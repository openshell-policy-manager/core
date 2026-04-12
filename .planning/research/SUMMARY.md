# Project Research Summary

**Project:** OpenShell Security Configuration macOS
**Domain:** macOS Security Configuration System (CLI Tool)
**Researched:** 2026-04-05
**Confidence:** MEDIUM

## Executive Summary

This project develops a CLI tool for step-by-step configuration of OpenCode security settings on macOS. The core principle is "security first" — minimal rights are gradually expanded instead of having permissive defaults. The research shows that OpenCode starts with extensive permissions by default (read/write/execute without prompting), which represents a "malware-class risk" (Issue #5076).

The recommended architecture follows a three-stage pattern: Definition Layer (YAML/JSON Configs) → Execution Engine (CLI + Logic) → System Interface (defaults, TCC, pfctl). The stack is based on macOS system tools (sandbox-exec, tccplus, pfctl) combined with modern CLI tools (Starship Prompt, Ghostty Terminal). The biggest risks are permissive defaults, unrestricted external directory access, and auto-update without verification.

## Key Findings

### Recommended Stack

**Core technologies:**
- **zsh 5.9+** — Standard shell on macOS, better plugin architecture than bash
- **Starship Prompt 1.21+** — 2026 Recommended alternative to Oh My Zsh, Rust-based, fast
- **Ghostty 1.3+** — Modern terminal emulator, Rust-based, more performant than iTerm2
- **sandbox-exec** — System-integrated (macOS 10.7+), for process isolation without third-party tools
- **tccplus** — Fork of tccutil with Grant/Remove functions for finer TCC control
- **stronghold** — Community tool (1.2k stars) for automated macOS security settings
- **pfctl** — Built-in firewall control for network rules

**Do not use:**
- Direct TCC.db manipulation via SQL (Apple changes schema on updates)
- Completely disabling SIP
- Full root access without minimal-necessary principle

### Expected Features

**Must have (table stakes):**
- **Configure GitLab access** — SSH keys or token auth for remote repository access
- **Control network access** — Firewall rules for GitLab + free internet access
- **OpenCode Installation/Setup** — Prerequisite for everything else
- **CLI command for step-by-step configuration** — Individual steps must be executable in isolation
- **Security configuration with minimal rights** — Core Value: minimal necessary rights
- **OpenCode sandbox with skills/MCP servers** — Skills must be defined in configuration

**Should have (competitive):**
- **Gradual permission expansion** — Differentiating model: security first, then step-by-step unlocking
- **Explicit approval check before network access** — Whitelist approach instead of blacklist
- **Security audit trail** — Logging which permissions were unlocked
- **Configuration backup/versioning** — Rollback capability if problems occur

**Defer (v2+):**
- Advanced sandbox constraints
- Integration with other tools
- Dashboard for configuration overview

### Architecture Approach

The recommended architecture separates configuration definition from execution and enables the "step-by-step" deployment model. Four layers:

1. **CLI Interface** — User-facing command parser (gsd configure --step N)
2. **Configuration Layer** — Definitions (YAML/JSON), Phase Manifest, Validation Rules
3. **Execution Engine** — Network Manager, Skill Registry, Permission Controller
4. **System Interface** — macOS APIs: defaults, tccutil, security, pfctl

**Major components:**
1. **Network Manager** — Firewall rules (pfctl), hosts file, allowed domains
2. **Skill Registry** — Enable/disable OpenCode capabilities
3. **Permission Controller** — TCC database, sandbox settings, entitlements

### Critical Pitfalls

1. **Overly Permissive Default Permissions** — OpenCode starts with "allow" defaults. Configure with `"*": "ask"`, `"edit": "deny"`, `*.env` deny rules
2. **Uncontrolled External Directory Access** — Wildcard `~` expands to full home directory. Only allow specific paths
3. **Destructive Commands Without Safeguards** — `rm -rf *` can execute without prompting. Explicit deny patterns for destructive commands
4. **Auto-Update Executes Unverified Code** — Auto-update is enabled by default. `"autoupdate": false` or `"notify"`
5. **GitLab Credentials Persist Beyond Session** — Tokens persist in ~/.config/opencode/. Short-lived tokens with expiration
6. **CVE-2026-22812 HTTP Server Authentication** — Unauthenticated HTTP Server allows Remote Code Execution. Update to v1.1.1+ and auth required

## Implications for Roadmap

Based on research, suggested phase structure:

### Phase 1: Base Configuration (CLI + Core Security)
**Rationale:** Foundation for everything else — without a working CLI and base security, all other phases are worthless. Critical pitfalls (permissive defaults, CVE-2026-22812) must be addressed first.

**Delivers:**
- CLI argument parser with --step N, --list, --verify flags
- State file management (.gsd/state.json)
- Configuration loader (YAML)
- Base security configuration: permission defaults, autoupdate disable, CVE fix

**Addresses (FEATURES.md):**
- OpenCode Installation/Setup (P1)
- Security configuration with minimal rights (P1)

**Avoids (PITFALLS.md):**
- Pitfall 1: Overly Permissive Defaults
- Pitfall 4: Auto-Update unverified code
- Pitfall 7: CVE-2026-22812

**Research Flags:** Standard patterns — CLI parsing and state management are established patterns, no deeper research needed.

---

### Phase 2: GitLab Integration
**Rationale:** Network access and GitLab access are the core requirements from PROJECT.md. Network rules are defined based on GitLab URLs — hence this order.

**Delivers:**
- GitLab authentication (OAuth/deploy tokens)
- Network whitelist for gitlab.com + company-internal instances
- Shell environment hardening (absolute paths)

**Addresses (FEATURES.md):**
- Configure GitLab access (P1)
- Control network access (P1)

**Avoids (PITFALLS.md):**
- Pitfall 5: GitLab Credentials Persistence
- Pitfall 6: Shell Environment Privilege Escalation

**Research Flags:** Needs research — GitLab OAuth Flow specifics, Token expiration handling

---

### Phase 3: Skills & MCP Configuration
**Rationale:** Skills and MCP servers run with full OpenCode rights. After network and permissions are configured, skills can be restricted.

**Delivers:**
- Skill whitelist/blacklist configuration
- MCP Server trust configuration
- Audit trail logging

**Addresses (FEATURES.md):**
- OpenCode sandbox with skills/MCP servers (P1)
- Gradual permission expansion (P1 Differentiator)
- Security audit trail (P2)

**Avoids (PITFALLS.md):**
- Pitfall 8: Skill/MCP Unrestricted Loading

**Research Flags:** Needs research — OpenCode skill permission schema specifics

---

### Phase 4: Verification & Polish
**Rationale:** After the three core phases, the result must be verifiable. Rollback and export/import make the tool production-ready.

**Delivers:**
- Verification reporter (defined vs. applied diff)
- Rollback support
- Dry-run mode
- Configuration export/import

**Addresses (FEATURES.md):**
- CLI command for step-by-step configuration (P1) — verifiable
- Configuration backup/versioning (P2)

**Research Flags:** Standard patterns — verification and rollback are known patterns, no deep research needed.

---

### Phase Ordering Rationale

1. **CLI + Base Security first** — Everything builds on the CLI; critical security vulnerabilities (CVE, permissive defaults) must be fixed before any operation
2. **GitLab + Network second** — PROJECT.md specifies "GitLab access" and "control network access" as core requirements; network rules need GitLab URLs
3. **Skills third** — Skills need network and permissions as foundation; run with full rights and must be restricted
4. **Verification last** — Each phase should be verifiable, but a dedicated verification phase makes the tool production-ready

**Dependencies from FEATURES.md:**
- GitLab access requires OpenCode installation
- Network access requires GitLab access
- Skills require network + permissions

## Confidence Assessment

| Area | Confidence | Notes |
|------|------------|-------|
| Stack | MEDIUM | Shell/Terminal Tools HIGH (Starship/Ghostty are current 2026 recommendations); TCC Management MEDIUM (tccplus actively maintained but Apple changes APIs) |
| Features | MEDIUM | Based on PROJECT.md requirements; Differentiator "gradual permission expansion" is unique |
| Architecture | MEDIUM | Layered architecture pattern is proven (stronghold, NIST mSCP); Build order has clear dependencies |
| Pitfalls | HIGH | Direct references to OpenCode issues (#5076, #17949, #15163); CVE-2026-22812 is verified |

**Overall confidence:** MEDIUM

### Gaps to Address

- **OpenCode permission schema details:** The exact schema for `opencode.json` permission rules is not fully documented. Verify during implementation.
- **TCC.db changes:** Apple changes TCC schema on updates. Tool must be resilient to schema changes — tccplus uses official APIs where possible.
- **GitLab API specifics:** OAuth flow and token handling must be verified during Phase 2 implementation.

## Sources

### Primary (HIGH confidence)
- OpenCode Permissions Documentation — https://opencode.ai/docs/permissions/
- Issue #5076: OpenCode should have better/safer defaults — https://github.com/anomalyco/opencode/issues/5076
- Issue #17949: OpenCode deleted Downloads folder — https://github.com/anomalyco/opencode/issues/17949
- CVE-2026-22812: HTTP Server Authentication — https://github.com/anomalyco/opencode/pull/9328

### Secondary (MEDIUM confidence)
- NVIDIA OpenShell Security Best Practices — docs.nvidia.com/openshell/latest/security/best-practices.html
- stronghold CLI tool — https://github.com/alichtman/stronghold
- tccplus fork — community fork of tccutil

### Tertiary (LOW confidence)
- Ghostty Terminal Emulator — https://github.com/ghostty-org/ghostty (newer, less established than iTerm2)
- OpenSnitch — Open-source network filtering (less polished than Little Snitch)

---

*Research completed: 2026-04-05*
*Ready for roadmap: yes*
