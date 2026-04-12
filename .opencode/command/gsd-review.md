---
description: Request cross-AI peer review of phase plans from external AI CLIs
<<<<<<< HEAD
argument-hint: "--phase N [--gemini] [--claude] [--codex] [--all]"
=======
argument-hint: "--phase N [--gemini] [--claude] [--codex] [--opencode] [--all]"
>>>>>>> 1763127 (chore: init gsd)
tools:
  read: true
  write: true
  bash: true
  glob: true
  grep: true
---

<objective>
<<<<<<< HEAD
Invoke external AI CLIs (Gemini, the agent, Codex) to independently review phase plans.
=======
Invoke external AI CLIs (Gemini, the agent, Codex, OpenCode) to independently review phase plans.
>>>>>>> 1763127 (chore: init gsd)
Produces a structured REVIEWS.md with per-reviewer feedback that can be fed back into
planning via /gsd-plan-phase --reviews.

**Flow:** Detect CLIs → Build review prompt → Invoke each CLI → Collect responses → Write REVIEWS.md
</objective>

<execution_context>
@/Users/tkopatz/dev/work/tk-repos/openshell-policy-manager/.opencode/get-shit-done/workflows/review.md
</execution_context>

<context>
Phase number: extracted from $ARGUMENTS (required)

**Flags:**
- `--gemini` — Include Gemini CLI review
- `--claude` — Include the agent CLI review (uses separate session)
- `--codex` — Include Codex CLI review
<<<<<<< HEAD
=======
- `--opencode` — Include OpenCode review (uses model from user's OpenCode config)
>>>>>>> 1763127 (chore: init gsd)
- `--all` — Include all available CLIs
</context>

<process>
Execute the review workflow from @/Users/tkopatz/dev/work/tk-repos/openshell-policy-manager/.opencode/get-shit-done/workflows/review.md end-to-end.
</process>
