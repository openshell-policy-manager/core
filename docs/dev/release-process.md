# Release Process

## Repositories

| Repo | URL | Purpose |
|------|-----|---------|
| **Core** | `https://github.com/openshell-policy-manager/core` | Go source, binaries |
| **Homebrew Stable** | `https://github.com/openshell-policy-manager/homebrew-ospm` | Stable tap |
| **Homebrew Alpha** | `https://github.com/openshell-policy-manager/homebrew-ospm-alpha` | Alpha tap |
| **Presets** | `https://github.com/openshell-policy-manager/presets` | Policy presets |

## Release Workflow

```
GitLab (main) ──[Pipeline: sync_to_github]──> GitHub core (sync-from-gitlab branch)
                                                          │
                                                          ▼ Merge PR
                                        GitHub core (main) ──[git tag]──> Release
                                                                   │
                                                                   ▼
                                                          GoReleaser builds
                                                                   │
                                                      ┌────────────┴────────────┐
                                                      ▼                         ▼
                                              GitHub Releases         Homebrew Taps
                                              (binaries)         homebrew-ospm / homebrew-ospm-alpha
```

## Schritt-für-Schritt

### 1. Source in GitLab entwickeln

Alle Entwickung in GitLab:
- `https://gitlab.communicode.de/tkopatz/openshell-policy-manager`
- Branch: `main`

### 2. Zu GitHub synchronisieren

```bash
# In GitLab Pipeline UI:
# https://gitlab.communicode.de/tkopatz/openshell-policy-manager/-/pipelines
# → Run Pipeline (main branch)
# → Job: sync_to_github → Play Button
```

### 3. PR erstellen und mergen

Nach Sync erscheint Branch `sync-from-gitlab` in GitHub:
- https://github.com/openshell-policy-manager/core/branches
- PR erstellen und mergen

### 4. Release Tag erstellen

**Im GitHub core Repo:**
```bash
# Clone das repo
git clone https://github.com/openshell-policy-manager/core.git
cd core

# Stable Release
git tag -a v0.1.0 -m "Release v0.1.0"
git push origin v0.1.0

# Oder Alpha Release
git tag -a v0.1.0-alpha.1 -m "Alpha v0.1.0-alpha.1"
git push origin v0.1.0-alpha.1
```

**ODER via GitHub UI:**
- https://github.com/openshell-policy-manager/core/tags/new

### 5. Was passiert beim Tag

GitHub Actions (`.github/workflows/release.yml`) wird getriggert:

1. GoReleaser baut binaries (darwin/linux/windows, amd64/arm64)
2. Erstellt GitHub Release
3. Homebrew Cask wird zu Taps gepusht:
   - Stable → `homebrew-ospm`
   - Mit `-alpha*` oder `-beta*` im Tag → `homebrew-ospm-alpha`

## Installation

### Stable
```bash
brew install openshell-policy-manager/ospm/ospm
```

### Alpha
```bash
brew install openshell-policy-manager/ospm-alpha/ospm
```

### Manuell
```bash
go install github.com/openshell-policy-manager/core/cmd/ospm@latest
```

## Quick Reference

| Release | Tag | Brew Tap |
|---------|-----|----------|
| Stable | `v0.1.0` | `homebrew-ospm` |
| Alpha | `v0.1.0-alpha.1` | `homebrew-ospm-alpha` |
| Beta | `v0.2.0-beta.1` | `homebrew-ospm-alpha` |

## Wichtig

- **Core Repo:** `github.com/openshell-policy-manager/core` — hier builden die Binaries
- **Homebrew taps:** Werden automatisch von GoReleaser aktualisiert
- **Tag name bestimmt:** Ob stable oder alpha ( `-alpha*` oder `-beta*` = alpha tap)