# Phase 2: GitLab Integration - Discussion Log

> **Audit trail only.** Do not use as input to planning, research, or execution agents.
> Decisions are captured in CONTEXT.md — this log preserves the alternatives considered.

**Date:** 2026-04-05
**Phase:** 02-gitlab-integration
**Areas discussed:** GitLab Authentication, Network Security, Shell Environment, Internet Access

---

## GitLab Authentication

| Option | Description | Selected |
|--------|-------------|----------|
| OAuth Token | Standard OAuth flow with PKCE | |
| Deploy Token | Simpler, more secure than OAuth, better for CI/CD | ✓ |
| SSH Keys | Traditional Git auth, but less suitable for OpenCode | |

**User's choice:** Deploy Token
**Notes:** [auto] Deploy Token recommended — more secure than long-lived OAuth tokens, no refresh overhead, easy to rotate

---

## Network Security

| Option | Description | Selected |
|--------|-------------|----------|
| Blacklist | Everything allowed except explicitly blocked | |
| Whitelist | Only explicitly allowed domains | ✓ |

**User's choice:** Whitelist
**Notes:** [auto] Whitelist approach per PROJECT.md constraints — "Only download data that has been explicitly approved beforehand"

---

## Firewall Tool

| Option | Description | Selected |
|--------|-------------|----------|
| pfctl | System-integrated, no third-party | ✓ |
| Little Snitch | Commercial, more powerful | |
| OpenSnitch | Open-source, less stable | |

**User's choice:** pfctl
**Notes:** [auto] pfctl recommended — system-integrated, no additional dependencies, full control

---

## Shell Environment

| Option | Description | Selected |
|--------|-------------|----------|
| Relative paths allowed | Flexible paths, potential security risks | |
| Absolute paths required | More secure, prevents PATH manipulation | ✓ |

**User's choice:** Absolute paths required
**Notes:** [auto] Addresses Pitfall 6 from research/SUMMARY.md — Prevent Shell Environment Privilege Escalation

---

## Internet Access

| Option | Description | Selected |
|--------|-------------|----------|
| Free access | All ports allowed | |
| Controlled with whitelist | Only research skills allowed | ✓ |

**User's choice:** Controlled with whitelist — Free internet access for research (codesearch, websearch, webfetch), must be explicitly activated

---

## Deferred Ideas

None — all discussed points fit within phase scope.

---

*Phase: 02-gitlab-integration*
*Discussion log created: 2026-04-05*
*Mode: auto (all decisions auto-selected)*
