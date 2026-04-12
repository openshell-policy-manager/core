# OpenShell Preset Repository Struktur

Dieses Dokument beschreibt die vorgeschlagene Struktur für das OpenShell Preset Repository unter `https://github.com/openshell-policy-manager/presets`.

## Überblick

Das Preset-Repository enthält wiederverwendbare Sicherheitsrichtlinien für OpenShell. Jedes Preset definiert einen Satz von Netzwerk-, Dateisystem- und Prozessrichtlinien, die einfach kombiniert werden können.

## Verzeichnisstruktur

```
presets/
├── blueprint.yaml              # Haupt-Blueprint (Konfiguration)
├── network/                   # Netzwerk-Policies (einzelne Dienste)
│   ├── gitlab.yaml
│   ├── github.yaml
│   ├── npm.yaml
│   ├── pypi.yaml
│   ├── docker.yaml
│   ├── apt.yaml
│   └── exa-ai.yaml
├── filesystem/                # Filesystem-Policies
│   ├── read-only.yaml
│   └── read-write.yaml
└── profiles/                  # Komplette Profile (Policy-Kombinationen)
    ├── minimal.yaml           # Git Read Only
    ├── dev.yaml               # npm + websearch
    └── full.yaml              # Alles inklusive
```

## Policy-Schema

### Netzwerk-Policy (`network/*.yaml`)

```yaml
name: "<preset-name>"
description: "<beschreibung>"
version: "0.1.0"

network:
  allowed:
    - host: "<hostname>"
      port: <port>
      tls: true|false

binaries:
  - "<pfad-zu-binär>"
```

### Komplettes Profil (`profiles/*.yaml`)

```yaml
name: "<profil-name>"
description: "<beschreibung>"
version: "0.1.0"

includes:
  - network/gitlab
  - network/npm
  - filesystem/read-write

network:
  allowed: []
  
binaries: []
```

## Bestehende Presets

### Netzwerk-Policies

| Preset | Host(s) | Binaries |
|--------|---------|----------|
| `gitlab` | gitlab.com:443 | git |
| `github` | github.com:443 | git |
| `npm` | registry.npmjs.org:443, api.npmjs.org:443 | npm, npx, node |
| `pypi` | pypi.org:443, files.pythonhosted.org:443 | pip, pip3, python3 |
| `docker` | registry-1.docker.io:443, auth.docker.io:443 | docker |
| `apt` | security.ubuntu.com:443, archive.ubuntu.com:443 | apt, apt-get |

### Profile

| Profil | Enthaltene Presets | Anwendungsfall |
|--------|-------------------|----------------|
| `minimal` | github, gitlab | Nur Git-Read-Zugriff |
| `dev` | npm, exa-ai | Frontend-Entwicklung |
| `full` | npm, pypi, docker, apt | Vollständige Entwicklung |

## Verwendung

### Presets referenzieren

In einem Blueprint können Presets eingebunden werden:

```yaml
components:
  policy:
    base: "policies/openshell-sandbox.yaml"
    additions:
      dev: "policies/presets/dev.yaml"
```

### Eigenes Profil erstellen

1. Netzwerk-Policies aus `network/` kombinieren
2. Filesystem-Policies aus `filesystem/` hinzufügen
3. Als neues Profil in `profiles/` speichern

## Entwicklung

### Neues Preset hinzufügen

1. YAML-Datei in entsprechendem Verzeichnis erstellen
2. Schema oben befolgen
3. Version bumpen (semantisch)
4. PR erstellen

### Validierung

```bash
# YAML-Syntax prüfen
yamllint presets/

# Blueprint validieren
openshell validate blueprint.yaml
```

## Lizenz

Siehe LICENSE Datei im Repository.
