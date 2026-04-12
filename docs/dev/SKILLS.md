# OpenCode Skills - Konfiguration

## Übersicht

Diese Dokumentation beschreibt die Skills, die für die OpenShell-Konfiguration aktiviert sind.

## Aktivierte Skills

### websearch / webfetch
- **Zweck:** Internet-Recherche
- **Berechtigungen:** network
- **Verwendung:** Exa-Search, Live-Web-Crawling

### codesearch
- **Zweck:** Code-Analyse
- **Berechtigungen:** read
- **Verwendung:** Exa Code API für Programmierfragen

### file_operations
- **Zweck:** Dateizugriffe
- **Berechtigungen:** read, write, glob, grep
- **Verwendung:** Projektdateien lesen und schreiben

### git_operations
- **Zweck:** Git-Repository-Interaktionen
- **Berechtigungen:** status, diff, commit, push, pull
- **Verwendung:** Git-Operationen via OpenCode

## Konfiguration

Skills werden in `.planning/config.json` unter `agent_skills` definiert.

## Sicherheit

Nur explizit benötigte Skills sind aktiviert (Minimalprinzip).
