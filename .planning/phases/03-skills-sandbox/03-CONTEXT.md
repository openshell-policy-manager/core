# Phase 3: Skills & Sandbox - Context

**Gathered:** 2026-04-05
**Status:** Ready for planning

<domain>
## Phase Boundary

OpenCode skills and MCP servers are configurable and usable within the sandbox with minimal rights.

</domain>

<decisions>
## Implementation Decisions

### Skill Whitelist Strategy
- **D-01:** Only explicitly needed skills are activated (not all available skills)
- **Rationale:** Minimal necessary rights — project only needs Websearch, Codesearch, File Operations, Git Operations

### MCP Server Configuration
- **D-02:** MCP servers are automatically set up from central configuration
- **Rationale:** Consistent with CLI approach — configuration defined centrally, not manually per instance

### Skill Permission Levels
- **D-03:** Permissions are defined per skill (not globally)
- **Rationale:** Enables granular control — e.g. Websearch gets network, File Operations read-only

### Sandbox Isolation
- **D-04:** Combined security model: sandbox-exec for process isolation + OpenCode-internal permissions
- **Rationale:** Multi-layered security — redundancy for critical operations

### the agent's Discretion
- Skill-specific paths and permission granularity can be optimized by the planner

</decisions>

<canonical_refs>
## Canonical References

**Downstream agents MUST read these before planning or implementing.**

### Project Context
- `.planning/PROJECT.md` §41-44 — Constraints: Skills limits, CLI command
- `.planning/REQUIREMENTS.md` §35-42 — SAND-01 through SAND-06 Requirements
- `.planning/ROADMAP.md` §63-82 — Phase 3 Details

### Security
- `.planning/research/PITFALLS.md` §239-274 — Pitfall 8: Skill/MCP Unrestricted Loading

### Architecture
- `.planning/research/ARCHITECTURE.md` §307-328 — Skill config schema and registry

</canonical_refs>

<code_context>
## Existing Code Insights

### Reusable Assets
- `bin/gsd-configure` — CLI tool with step logic (Phase 3 already referenced)
- `.planning/config.json` — Central configuration with agent_skills field

### Integration Points
- OpenCode Permissions System — controls what skills are allowed to do
- MCP Server Configuration — via opencode.json or central config

</code_context>

<specifics>
## Specific Ideas

**Skills for Phase 3:**
- Websearch/Webfetch — for internet research (NETW-02)
- Codesearch — for code analysis
- File Operations — for file access
- Git Operations — for repository interaction

**Do not activate:**
- Agent tools that are not explicitly needed
- Skills not specified in requirements

</specifics>

<deferred>
## Deferred Ideas

None — discussion stayed within phase scope

</deferred>

---

*Phase: 03-skills-sandbox*
*Context gathered: 2026-04-05*
