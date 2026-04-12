# Phase 1: Base Configuration - Context

**Gathered:** 2026-04-05
**Status:** Ready for planning

<domain>
## Phase Boundary

CLI-Tool exists and functions with secure base configuration. This phase delivers the foundational CLI for step-by-step permission management with secure defaults.

</domain>

<decisions>
## Implementation Decisions

### CLI Interface Design
- **D-01:** CLI command is `gsd configure --step N` where N is the phase number
- **D-02:** Flags: `--step N` for phase selection, `--list` for showing available steps, `--verify` for dry-run verification
- **D-03:** Output format: human-readable with color-coded status indicators
- **D-04:** Error handling: descriptive error messages with suggested fixes

### State Management
- **D-05:** State file location: `.gsd/state.json` (in user's home or project root)
- **D-06:** State tracks: current phase, applied permissions, configuration version
- **D-07:** Migration support: state file includes schema version for future updates

### Security Defaults
- **D-08:** Default permissions: restrictive — ask before read, deny by default for write operations
- **D-09:** Auto-update: disabled by default (`"autoupdate": false`)
- **D-10:** External directory access: restricted to explicitly configured paths only
- **D-11:** Destructive commands: require explicit confirmation (`rm -rf` prompts)
- **D-12:** Permission wildcards: expand to specific paths, never full home directory

### Installation
- **D-13:** OpenCode installation: documented for both `brew install` and direct download
- **D-14:** Installation guide: step-by-step documentation in project root

</decisions>

<canonical_refs>
## Canonical References

**Downstream agents MUST read these before planning or implementing.**

No external specs — requirements fully captured in decisions above

### Research Summary
- `.planning/research/SUMMARY.md` — Phase 1 rationale and critical pitfalls
- `.planning/ROADMAP.md` §14-33 — Phase 1 details and success criteria
- `.planning/REQUIREMENTS.md` — INST-01 through CLI-02 requirements

### Project Context
- `.planning/PROJECT.md` — Core value: secure development with minimal necessary rights
- `.planning/PROJECT.md` §41-44 — Constraints: CLI command, network restrictions, skills limits

</canonical_refs>

<code_context>
## Existing Code Insights

### Reusable Assets
- None yet — Phase 1 creates the foundational CLI tool

### Established Patterns
- No existing patterns — this is the first phase building the core tool

### Integration Points
- CLI tool integrates with: macOS system APIs (defaults, TCC), state file (.gsd/state.json)

</code_context>

<specifics>
## Specific Ideas

The CLI tool must support step-by-step configuration. Each step unlocks specific permissions incrementally:
- Step 1: Install OpenCode + basic security defaults
- Step 2-4: Progressive permission expansion based on user needs

</specifics>

<deferred>
## Deferred Ideas

None — discussion stayed within phase scope

</deferred>

---

*Phase: 01-base-configuration*
*Context gathered: 2026-04-05*
