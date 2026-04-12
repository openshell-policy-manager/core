# Phase 999.1: ospm alpha versionen per brew zur Verfügung stellen - Discussion Log

> **Audit trail only.** Do not use as input to planning, research, or execution agents.
> Decisions captured in CONTEXT.md — this log preserves the alternatives considered.

**Date:** 2026-04-07
**Phase:** 999.1-ospm-alpha-versionen-per-brew-zur-verf-gung-stellen
**Mode:** discuss
**Areas discussed:** Alpha tap repository, Release automation, Packaging approach

---

## Alpha tap repository

| Option | Description | Selected |
|--------|-------------|----------|
| Existing tap + pre-release tags | Use the existing tap. Alpha versions are distributed via pre-release GitHub tags. Users would `brew install openshell-policy-manager/ospm/ospm --head` or similar. | |
| Separate alpha tap | Create a separate tap like `openshell-policy-manager/homebrew-ospm-alpha`. Keeps alpha separated from stable. | ✓ |
| Custom URL scheme | Use a custom brew installer URL pointing directly to GitHub releases. | |

**User's choice:** Separate alpha tap
**Notes:** keeps alpha isolated from stable `homebrew-ospm`

---

## Release automation

| Option | Description | Selected |
|--------|-------------|----------|
| Git tag push (Recommended) | Tag like v0.1.0-alpha.1 triggers build → brew update → users get it | ✓ |
| GitHub Actions workflow dispatch | Manual workflow dispatch with version input field | |
| Branch push (alpha branch) | Build whenever alpha branch has new commits | |

**User's choice:** Git tag push (Recommended)
**Notes:** simple, familiar developer workflow

---

## Alpha brew tap

| Option | Description | Selected |
|--------|-------------|----------|
| homebrew-ospm-alpha (Recommended) | github.com/openshell-policy-manager/homebrew-ospm-alpha | ✓ |
| Separate repo (homebrew-ospm) | github.com/openshell-policy-manager/homebrew-ospm | |
| Same repo, different branch | github.com/openshell-policy-manager/homebrew-ospm.git#alpha | |

**User's choice:** homebrew-ospm-alpha (Recommended)

---

## Packaging approach

| Option | Description | Selected |
|--------|-------------|----------|
| goreleaser + GitHub Releases (Recommended) | Build binaries → attach to GitHub Release → goreleaser creates the Formula from template → commits to tap repo | ✓ |
| goreleaser + bottles | goreleaser generates bottles (pre-built binaries) → uploads to bintray or GitHub Releases → serves to brew | |
| Brew from source | Brew formula pulls source and builds on install (slower but always fresh) | |

**User's choice:** goreleaser + GitHub Releases (Recommended)
**Notes:** Uses existing `.goreleaser.yaml` pattern

---

## User's Discretion

No areas deferred to agent discretion.

---

## Deferred Ideas

None — discussion stayed within phase scope