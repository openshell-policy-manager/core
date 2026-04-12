# Phase 2: GitLab Integration - Context

**Gathered:** 2026-04-05
**Status:** Ready for planning

<domain>
## Phase Boundary

GitLab access works with secure network configuration. This phase configures authentication and network access for OpenCode with GitLab repositories.

</domain>

<decisions>
## Implementation Decisions

### GitLab Authentication
- **D-01:** Authentication method: **Deploy-Token** — simpler than OAuth, more secure than long-lived OAuth tokens, better than SSH keys for OpenCode workflows
- **D-02:** Deploy-Token Storage: **`~/.config/opencode/credentials/gitlab-deploy-token`** — separate directory, not ~/.git-credentials (prevents override)

### Network Security
- **D-03:** Network policy: **Whitelist** — only explicitly allowed domains (gitlab.com, gitlab.*.com)
- **D-04:** Firewall tool: **pfctl** — system-integrated, no third-party dependency
- **D-05:** Allowed network targets:
  - `*.gitlab.com` — GitLab SaaS
  - `gitlab.*.company.com` — company-specific instances (placeholder)
  - `api.github.com` only if explicitly configured (no default permission)

### Shell Environment
- **D-06:** Shell environment hardening: **Absolute paths required** — prevents PATH manipulation and privilege escalation (Pitfall 6 from research/SUMMARY.md)

### Internet Access
- **D-07:** Web research: **Free internet access for research** (codesearch, websearch, webfetch) — must be explicitly activated, not default-on

</decisions>

<canonical_refs>
## Canonical References

**Downstream agents MUST read these before planning or implementing.**

### Research
- `.planning/research/SUMMARY.md` §102-119 — Phase 2 rationale and pitfalls (Pitfall 5: Credentials Persistence, Pitfall 6: Shell Environment Privilege Escalation)

### Requirements
- `.planning/REQUIREMENTS.md` — GITL-01 through GITL-04, NETW-01 through NETW-03

### Project Context
- `.planning/PROJECT.md` — Core value: minimal necessary rights
- `.planning/PROJECT.md` §41-44 — Constraints: Network only for explicitly approved sources

</canonical_refs>

<code_context>
## Existing Code Insights

### Reusable Assets
- Phase 1 CLI tool (gsd configure) — will be extended for this phase with new sub-commands
- State file (.gsd/state.json) — already stores configured credentials and network rules

### Integration Points
- CLI extended: `gsd configure --step 2` applies GitLab + Network config
- State file: new fields for `gitlab_token`, `network_whitelist`, `websearch_enabled`

</code_context>

<specifics>
## Specific Ideas

**Token Handling:**
- Deploy token must be configured with expiration (short-lived, max 1 year)
- Token must be stored in environment variable, not as plain text in config

**Network Enforcement:**
- pfctl rules must be stored in /etc/pf.rules or equivalent
- Backup rules if pfctl is not available (fallback to hosts file)

</specifics>

<deferred>
## Deferred Ideas

None — discussion stayed within phase scope

</deferred>

---

*Phase: 02-gitlab-integration*
*Context gathered: 2026-04-05*
