# Phase 5: ospm alpha versionen per brew zur Verfügung stellen - Context

**Gathered:** 2026-04-07
**Status:** Ready for planning

<domain>
## Phase Boundary

Distribute ospm alpha versions via a separate Homebrew tap so users can install pre-release builds with `brew install` without affecting stable installation.

</domain>

<decisions>
## Implementation Decisions

### Alpha tap repository
- **D-01:** Separate tap repository: `homebrew-ospm-alpha`
- **D-02:** Full GitHub path: `github.com/openshell-policy-manager/homebrew-ospm-alpha`
- **D-03:** Tap is owned by `openshell-policy-manager` org

### Release automation
- **D-04:** Alpha releases triggered by git tag push (format: `v*.*.*-alpha.*`)
- **D-05:** GitHub Actions workflow builds and publishes on tag push
- **D-06:** goreleaser handles formula generation from template

### Packaging approach
- **D-07:** goreleaser + GitHub Releases for binary distribution
- **D-08:** Binaries attached to GitHub Release as artifacts
- **D-09:** goreleaser generates Formula from template and commits to tap repo

### Installation command
- **D-10:** Users install alpha via: `brew install openshell-policy-manager/ospm-alpha/ospm`

### Headless/formula naming
- **D-11:** Standard formula approach (not `--head` since we're distributing binaries)

</decisions>

<specifics>
## Specific Ideas

- "Use separate alpha tap" — keeps alpha isolated from stable `homebrew-ospm`
- "git tag push triggers release" — simple, familiar developer workflow

</specifics>

<canonical_refs>
## Canonical References

**Downstream agents MUST read these before planning or implementing.**

No external specs — requirements fully captured in decisions above.

</canonical_refs>

<code_context>
## Existing Code Insights

### Reusable Assets
- `.goreleaser.yaml` — Already configured with `brews` section; can be extended for alpha tap
- `.github/workflows/` — Existing GitHub Actions workflow patterns

### Established Patterns
- goreleaser `brews` formula: Uses `bin.install "ospm"` standard install
- GitHub Releases: Already used for distributing binaries

### Integration Points
- GitHub Actions: Trigger on tag push, run goreleaser
- Tap repository: Formula files committed by goreleaser post-build

</code_context>

<deferred>
## Deferred Ideas

None — discussion stayed within phase scope

</deferred>

---

*Phase: 5-ospm-alpha-versionen-per-brew-zur-verf-gung-stellen*
*Context gathered: 2026-04-07*