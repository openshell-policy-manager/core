# OpenShell Policy Manager

## Project Context

This project is a **Policy Preset Manager** for NVIDIA OpenShell — an open-source runtime that provides sandbox isolation for autonomous AI agents. OpenShell applies browser-like isolation principles to the agentic workflow: every session is sandboxed, every resource is metered, and every permission is verified by the runtime before execution.

**What is OpenShell?**  
OpenShell is NVIDIA's open-source security runtime for autonomous agents. It provides three core pillars:

1. **Programmable Sandboxes** — Purpose-built isolation for agents that modify their own environment. Handles skill verification and network isolation, providing a "break-safe" environment where agents can experiment without touching the host.

2. **Granular Policy Engine** — Controls the "What," "Where," and "How" of execution. The engine evaluates actions at binary, path, and method levels. It allows agents to be autonomous where it matters — like installing a verified skill — while blocking unreviewed binaries or unauthorized network calls.

3. **Private Inference Router** — Keeps sensitive data on-device using local open models, routing to frontier models only when cost and privacy policies allow.

### Policy Structure

OpenShell policies are defined in YAML and control:

| Section | Type | Description |
|---------|------|-------------|
| `filesystem_policy` | Static | Controls which directories the agent can read vs read/write. Paths not listed are inaccessible. |
| `landlock` | Static | Kernel-level enforcement. `best_effort` uses highest ABI the host supports. |
| `process` | Static | OS-level identity for the agent process (`run_as_user`, `run_as_group`). |
| `network_policies` | Dynamic | Hot-reloadable network access controls. Named blocks of endpoints + binaries allowed to reach them. |

### This Project's Purpose

This project manages **policy presets** for OpenShell sandboxes. It provides:

- A collection of reusable YAML policy configurations for common use cases (brew, npm, pypi, docker, discord, etc.)
- A CLI for applying and managing these presets
- Integration with OpenShell's sandbox creation and policy update workflows

The presets in this project are compatible with OpenShell's policy schema and can be used directly with `openshell sandbox create --policy`.

### Reference Presets

This project includes presets inspired by NVIDIA's NemoClaw blueprint presets:

- `brew.yaml` — Homebrew package manager access
- `npm.yaml` — npm package registry access
- `pypi.yaml` — Python package index access
- `docker.yaml` — Docker registry access
- `discord.yaml` — Discord API access
- `slack.yaml` — Slack API access
- And more...

See [NVIDIA/NemoClaw nemoclaw-blueprint/policies/presets](https://github.com/NVIDIA/NemoClaw/tree/main/nemoclaw-blueprint/policies/presets) for the official preset collection.

## Current State

**Shipped:** v1.0 MVP (2026-04-05)
**Phases:** 5 complete | **Plans:** 5 | **Tasks:** 9

v1.0 delivered a secure OpenShell configuration CLI with step-by-step permission management, GitLab integration via Deploy-Token, skills/sandbox configuration, and verification capabilities. Phase 999.3 added automated Homebrew distribution via GoReleaser with dual alpha/stable taps.

**Known gaps carried forward:**
- SAND-01 through SAND-06: Skills & Sandbox requirements lack formal verification (no VERIFICATION.md files)
- OSHELL-07: GitLab URL memorization not yet implemented

## What This Is

A secure OpenShell configuration for local development with OpenCode and a remote GitLab repository. The configuration is applied step by step through a CLI command to gradually unlock rights and access.

## Core Value

Secure development environment: OpenCode can perform all required operations, but only with the minimal necessary rights and access.

## Requirements

### Validated

- ✓ INST-01: OpenCode can be installed via brew — v1.0
- ✓ INST-02: OpenCode can be installed via direct download — v1.0
- ✓ INST-03: Installation guide is documented — v1.0
- ✓ BASE-01: CLI command exists for configuration application — v1.0
- ✓ BASE-02: Configuration uses secure defaults — v1.0
- ✓ BASE-03: Auto-update is disabled by default — v1.0
- ✓ BASE-04: External directory access is restricted — v1.0
- ✓ BASE-05: Destructive commands require confirmation — v1.0
- ✓ GITL-01: GitLab OAuth tokens can be configured — v1.0
- ✓ GITL-02: Read access to GitLab repositories works — v1.0
- ✓ GITL-03: Write access to GitLab repositories works — v1.0
- ✓ GITL-04: GitLab credentials are stored correctly — v1.0
- ✓ NETW-01: Network access to GitLab is allowed — v1.0
- ✓ NETW-02: Free internet access for web research is enabled — v1.0
- ✓ NETW-03: Only data from explicitly approved sources — v1.0
- ✓ CLI-01: A single command applies the configuration — v1.0
- ✓ CLI-02: Command can gradually unlock permissions — v1.0
- ✓ CLI-03: Command shows current configuration status — v1.0
- ✓ CLI-04: Dry-run mode for testing without changes — v1.0

### Active

- [ ] SAND-01: OpenCode sandbox with all skills is configurable (verification needed)
- [ ] SAND-02: MCP servers approved in OpenCode can be used (verification needed)
- [ ] SAND-03: Websearch/Webfetch skills are usable (verification needed)
- [ ] SAND-04: Codesearch skill is usable (verification needed)
- [ ] SAND-05: File operations skills are usable (verification needed)
- [ ] SAND-06: Git operations (status, diff, commit) are usable (verification needed)
- [ ] OSHELL-07: Remember GitLab URL (do not ask again)

### Out of Scope

| Feature | Reason |
|---------|--------|
| OpenShell installation | Configuration only, no installation |
| Full dev environment (npm/node, git cli) | Out of scope — OpenCode only |
| Admin features on GitLab | Not required for development |
| Docker container isolation | macOS native sandbox-exec sufficient |

## Context

- **Target system:** macOS
- **GitLab:** Remote repository with read and write access
- **Network:** Selective access to GitLab, free internet access for web research
- **Security model:** Security first — minimal necessary rights
- **Tech stack:** Bash CLI scripts, JSON state management
- **Codebase:** ~435 LOC across bin/gsd, bin/gsd-configure, bin/gsd-gitlab
- **Key files:** .gsd/state.json, .gsd/network-rules.json, .planning/config.json

## Constraints

- **Network:** Only download data that has been explicitly approved beforehand
- **Skills:** Only skills and MCP servers approved in OpenCode
- **CLI:** Simple command for step-by-step configuration application
- **Dev Tools:** OpenCode only (no full dev environment)

## Key Decisions

| Decision | Rationale | Outcome |
|----------|-----------|---------|
| CLI command is `gsd configure --step N` | Simple, discoverable, phased approach | ✓ Good — used across all 4 phases |
| Deploy-Token auth over OAuth | Simpler, more secure for CI/CD workflows | ✓ Good — implemented in v1.0 |
| Network whitelist policy | Only explicitly allowed domains | ✓ Good — gitlab.com whitelisted |
| Skill whitelist strategy | Only explicitly needed skills activated | ✓ Good — 5 skills configured |
| State file at `.gsd/state.json` | Centralized, versioned state tracking | ✓ Good — schema v1.0 |
| Auto-update disabled by default | Security first principle | ✓ Good — matches core value |
| Permission wildcards never expand to full home | Prevents over-broad access | ✓ Good — specific paths only |
| Destructive commands require confirmation | Prevents accidental data loss | ✓ Good — safety mechanism |

---

*Last updated: 2026-04-12 after phase 999.3 completion*
