---
description: Archive accumulated phase directories from completed milestones
<<<<<<< HEAD
=======
tools:
  read: true
  write: true
  bash: true
  question: true
>>>>>>> 1763127 (chore: init gsd)
---
<objective>
Archive phase directories from completed milestones into `.planning/milestones/v{X.Y}-phases/`.

Use when `.planning/phases/` has accumulated directories from past milestones.
</objective>

<execution_context>
@/Users/tkopatz/dev/work/tk-repos/openshell-policy-manager/.opencode/get-shit-done/workflows/cleanup.md
</execution_context>

<process>
Follow the cleanup workflow at @/Users/tkopatz/dev/work/tk-repos/openshell-policy-manager/.opencode/get-shit-done/workflows/cleanup.md.
Identify completed milestones, show a dry-run summary, and archive on confirmation.
</process>
