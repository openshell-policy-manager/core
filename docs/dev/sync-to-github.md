# Sync to GitHub on tag push

This workflow syncs only release-relevant files to GitHub when a tag is pushed in GitLab.

## Files to sync (keep in GitHub)

```
# Go build
go.mod
go.sum
cmd/
internal/

# Release config
.goreleaser.yaml

# Docs
README.md
LICENSE
INSTALL.md
```

## Files to exclude (stay in GitLab only)

```
.planning/
.github/
bin/
docs/
*.md (except README, LICENSE, INSTALL)
.gitignore
```

## Setup

### 1. Create GitHub Repository

https://github.com/openshell-policy-manager/ospm

### 2. Create GitHub PAT

Create a Personal Access Token on GitHub with `repo` scope.

### 3. Add GitLab CI/CD Variable

In GitLab (`gitlab.communicode.de/openshell-policy/ospm`):
- Settings → CI/CD → Variables
- Add variable: `GITHUB_TOKEN` = your PAT

### 4. Add .gitlab-ci.yml

Create `.gitlab-ci.yml` in the GitLab repo:

```yaml
sync_to_github:
  image: alpine/git:latest
  only:
    - tags
  script:
    - |
      # Clone GitHub repo
      git clone https://openshell-policy-manager:$GITHUB_TOKEN@github.com/openshell-policy-manager/ospm.git /tmp/ospm-github
      cd /tmp/ospm-github
      
      # Copy release files from GitLab source
      cp -r ${CI_PROJECT_DIR}/go.mod .
      cp -r ${CI_PROJECT_DIR}/go.sum .
      cp -r ${CI_PROJECT_DIR}/cmd .
      cp -r ${CI_PROJECT_DIR}/internal .
      cp ${CI_PROJECT_DIR}/.goreleaser.yaml .
      cp ${CI_PROJECT_DIR}/README.md .
      cp ${CI_PROJECT_DIR}/LICENSE .
      cp ${CI_PROJECT_DIR}/INSTALL.md .
      
      # Commit and push
      git add .
      git config --local user.name "GitLab CI"
      git config --local user.email "ci@gitlab.communicode.de"
      git commit -m "Sync from GitLab: ${CI_COMMIT_TAG}" || true
      git push origin main
  when: manual

> **Empfohlen:** `when: manual` — du startest den Sync bewusst bevor du den eigentlichen Release baust.
```

**Alternative: Automatic sync on tag**

```yaml
sync_to_github:
  image: alpine/git:latest
  only:
    - tags
  script:
    - |
      git clone https://openshell-policy-manager:$GITHUB_TOKEN@github.com/openshell-policy-manager/ospm.git /tmp/ospm-github
      cd /tmp/ospm-github
      
      # Sync files
      cp ${CI_PROJECT_DIR}/go.mod .
      cp ${CI_PROJECT_DIR}/go.sum .
      cp -r ${CI_PROJECT_DIR}/cmd .
      cp -r ${CI_PROJECT_DIR}/internal .
      cp ${CI_PROJECT_DIR}/.goreleaser.yaml .
      cp ${CI_PROJECT_DIR}/README.md .
      cp ${CI_PROJECT_DIR}/LICENSE .
      cp ${CI_PROJECT_DIR}/INSTALL.md .
      cp ${CI_PROJECT_DIR}/.github/workflows/release.yml .
      
      git add .
      git config --local user.name "GitLab CI"
      git config --local user.email "ci@gitlab.communicode.de"
      git commit -m "Sync from GitLab: ${CI_COMMIT_TAG}"
      git push origin main
```

## Workflow

```
GitLab (full repo) ──[tag push]──> GitHub (release files) ──[GitHub Actions]──> Release
```

1. **Developer** works in GitLab (all files including `.planning/`)
2. **Tag push** in GitLab triggers CI/CD
3. **CI/CD** copies only release files to GitHub
4. **GitHub Actions** builds and releases

## Notes

- GitHub Actions workflow stays in `.github/workflows/release.yml`
- Keep it synced to GitHub manually or in step 3
- Manual trigger in GitLab gives you control before sync