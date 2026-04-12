# Plan 05-01 Summary

**Objective:** Configure alpha release automation with goreleaser + GitHub Actions
**Status:** ✅ Complete

## Tasks Executed

| # | Task | Status |
|---|------|--------|
| 1 | Extend goreleaser config for alpha tap | ✅ |
| 2 | Create GitHub Actions workflow for alpha releases | ✅ |

## Files Modified/Created

- `.goreleaser.yaml` — Added alpha brews section with `ospm-alpha` id
- `.github/workflows/release-alpha.yaml` — New workflow, triggers on `v*.*.*-alpha.*` tags

## Setup Required

Before first alpha release, user must:
1. Create empty `homebrew-ospm-alpha` repository on GitHub
2. Generate classic PAT with `repo` scope at github.com/settings/tokens
3. Add as secret `TAP_GITHUB_TOKEN` in repo settings

## Usage

```bash
# Tag an alpha release
git tag v0.1.0-alpha.1 && git push

# Users install alpha
brew install openshell-policy-manager/ospm-alpha/ospm
```

---
status: complete