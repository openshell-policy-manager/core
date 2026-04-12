# Development & Release Process

## Repositories

| Repo | URL | Purpose |
|------|-----|---------|
| **Core** | `https://github.com/openshell-policy-manager/core.git` | Go source code, CLI binary |
| **Presets** | `https://github.com/openshell-policy-manager/presets.git` | Policy YAML presets (brew, npm, pypi, docker, etc.) |
| **Homebrew Alpha** | `https://github.com/openshell-policy-manager/homebrew-ospm-alpha.git` | Homebrew tap for alpha releases |
| **Homebrew Stable** | `https://github.com/openshell-policy-manager/homebrew-ospm.git` | Homebrew tap for stable releases |

## GitLab (Development)

**URL:** `https://gitlab.communicode.de/tkopatz/openshell-policy-manager`

- Full source code including `.planning/` for roadmap and milestones
- Working branch: `main`

## GitHub (Release Distribution)

**URL:** `https://github.com/openshell-policy-manager/`

- Only release-relevant files synced from GitLab
- Files synced: `go.mod`, `go.sum`, `cmd/`, `internal/`, `.goreleaser.yaml`, `README.md`, `LICENSE`, `INSTALL.md`, `.github/workflows/`

## Workflow

```
GitLab (main) ──[Pipeline: sync_to_github]──> GitHub/core (main)
                                            │
                                            └──[Tag push]──> GitHub Actions (Release)
                                                              │
                                                              └── GoReleaser
                                                                  ├── Binaries to GitHub Releases
                                                                  └── Homebrew Cask to taps
```

## CI/CD Pipeline

**.gitlab-ci.yml** in GitLab:

- **Job:** `sync_to_github`
- **Trigger:** Manual (nicht automatisch)
- **Variables:**
  - `GORELEASER_TOKEN` — GitHub PAT with `repo` scope

## Release Process

### 1. Sync to GitHub (Manual)
```
GitLab → Pipeline → Run sync_to_github
```

### 2. Create Release Tag (in GitHub)
```bash
git tag -a v0.1.0 -m "Release v0.1.0"
git push origin v0.1.0
```

### 3. GitHub Actions Builds
- GoReleaser builds binaries (darwin/linux, amd64/arm64)
- Creates GitHub Release
- Updates Homebrew taps

## Installation

### Alpha (from Homebrew)
```bash
brew install openshell-policy-manager/ospm-alpha/ospm
```

### Stable (from Homebrew)
```bash
brew install openshell-policy-manager/ospm/ospm
```

### Manual
```bash
go build -o ospm ./cmd/ospm
```

## Links

- GitHub: https://github.com/openshell-policy-manager
- GitLab: https://gitlab.communicode.de/tkopatz/openshell-policy-manager