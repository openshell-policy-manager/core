# Roadmap: OpenShell Configuration

## Milestones

- ✅ **v1.0 MVP** — Phases 1-4 (shipped 2026-04-05) — [Archive](milestones/v1.0-ROADMAP.md)
- ◆ **v1.1 Alpha Releases** — Phase 5 (in progress)

## Phases

<details>
<summary>✅ v1.0 MVP (Phases 1-4) — SHIPPED 2026-04-05</summary>

- [x] Phase 1: Base Configuration (1/1 plans) — completed 2026-04-05
- [x] Phase 2: GitLab Integration (1/1 plans) — completed 2026-04-05
- [x] Phase 3: Skills & Sandbox (1/1 plans) — completed 2026-04-05
- [x] Phase 4: Verification & Polish (1/1 plans) — completed 2026-04-05

</details>

<details>
<summary>◆ v1.1 Alpha Releases (Phase 5)</summary>

- [x] Phase 5: OSPM Alpha-Versionen per Brew zur Verfügung stellen (0/0 plans) — completed 2026-04-07
  - Alpha release automation configured (goreleaser + GitHub Actions)
  - Verification passed — manual setup required for first release

</details>

## Progress

| Phase | Milestone | Plans Complete | Status | Completed |
|-------|-----------|----------------|--------|-----------|
| 1. Base Configuration | v1.0 | 1/1 | Complete | 2026-04-05 |
| 2. GitLab Integration | v1.0 | 1/1 | Complete | 2026-04-05 |
| 3. Skills & Sandbox | v1.0 | 1/1 | Complete | 2026-04-05 |
| 4. Verification & Polish | v1.0 | 1/1 | Complete | 2026-04-05 |
| 5. OSPM Alpha-Versionen per Brew | v1.1 | 0/0 | Complete | 2026-04-07 |

---

## Backlog

### Phase 999.1: OSPM Reset löscht Custom Presets (BACKLOG)

**Goal:** ospm reset soll die bisher angelegten custom presets wieder auf den auslieferungszustand bringen, also löschen. Es sollen aber nur die Dateien in den config-Ordnern gelöscht werden die nicht zu ospm gehören. Andere Konfigurationen sollen unangetastet bleiben.
**Requirements:** TBD
**Plans:** 0 plans

Plans:
- [ ] TBD (promote with /gsd-review-backlog when ready)

### Phase 999.2: Warning wenn OpenShell Gateway nicht verfügbar (BACKLOG)

**Goal:** Warning wenn openshell gateway nicht verfügbar
**Requirements:** TBD
**Plans:** 0 plans

Plans:
- [ ] TBD (promote with /gsd-review-backlog when ready)

### Phase 999.3: Brew Channel Updater

**Goal:** Automated Homebrew formula/channel updates for OSPM releases. Keep the Homebrew tap in sync with GitHub releases using GoReleaser.
**Requirements:** TBD
**Plans:** 1/1 plans complete

Plans:
- [x] 999.3-01-PLAN.md — Dual tap support (alpha + stable)

### Phase 999.4: Fix Build and Release (BACKLOG)

**Goal:** Fix build and release errors so alpha releases work with Homebrew
**Requirements:** TBD
**Plans:** 1/1 plans complete

Plans:
- [ ] TBD (promote with /gsd-review-backlog when ready)

---

*Roadmap updated: 2026-04-12 — Phase 999.4 added to backlog*
