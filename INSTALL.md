# OSPM Installation

## Prerequisites

- macOS or Linux
- Git

## Stable Release

```bash
brew install openshell-policy-manager/ospm/ospm
```

## Alpha Release

```bash
brew install openshell-policy-manager/ospm-alpha/ospm
```

## Manual

Download the latest release from the [releases page](https://github.com/openshell-policy-manager/core/releases).

## Build from Source

```bash
git clone https://github.com/openshell-policy-manager/core.git
cd core
go build -o ospm ./cmd/ospm
```

## Usage

```bash
ospm init
ospm wizard
ospm generate --sandbox demo
```