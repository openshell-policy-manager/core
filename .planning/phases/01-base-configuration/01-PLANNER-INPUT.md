# Phase 1 Plan: Base Configuration

**Phase:** 01-base-configuration
**Mode:** standard

## Files to Read

- `.planning/STATE.md` - Project State
- `.planning/ROADMAP.md` - Roadmap
- `.planning/REQUIREMENTS.md` - Requirements
- `.planning/phases/01-base-configuration/01-CONTEXT.md` - Phase Context

## Context

### Phase Goal

CLI tool exists and functions with secure base configuration. This phase delivers the foundational CLI for step-by-step permission management with secure defaults.

### Requirements for Phase 1

- **INST-01**: OpenCode can be installed via brew
- **INST-02**: OpenCode can be installed via direct download
- **INST-03**: Installation guide is documented and easy to follow
- **BASE-01**: CLI command exists for configuration application
- **BASE-02**: Configuration uses secure defaults (no permissive rights)
- **BASE-03**: Auto-update is disabled by default
- **BASE-04**: External directory access is restricted
- **BASE-05**: Destructive commands (rm -rf) require confirmation
- **CLI-01**: A single command applies the configuration
- **CLI-02**: Command can gradually unlock permissions (phase by phase)

### Implementation Decisions from CONTEXT.md

- **D-01:** CLI command is `gsd configure --step N` where N is the phase number
- **D-02:** Flags: `--step N` for phase selection, `--list` for showing available steps, `--verify` for dry-run verification
- **D-03:** Output format: human-readable with color-coded status indicators
- **D-04:** Error handling: descriptive error messages with suggested fixes
- **D-05:** State file location: `.gsd/state.json`
- **D-06:** State tracks: current phase, applied permissions, configuration version
- **D-07:** Migration support: state file includes schema version for future updates
- **D-08:** Default permissions: restrictive — ask before read, deny by default for write operations
- **D-09:** Auto-update: disabled by default (`"autoupdate": false`)
- **D-10:** External directory access: restricted to explicitly configured paths only
- **D-11:** Destructive commands: require explicit confirmation (`rm -rf` prompts)
- **D-12:** Permission wildcards: expand to specific paths, never full home directory
- **D-13:** OpenCode installation: documented for both `brew install` and direct download
- **D-14:** Installation guide: step-by-step documentation in project root

## Project Instructions

Read `./AGENTS.md` if it exists in the working directory. Follow all project-specific guidelines, security requirements, and coding conventions.

## Output

Create PLAN.md files in `.planning/phases/01-base-configuration/`.

Each plan must include:
- Frontmatter with wave, depends_on, files_modified, autonomous
- Tasks in XML format with read_first and acceptance_criteria
- Verification criteria
- must_haves for goal-backward verification

## Quality Gate

Before returning PLANNING COMPLETE:
- [ ] PLAN.md files created in phase directory
- [ ] Each plan has valid frontmatter
- [ ] Tasks are specific and actionable
- [ ] Every task has `<read_first>` with at least the file being modified
- [ ] Every task has `<acceptance_criteria>` with grep-verifiable conditions
- [ ] Every `<action>` contains concrete values
- [ ] Dependencies correctly identified
- [ ] Waves assigned for parallel execution
- [ ] must_haves derived from phase goal
