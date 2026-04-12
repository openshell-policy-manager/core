# Phase 3: Skills & Sandbox - Discussion Log

> **Audit trail only.** Do not use as input to planning, research, or execution agents.
> Decisions are captured in CONTEXT.md — this log preserves the alternatives considered.

**Date:** 2026-04-05
**Phase:** 03-skills-sandbox
**Areas discussed:** Skill Whitelist Strategy, MCP Server Configuration, Skill Permission Levels, Sandbox Isolation

---

## Skill Whitelist Strategy

| Option | Description | Selected |
|--------|-------------|----------|
| Enable all skills | Maximum functionality — all available OpenCode skills | |
| Only explicitly needed skills | Minimal necessary rights — only skills that requirements specify | ✓ |
| No skills enabled | Maximum security — no external capabilities | |

**User's choice:** Only explicitly needed skills (SAND-01 through SAND-06 from REQUIREMENTS.md)
**Notes:** Minimal Necessary Rights — project only needs Websearch, Codesearch, File Operations, Git Operations

---

## MCP Server Configuration

| Option | Description | Selected |
|--------|-------------|----------|
| Centrally configured via CLI | Consistent approach — config defined centrally | ✓ |
| Manually per instance | More flexible but error-prone | |

**User's choice:** Automatically from central configuration (matching CLI philosophy)
**Notes:** Combinable with `gsd configure --step 3`

---

## Skill Permission Levels

| Option | Description | Selected |
|--------|-------------|----------|
| Global for all skills | Simpler — one permission rule for all | |
| Per skill individually | Granular — e.g. Websearch gets network, Files read-only | ✓ |

**User's choice:** Define per skill (enables granular control)
**Notes:** Detailed permissions can be specified by the planner

---

## Sandbox Isolation

| Option | Description | Selected |
|--------|-------------|----------|
| Only OpenCode-internal permissions | Simple configuration | |
| sandbox-exec for process isolation | Additional shell isolation layer | ✓ |
| Combined: OpenCode + sandbox-exec | Multi-layered security — redundancy | ✓ |

**User's choice:** Combined model with sandbox-exec + OpenCode permissions
**Notes:** Redundancy for critical operations increases security

---

## the agent's Discretion

- Skill-specific paths and permission granularity — planner can optimize based on OpenCode documentation

## Deferred Ideas

None — discussion stayed within phase scope.
