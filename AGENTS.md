<!-- GSD:project-start source:PROJECT.md -->
## Project

**OpenShell Configuration**

Eine sichere OpenShell-Konfiguration für die lokale Entwicklung mit OpenCode und einem Remote GitLab-Repository. Die Konfiguration wird schrittweise durch einen CLI-Befehl angewendet, um nach und nach Rechte und Zugriffe freizuschalten.

**Core Value:** Sichere Entwicklungsumgebung: OpenCode kann alle erforderlichen Operationen ausführen, aber nur mit den minimal notwendigen Rechten und Zugriffen.

### Constraints

- **Netzwerk:** Nur Daten herunterladen die vorher explizit freigegeben wurden
- **Skills:** Nur Skills und MCP-Server die in OpenCode freigegeben sind
- **CLI:** Einfacher Befehl zur schrittweisen Konfigurationsanwendung
- **Dev Tools:** Nur OpenCode (keine volle Dev-Umgebung)
<!-- GSD:project-end -->

<!-- GSD:stack-start source:research/STACK.md -->
## Technology Stack

## Empfohlener Technologie-Stack
### Betriebssystem-Sicherheitskomponenten
| Technologie | Version | Zweck | Warum empfohlen |
|-------------|---------|-------|-----------------|
| TCC (Transparency, Consent, and Control) | Systemintegriert | Zugriffskontrolle auf sensible Daten (Kontakte, Kalender, Mikrofon, Kamera) | Apple's Standard-Mechanismus für Datenschutzberechtigungen; wird bei jedem App-Start geprüft |
| sandbox-exec | Systemintegriert (macOS 10.7+) | Kommandozeilen-Sandboxing | Eingebaute macOS-Funktion für Prozess-Isolation ohne Drittanbieter-Tools; nutzt Seatbelt-Sandbox-Profile |
| SIP (System Integrity Protection) | Systemintegriert | Schutz kritischer Systemdateien | Verhindert Modifikation von Systemverzeichnissen; essentiell für macOS-Sicherheit |
| Gatekeeper | Systemintegriert | App-Signatur-Prüfung | Verhindert Ausführung nicht signierter Apps; erster Schutzwall gegen Malware |
### Shell und Terminal
| Technologie | Version | Zweck | Warum empfohlen |
|-------------|---------|-------|-----------------|
| zsh | 5.9+ (macOS Default) | Standard-Shell auf macOS | Seit macOS Catalina der Default; bessere Plugin-Architektur als bash |
| Starship Prompt | 1.21+ | Cross-Shell Prompt-Optimierung | 2026 Recommended Alternative zu Oh My Zsh; schneller, weniger Overhead, Rust-basiert |
| Ghostty | 1.3+ | Terminal-Emulator | Schneller als iTerm2, Rust-basiert, moderne Architektur |
### Sicherheitswerkzeuge
| Technologie | Version | Zweck | Warum empfohlen |
|-------------|---------|-------|-----------------|
| tccutil | Systemintegriert | TCC-Datenbank-Verwaltung | Apple's offizielles CLI-Tool zum Zurücksetzen von Berechtigungen; nur Reset-Funktion |
| tccplus (Fork) | Aktuellste Version | Erweiterte TCC-Verwaltung | Fork von tccutil mit Grant/Remove-Funktionen; ermöglicht feinere Kontrolle |
| stronghold | Latest | Automatisierte macOS-Sicherheitseinstellungen | Community-Tool mit 1.2k Stars; konfiguriert Firewall, Gatekeeper, SIP automatisch |
### Netzwerk-Sicherheit
| Technologie | Version | Zweck | Warum empfohlen |
|-------------|---------|-------|-----------------|
| Little Snitch | 6.x | Netzwerk-Monitoring und -Filterung | Industry Standard für macOS-Netzwerküberwachung; blockiert unerwünschte Outbound-Verbindungen |
| Murena / OpenSnitch | Latest | Linux-macOS Netzwerk-Filterung | Open-Source-Alternative zu Little Snitch; weniger polished aber funktional |
## Installation
# Shell-Erweiterungen
# Sicherheits-Tools
# Netzwerk-Monitoring (optional)
# Little Snitch ist kommerziell, keine Homebrew-Option
# OpenSnitch: brew install opensnitch
## Alternativen betrachtet
| Kategorie | Empfohlen | Alternative | Warum nicht empfohlen |
|-----------|-----------|-------------|----------------------|
| Shell | zsh + Starship | Oh My Zsh + Powerlevel10k | 2026 zu langsam für täglichen Gebrauch; Overhead durchRuby-Plugins |
| Terminal | Ghostty | iTerm2 | iTerm2 hat tech debt; Rust-basierte Tools performancetechnisch überlegen |
| Prompt | Starship | Fig + Powerline | Starship ist shell-agnostisch und schneller |
| TCC-Management | tccplus | manuelles SQL auf TCC.db | Riskant bei Updates; tccplus nutzt offizielle APIs wo möglich |
| Sandbox | sandbox-exec | Docker-Container | sandbox-exec ist leichter und systemnäher; Docker overhead für einfache Isolation |
## Was NICHT verwenden
| Vermeiden | Warum | Stattdessen verwenden |
|-----------|-------|----------------------|
| Direkte TCC.db Manipulation via SQL | Apple ändert Schema bei Updates; führt zu Korruption | tccplus oder offizielle Apple-Methoden |
| SIP vollständig deaktivieren | Sicherheitsrisiko; macht System angreifbar | Nur für spezifische Entwicklungsaufgaben temporär deaktivieren |
| Vollständige Root-Zugriffe ohne Grundprinzip | Minimal Necessary Rights verletzt | Nur benötigte Rechte schrittweise freischalten |
| Alte macOS-Versionen ohne aktuelle Sicherheitspatches | Bekannte CVE-Exploits | Immer aktuelles macOS verwenden |
## Stack-Varianten nach Anwendungsfall
- Starship für schnelle Prompts
- sandbox-exec für Isolation einzelner Befehle
- TCCplus für Berechtigungsmanagement
- Keine vollständige Firewall-Blockierung (Networking muss für GitLab funktionieren)
- Little Snitch für vollständige Netzwerkkontrolle
- Starke Sandbox-Profile via sandbox-exec
- Vollständige TCC-Beschränkungen
- regelmäßige Sicherheitsaudits via Scripts
## Versions-Kompatibilität
| Paket | Kompatibel mit | Hinweise |
|-------|----------------|----------|
| Starship 1.21+ | zsh 5.0+, bash 4+, fish 3+ | Funktioniert mit allen gängigen Shells |
| Ghostty 1.3+ | macOS 13+ | Benötigt Apple Silicon oder moderne Intel |
| tccplus | macOS 11+ | Funktioniert mit Big Sur und neuer |
| stronghold | macOS 10.15+ | Catalina und neuer unterstützt |
## Konfidenz-Bewertung
| Bereich | Konfidenz | Begründung |
|---------|-----------|------------|
| Shell/Terminal Tools | HOCH | Starship/Ghostty sind aktuelle 2026-Empfehlungen; verifiziert durch mehrere Quellen |
| TCC-Management | MEDIUM | tccplus ist aktiv gepflegter Fork; direkte Manipulation ist riskant |
| Sandbox-Technologien | MEDIUM | sandbox-exec ist systemintegriert; Apple ändert jedoch regelmäßig APIs |
| Netzwerk-Security | MEDIUM | Little Snitch ist bewährt; OpenSnitch weniger stabil |
## Quellen
- Apple Developer Documentation — sandbox-exec und entitlements
- HackTricks macOS TCC — aktuelle Sicherheitsmechanismen
- starship.rs — Cross-Shell Prompt (offizielle Docs)
- Ghostty GitHub — Terminal-Emulator
- community/foren — tccplus und stronghold Nutzung
<!-- GSD:stack-end -->

<!-- GSD:conventions-start source:CONVENTIONS.md -->
## Conventions

Conventions not yet established. Will populate as patterns emerge during development.
<!-- GSD:conventions-end -->

<!-- GSD:architecture-start source:ARCHITECTURE.md -->
## Architecture

Architecture not yet mapped. Follow existing patterns found in the codebase.
<!-- GSD:architecture-end -->

<!-- GSD:skills-start source:skills/ -->
## Project Skills

No project skills found. Add skills to any of: `.claude/skills/`, `.agents/skills/`, `.cursor/skills/`, or `.github/skills/` with a `SKILL.md` index file.
<!-- GSD:skills-end -->

<!-- GSD:workflow-start source:GSD defaults -->
## GSD Workflow Enforcement

Before using Edit, Write, or other file-changing tools, start work through a GSD command so planning artifacts and execution context stay in sync.

Use these entry points:
- `/gsd-quick` for small fixes, doc updates, and ad-hoc tasks
- `/gsd-debug` for investigation and bug fixing
- `/gsd-execute-phase` for planned phase work

Do not make direct repo edits outside a GSD workflow unless the user explicitly asks to bypass it.
<!-- GSD:workflow-end -->



<!-- GSD:profile-start -->
## Developer Profile

> Profile not yet configured. Run `/gsd-profile-user` to generate your developer profile.
> This section is managed by `generate-claude-profile` -- do not edit manually.
<!-- GSD:profile-end -->
