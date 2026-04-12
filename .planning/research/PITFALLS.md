# Pitfalls Research

**Domain:** OpenShell Security Configuration for OpenCode on macOS
**Researched:** 2026-04-05
**Confidence:** HIGH

## Critical Pitfalls

### Pitfall 1: Overly Permissive Default Permissions

**What goes wrong:**
OpenCode starts with permissive defaults — most permissions default to `"allow"`. Without explicit configuration, the agent can read/write files and execute shell commands freely. This is equivalent to running an auto-updating remote access tool with full user privileges.

**Why it happens:**
The UX-first design philosophy prioritizes "it just works" over security. Issue #5076 (49 reactions) explicitly calls this out as "malware-class risk" — if the update pipeline, binary distribution, or model provider is compromised, attackers get full access to the system.

**How to avoid:**
Override defaults explicitly in `opencode.json`:

```json
{
  "permission": {
    "*": "ask",
    "read": { "*": "allow", "*.env": "deny", "*.env.*": "deny" },
    "edit": "deny",
    "bash": { "*": "deny", "git status *": "allow", "git pull *": "allow" },
    "external_directory": "ask"
  }
}
```

**Warning signs:**
- No permission prompts when OpenCode runs first time
- Agent can execute any bash command without confirmation
- `.env` files can be read (verify deny rules are in place)

**Phase to address:** Phase 1 — Base Configuration (foundational security)

---

### Pitfall 2: Uncontrolled External Directory Access

**What goes wrong:**
OpenCode can access paths outside the working directory where it was started. Without `external_directory` configuration, the agent can read/write to the entire filesystem. Issue #15163 shows OpenCode CLI scans outside workspace, triggering macOS security alerts.

**Why it happens:**
The `external_directory` permission defaults to `"ask"`, but users often click "always allow" without understanding the scope. The wildcard expansion (`~` → `/Users/username`) makes it easy to accidentally grant broad access.

**How to avoid:**
Explicitly whitelist only required directories:

```json
{
  "permission": {
    "external_directory": {
      "~/projects/myproject/**": "allow",
      "~/.ssh/**": "deny"
    }
  }
}
```

**Warning signs:**
- OpenCode asks to access directories outside project
- Agent can read `~/Downloads`, `~/Documents`, or other sensitive paths
- First-run prompts grant blanket access

**Phase to address:** Phase 1 — Base Configuration

---

### Pitfall 3: Destructive Commands Without Safeguards

**What goes wrong:**
Issue #17949 reports OpenCode deleted the entire Downloads folder without asking. Without proper `bash` permission patterns, `rm -rf *` or similar destructive commands can execute silently.

**Why it happens:**
The default `bash` permission is `"allow"`, and pattern matching requires explicit command prefixes. Users don't realize that `rm *` needs the same permission as `rm` alone.

**How to avoid:**
Configure granular bash rules with explicit deny patterns:

```json
{
  "permission": {
    "bash": {
      "*": "ask",
      "git *": "allow",
      "npm *": "allow",
      "grep *": "allow",
      "rm -rf *": "deny",
      "rm -rf": "deny",
      "dd *": "deny",
      "mkfs *": "deny"
    }
  }
}
```

**Warning signs:**
- Bash commands execute without prompt
- No pattern matching on command arguments
- Agent can run `curl | bash` or similar remote execution

**Phase to address:** Phase 1 — Base Configuration

---

### Pitfall 4: Auto-Update Executes Unverified Code

**What goes wrong:**
Auto-update is enabled by default. OpenCode silently replaces its own binary on startup. Combined with full file/shell access, a compromised update pipeline means full system compromise.

**Why it happens:**
Auto-update defaults to enabled for UX convenience. Users don't review each new version before it's applied.

**How to avoid:**
Disable auto-update and use manual updates:

```json
{
  "autoupdate": false
}
```

Or set to "notify" to be informed of updates without auto-installing:

```json
{
  "autoupdate": "notify"
}
```

**Warning signs:**
- OpenCode updates without user notification
- New version appears without explicit user action
- No changelog review before update

**Phase to address:** Phase 1 — Base Configuration

---

### Pitfall 5: GitLab Credentials Persist Beyond Session

**What goes wrong:**
Issue #10950 shows stored OAuth credentials silently override explicit provider config. Tokens persist and may be reused across sessions, creating credential leakage risk.

**Why it happens:**
OAuth tokens are stored in `~/.config/opencode/` and persist. The config loading order means stored credentials take precedence over explicitly configured ones.

**How to avoid:**
- Use short-lived tokens with explicit expiration
- Clear credentials before sharing config files
- Verify `OPENCODE_CONFIG_CONTENT` precedence (Issue #11628 — fixed)

```json
{
  "providers": {
    "gitlab": {
      "type": "oauth",
      "expires_at": "2026-04-06T00:00:00Z"
    }
  }
}
```

**Warning signs:**
- Config file shows credentials but agent ignores them
- Token persists after logout
- Different behavior between sessions

**Phase to address:** Phase 2 — GitLab Integration

---

### Pitfall 6: Shell Environment Inherited Privilege Escalation

**What goes wrong:**
Bash commands inherit the full user environment including potentially dangerous variables (`PATH`, `LD_PRELOAD`, etc.). A compromised or maliciously crafted environment can escalate privileges.

**Why it happens:**
OpenCode executes bash in the user login shell context. All shell initialization scripts (`~/.bashrc`, `~/.zshrc`, `/etc/zshenv`) execute before commands.

**How to avoid:**
Use absolute paths in permission rules and consider a clean environment:

```json
{
  "permission": {
    "bash": {
      "*": "ask",
      "/usr/bin/git *": "allow",
      "/usr/local/bin/npm *": "allow"
    }
  }
}
```

For higher security, run OpenCode in a minimal environment or use sandboxing (Issue #9647 — implemented).

**Warning signs:**
- Commands execute with unexpected `PATH`
- Aliases or functions from shell init affect behavior
- Environment variables leak between sessions

**Phase to address:** Phase 2 — GitLab Integration

---

### Pitfall 7: Missing CVE-2026-22812 HTTP Server Authentication

**What goes wrong:**
CVE-2026-22812 is a critical vulnerability making HTTP server authentication mandatory. If running OpenCode's HTTP server without authentication, remote attackers can execute commands.

**Why it happens:**
Previous versions allowed unauthenticated HTTP server access. The fix (PR #9328) makes authentication mandatory.

**How to avoid:**
- Update to latest OpenCode version (v1.1.1+)
- If running HTTP server, configure authentication:

```json
{
  "server": {
    "auth": "required"
  }
}
```

**Warning signs:**
- OpenCode running with `--server` flag without auth config
- HTTP interface accessible without login
- Version pre-v1.1.1

**Phase to address:** Phase 1 — Base Configuration

---

### Pitfall 8: Skill/MCP Server Unrestricted Loading

**What goes wrong:**
Skills and MCP servers run with full OpenCode privileges. A compromised skill can read files, execute commands, and access network resources. Issue #2242 shows users want sandboxing but it's not fully implemented.

**Why it happens:**
The `skill` permission defaults to `"allow"` and loads from any configured source. No verification of skill integrity.

**How to avoid:**
Configure skill permissions restrictively:

```json
{
  "permission": {
    "skill": {
      "*": "deny",
      "git-helper": "allow",
      "code-review": "allow"
    }
  },
  "mcp": {
    "servers": {
      "trusted-server": {
        "allowed": true
      }
    }
  }
}
```

**Warning signs:**
- Skills load without confirmation
- MCP servers execute arbitrary code
- No skill source verification

**Phase to address:** Phase 3 — Skills & MCP Configuration

---

## Technical Debt Patterns

| Shortcut | Immediate Benefit | Long-term Cost | When Acceptable |
|----------|-------------------|----------------|-----------------|
| Using `"allow"` for all permissions | No prompts, fast testing | Full system access for any agent/model | Never — use `"ask"` minimum |
| Skipping `.env` deny rules | Can read all config files | Credential exposure risk | When using secret management |
| Disabling `external_directory` entirely | No prompts for external access | Cannot work across projects | Only with per-project config |
| Using personal access tokens | Simple auth setup | Token lives forever if not revoked | When tokens have expiry + limited scope |

---

## Integration Gotchas

| Integration | Common Mistake | Correct Approach |
|-------------|----------------|------------------|
| GitLab | Using long-lived tokens | Use OAuth with short expiry or deploy tokens with limited scope |
| GitLab | Not configuring redirect URI | Verify OAuth callback URL matches config |
| Network | Allowing all webfetch URLs | Whitelist domains: `"webfetch": { "github.com/*": "allow" }` |
| MCP Servers | Trusting all server configurations | Verify server sources, use signed configs |

---

## Security Mistakes

| Mistake | Risk | Prevention |
|---------|------|------------|
| No permission config at all | Full file/system access | Explicit permission config before first run |
| Granting `bash: "allow"` | Arbitrary command execution | Use pattern matching: `"git *": "allow"` |
| External directory with `~/**` | Full home directory access | Specific paths only: `"~/projects/myproject/**"` |
| Auto-update enabled | Unknown code execution | `"autoupdate": false` or `"notify"` |
| No `.env` protection | Credential theft | Explicit deny rules in place |

---

## "Looks Done But Isn't" Checklist

- [ ] **Permission config:** File exists but missing granular rules — verify patterns cover commands with arguments
- [ ] **External directory:** Configured but allows too many paths — verify only needed directories allowed
- [ ] **GitLab auth:** Token works but may persist incorrectly — verify logout clears credentials
- [ ] **Security defaults:** OpenCode asks for permission but "always" was clicked — reset to test config
- [ ] **Auto-update:** Assumed disabled but no explicit config — verify in `opencode.json`

---

## Recovery Strategies

| Pitfall | Recovery Cost | Recovery Steps |
|---------|---------------|----------------|
| Accidentally granted all permissions | LOW | Delete `~/.config/opencode/permission` cache, rewrite config with deny rules |
| Deleted files with rm -rf | HIGH | Use Time Machine or backups — no OpenCode recovery possible |
| Token leaked | MEDIUM | Revoke in GitLab, regenerate, update config |
| CVE vulnerability | MEDIUM | Update OpenCode: `brew upgrade opencode` or download new version |

---

## Pitfall-to-Phase Mapping

| Pitfall | Prevention Phase | Verification |
|---------|------------------|--------------|
| Overly permissive defaults | Phase 1: Base Configuration | Test first run — should prompt for actions |
| External directory access | Phase 1: Base Configuration | Verify agent cannot read ~/Downloads without config |
| Destructive bash commands | Phase 1: Base Configuration | Test `rm -rf *` — should be denied |
| Auto-update unverified code | Phase 1: Base Configuration | Verify config has `"autoupdate": false` |
| GitLab credential persistence | Phase 2: GitLab Integration | Logout should clear stored tokens |
| Shell environment escalation | Phase 2: GitLab Integration | Verify commands run with clean PATH |
| CVE-2026-22812 HTTP auth | Phase 1: Base Configuration | Verify version ≥ v1.1.1 |
| Skill/MCP unrestricted loading | Phase 3: Skills & MCP | Verify only trusted skills load |

---

## Sources

- OpenCode Permissions Documentation — https://opencode.ai/docs/permissions/
- Issue #5076: OpenCode should have better/safer defaults — https://github.com/anomalyco/opencode/issues/5076
- Issue #17949: OpenCode deleted Downloads folder — https://github.com/anomalyco/opencode/issues/17949
- Issue #15163: OpenCode CLI scans outside workspace — https://github.com/anomalyco/opencode/issues/15163
- Issue #6355: RCE and file read vulnerability — https://github.com/anomalyco/opencode/issues/6355
- Issue #10950: Stored OAuth credentials override config — https://github.com/anomalyco/opencode/issues/10950
- CVE-2026-22812: HTTP Server Authentication — https://github.com/anomalyco/opencode/pull/9328
- Issue #2242: Sandbox feature request — https://github.com/anomalyco/opencode/issues/2242

---

*Pitfalls research for: OpenShell Security Configuration*
*Researched: 2026-04-05*
