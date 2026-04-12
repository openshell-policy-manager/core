# Feature Research

**Domain:** OpenShell Security Configuration for macOS
**Researched:** 2026-04-05
**Confidence:** MEDIUM

## Feature Landscape

### Table Stakes (Users Expect These)

Features users assume exist. Missing these = product feels incomplete.

| Feature | Why Expected | Complexity | Notes |
|---------|--------------|------------|-------|
| Configure GitLab access | PROJECT.md explicitly requires read/write access to remote GitLab repository | MEDIUM | SSH keys or token authentication required |
| Control network access | Selective access to GitLab + free internet access for research is core requirement | HIGH | Firewall rules or proxy configuration for macOS |
| OpenCode Installation/Setup | OpenCode must be installed and configured | LOW | brew or direct download |
| CLI command for step-by-step configuration | Core feature from PROJECT.md - repeatable configuration application | MEDIUM | Individual steps must be executable in isolation |
| Security configuration with minimal rights | Core Value: minimal necessary rights for maximum security | HIGH | Sandbox constraints, filesystem access, process permissions |
| OpenCode sandbox with skills/MCP servers | PROJECT.md specifies all OpenCode skills within the sandbox | MEDIUM | Skills and MCP servers must be defined in configuration |

### Differentiators (Competitive Advantage)

Features that set the product apart. Not required, but valuable.

| Feature | Value Proposition | Complexity | Notes |
|---------|-------------------|------------|-------|
| Gradual permission expansion | Unique model: security first, then step-by-step unlocking — not the other way around | HIGH | Each step must be independently verifiable |
| Explicit approval check before network access | PROJECT.md: "Only download data that has been explicitly approved beforehand" | MEDIUM | Whitelist approach instead of blacklist |
| Security audit trail | Traceability of which permissions were unlocked when | LOW | Logging per CLI invocation |
| Configuration backup/versioning | Configurations can be rolled back if problems occur | LOW | Git-based configuration management |

### Anti-Features (Commonly Requested, Often Problematic)

Features that seem good but create problems.

| Feature | Why Requested | Why Problematic | Alternative |
|---------|---------------|-----------------|-------------|
| Full development environment (npm/node) | Developers expect a complete dev environment | PROJECT.md explicitly out of scope, contradicts Core Value | OpenCode-only, as specified in PROJECT.md |
| OpenShell installation | Product is called OpenShell, so it must be installed | PROJECT.md: "configuration only" | Configuration scripts for existing OpenShell instance |
| Admin features on GitLab | Makes product more powerful | PROJECT.md explicitly out of scope | Repository access only, no admin operations |
| Automatic permission expansion | "Smarter" configuration sounds better | Contradicts security principle: minimal necessary rights | Explicit, manual unlocking per step |
| Real-time monitoring dashboard | Modern products have dashboards | Increases complexity without security value, distracts from Core Value | CLI output + optional log files |

## Feature Dependencies

```
[Configure GitLab access]
        └──requires──> [OpenCode Installation/Setup]

[Control network access]
        └──requires──> [Configure GitLab access]

[CLI command for step-by-step configuration]
        ├──requires──> [Security configuration with minimal rights]
        └──requires──> [OpenCode sandbox with skills/MCP servers]

[Security configuration with minimal rights]
        └──requires──> [Control network access]

[Gradual permission expansion]
        └──enhances──> [CLI command for step-by-step configuration]
```

### Dependency Notes

- **GitLab access requires OpenCode installation:** Before GitLab can be configured, OpenCode must be installed
- **Network access requires GitLab access:** Network rules are defined based on GitLab URLs
- **Security configuration requires network access:** Permissions control what is reachable over the network
- **Gradual permission expansion enhances CLI:** The differentiating feature builds on the CLI

## MVP Definition

### Launch With (v1)

Minimum viable product — what's needed to validate the concept.

- [x] OpenCode Installation/Setup — Prerequisite for everything else, PROJECT.md requirement
- [x] Configure GitLab access — Core requirement for remote repository access
- [x] Control network access — GitLab + free internet access, PROJECT.md specifies
- [x] CLI command for step-by-step configuration — Core feature, PROJECT.md "CLI command for step-by-step configuration application"
- [x] Security configuration: Priority on security — Core Value "security first - minimal necessary rights"

### Add After Validation (v1.x)

Features to add once core is working.

- [ ] Security audit trail — Logging which steps were unlocked
- [ ] Configuration backup/versioning — Rollback capability if problems occur

### Future Consideration (v2+)

Features to defer until product-market fit is established.

- [ ] Advanced sandbox constraints — Advanced permission options
- [ ] Integration with other tools — More than just OpenCode
- [ ] Dashboard for configuration overview — Monitoring without real-time complexity

## Feature Prioritization Matrix

| Feature | User Value | Implementation Cost | Priority |
|---------|------------|---------------------|----------|
| OpenCode Installation/Setup | HIGH | LOW | P1 |
| Configure GitLab access | HIGH | MEDIUM | P1 |
| Control network access | HIGH | HIGH | P1 |
| CLI command for step-by-step configuration | HIGH | MEDIUM | P1 |
| Security configuration with minimal rights | HIGH | HIGH | P1 |
| OpenCode sandbox with skills/MCP servers | HIGH | MEDIUM | P1 |
| Gradual permission expansion | HIGH | HIGH | P1 (Differentiator) |
| Security audit trail | MEDIUM | LOW | P2 |
| Configuration backup/versioning | MEDIUM | LOW | P2 |

**Priority key:**
- P1: Must have for launch
- P2: Should have, add when possible
- P3: Nice to have, future consideration

## Competitor Feature Analysis

| Feature | OpenShell (NVIDIA) | OpenClaw | Our Approach |
|---------|-------------------|----------|--------------|
| Sandbox security | 4-Layer-Security (network, filesystem, process, inference) | Active security-hardening commits | Minimal necessary rights, step-by-step unlocking |
| Network access | Policy-based | Allowlist for exec | Explicit approval check + gradual expansion |
| CLI tool | Rust-based | TypeScript-based | Shell scripts for macOS configuration |
| Configuration format | YAML Policy | TypeScript Config | JSON/Shell-based for PROJECT.md spec |

**Our Differentiator:** Gradual permission expansion — while other tools have a static policy, our approach enables step-by-step unlocking with explicit approval checks.

## Sources

- NVIDIA OpenShell Security Best Practices (docs.nvidia.com/openshell/latest/security/best-practices.html)
- OpenClaw Security Commits (github.com/openclaw/openclaw)
- OpenCode Config Documentation (opencode.ai/docs/config/)
- PROJECT.md — Specifications and requirements for this project

---

*Feature research for: OpenShell Security Configuration macOS*
*Researched: 2026-04-05*
