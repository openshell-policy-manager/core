# ospm

**OpenShell Policy Manager** — A secure CLI tool for managing NVIDIA OpenShell sandbox policies.

## What is OpenShell?

[OpenShell](https://build.nvidia.com/openshell) is NVIDIA's open-source security runtime for autonomous AI agents. It applies browser-like isolation principles to the agentic workflow:

- **Programmable Sandboxes** — Purpose-built isolation for agents that modify their own environment
- **Granular Policy Engine** — Controls what, where, and how agents can execute
- **Private Inference Router** — Keeps sensitive data on-device using local models

Policies are YAML-based and control filesystem access, network endpoints, and process execution. See [Policy Documentation](https://docs.nvidia.com/openshell/latest/sandboxes/policies.html) for details.

## ospm Features

- **Interactive Wizard** - Guided configuration with sensible defaults
- **Template System** - Minimal, dev, and full profiles
- **Preset Collection** - Reusable policies for common tools (brew, npm, pypi, docker, discord, slack, etc.)
- **Granular Options**:
  - apt-get update vs install (different paths)
  - npm install vs npx execute
  - git read vs write access
- **XDG Base Directory Support** - Follows standard conventions
- **Offline Support** - ZIP-based installation when GitHub unreachable
- **Default Storage** - Remembers last configuration


## Installation

[INSTALL.md](docs/user/INSTALL.md)

### Stable Release
```bash
brew install openshell-policy-manager/ospm/ospm
```

### Alpha Release
```bash
brew install openshell-policy-manager/ospm-alpha/ospm
```

### Manual
Download the latest release from the [releases page](https://github.com/openshell-policy-manager/ospm/releases).

### Build from Source
```bash
go build -o ospm ./cmd/ospm
```

## Usage

### Initialize
```bash
ospm init
```

### Start Wizard
```bash
ospm wizard
```

### Generate Policy
```bash
ospm generate --sandbox demo
```

### Apply Policy
```bash
ospm apply --sandbox demo
```

### Update Presets
```bash
ospm preset update
```

### List Templates
```bash
ospm template list
```

## Configuration

The tool follows XDG Base Directory Specification:

- **Global Config**: `~/.config/ospm/`
- **Credentials**: `~/.config/opencode/credentials/`
- **OpenShell Policies**: `~/.ospm/`

### Hierarchy
Configuration priority (highest to lowest):
1. Project-local config (`.ospm/` in project root)
2. Global config (`~/.config/ospm/`)
3. Built-in defaults

## Environment Variables

| Variable | Description |
|----------|-------------|
| `OSPM_PRESETS_URL` | Override preset GitHub URL |
| `OSPM_PRESETS_REPO` | Custom preset repository URL |
| `OSPM_OFFLINE_ZIP` | Path to offline ZIP file |
| `OSPM_LOCALE` | Language (en, de) |
| `XDG_CONFIG_HOME` | Override config directory |

## Development

```bash
# Install dependencies
go mod download

# Build
go build -o bin/ospm ./cmd/ospm

# Run tests
go test ./...

# Release with goreleaser
goreleaser build --snapshot --clean
```

## Project Structure

```
├── cmd/                   # CLI entry point
│   └── ospm/              # Main application
├── internal/
│   ├── config/            # Config, template, policy types
│   ├── wizard/            # Interactive wizard
│   └── i18n/              # Internationalization
├── .goreleaser.yaml       # Build configuration
└── go.mod
```

## License

MIT License - see LICENSE file for details

## Contributing

Contributions are welcome! Please open an issue or submit a PR.

## Related

- [OpenShell Documentation](https://docs.nvidia.com/openshell/latest/index.html)
- [OpenShell on NVIDIA Build](https://build.nvidia.com/openshell)
- [NemoClaw Blueprint Presets](https://github.com/NVIDIA/NemoClaw/tree/main/nemoclaw-blueprint/policies/presets)
