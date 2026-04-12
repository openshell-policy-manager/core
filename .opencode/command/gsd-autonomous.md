---
description: Run all remaining phases autonomously — discuss→plan→execute per phase
<<<<<<< HEAD
argument-hint: "[--from N]"
=======
argument-hint: "[--from N] [--to N] [--only N] [--interactive]"
>>>>>>> 1763127 (chore: init gsd)
tools:
  read: true
  write: true
  bash: true
  glob: true
  grep: true
  question: true
  task: true
---
<objective>
Execute all remaining milestone phases autonomously. For each phase: discuss → plan → execute. Pauses only for user decisions (grey area acceptance, blockers, validation requests).

Uses ROADMAP.md phase discovery and Skill() flat invocations for each phase command. After all phases complete: milestone audit → complete → cleanup.

**Creates/Updates:**
- `.planning/STATE.md` — updated after each phase
- `.planning/ROADMAP.md` — progress updated after each phase
- Phase artifacts — CONTEXT.md, PLANs, SUMMARYs per phase

**After:** Milestone is complete and cleaned up.
</objective>

<execution_context>
@/Users/tkopatz/dev/work/tk-repos/openshell-policy-manager/.opencode/get-shit-done/workflows/autonomous.md
@/Users/tkopatz/dev/work/tk-repos/openshell-policy-manager/.opencode/get-shit-done/references/ui-brand.md
</execution_context>

<context>
<<<<<<< HEAD
Optional flag: `--from N` — start from phase N instead of the first incomplete phase.
=======
Optional flags:
- `--from N` — start from phase N instead of the first incomplete phase.
- `--to N` — stop after phase N completes (halt instead of advancing to next phase).
- `--only N` — execute only phase N (single-phase mode).
- `--interactive` — run discuss inline with questions (not auto-answered), then dispatch plan→execute as background agents. Keeps the main context lean while preserving user input on decisions.
>>>>>>> 1763127 (chore: init gsd)

Project context, phase list, and state are resolved inside the workflow using init commands (`gsd-tools.cjs init milestone-op`, `gsd-tools.cjs roadmap analyze`). No upfront context loading needed.
</context>

<process>
Execute the autonomous workflow from @/Users/tkopatz/dev/work/tk-repos/openshell-policy-manager/.opencode/get-shit-done/workflows/autonomous.md end-to-end.
Preserve all workflow gates (phase discovery, per-phase execution, blocker handling, progress display).
</process>
