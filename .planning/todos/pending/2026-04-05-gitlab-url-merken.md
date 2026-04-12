---
created: 2026-04-05T11:55:00Z
title: GitLab URL merken und nicht erneut abfragen
area: tooling
files:
  - bin/oshell-config-interactive
---

## Problem

Bei jedem Aufruf des oshell-config interactive Wizards muss die GitLab URL erneut eingegeben werden, auch wenn sie bereits vorher konfiguriert wurde. Dies ist umständlich für wiederholte Nutzung.

## Solution

- Beim Start des Wizards die gespeicherte Konfiguration aus .oshell-config/policy.json laden
- GitLab URL bereits vorab anzeigen und mit Enter bestätigen lassen ( statt neu eingeben)
- Wenn URL geändert werden soll, kann sie überschrieben werden
- Token nur abfragen wenn noch nicht konfiguriert oder geändert werden soll