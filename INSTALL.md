# OpenShell Configuration - Installation Guide

This guide explains how to install OpenCode and set up the GSD (Get Shit Done) CLI tool for secure, step-by-step configuration.

## Prerequisites

- macOS (tested on macOS 12+)
- Terminal access
- Git (for cloning this repository)

## Installation Methods

### Method 1: Install via Homebrew (Recommended)

```bash
# Add OpenCode to Homebrew (if available)
brew install opencode

# Or install from source
git clone https://github.com/opencode-ai/opencode.git
cd opencode
make install
```

### Method 2: Direct Download

```bash
# Download the latest release
curl -L -o /tmp/opencode.zip https://github.com/opencode-ai/opencode/releases/latest/download/opencode-darwin-arm64.zip

# Extract to a location in your PATH
unzip /tmp/opencode.zip -d ~/Applications/opencode

# Add to PATH (add to ~/.zshrc or ~/.bashrc)
export PATH="$HOME/Applications/opencode:$PATH"

# Verify installation
opencode --version
```

## GSD CLI Setup

### Step 1: Clone this Repository

```bash
git clone <repository-url> openshell-config
cd openshell-config
```

### Step 2: Add GSD to Your PATH

You can either:

**Option A: Add to PATH (Recommended)**
Add the following to your `~/.zshrc` (or `~/.bashrc`):

```bash
# Add GSD CLI to PATH
export PATH="$HOME/openshell-config/bin:$PATH"
```

Then reload your shell:
```bash
source ~/.zshrc
```

**Option B: Use Absolute Path**
You can also run the CLI directly:
```bash
./bin/oshell-config --help
```

### Step 3: Verify Installation

```bash
# Show help
oshell-config --help

# Show available configuration steps
oshell-config configure --list

# Show current status
oshell-config status
```

## Quick Start

### Apply Your First Configuration Step

```bash
oshell-config configure --step 1
```

This will:
- Set up the initial state file
- Configure basic security defaults
- Disable auto-update by default

### Check Configuration Status

```bash
oshell-config status
```

### Dry-Run Verification

Before applying any changes, you can test without making modifications:

```bash
oshell-config configure --step 2 --verify
```

## Configuration Steps

| Step | Description | What It Does |
|------|-------------|--------------|
| 1 | Base Configuration | Install OpenCode, basic security defaults |
| 2 | GitLab Integration | Network access, OAuth configuration |
| 3 | Skills & Sandbox | OpenCode skills, MCP servers |
| 4 | Verification & Polish | Status display, dry-run mode |

### Step 1: Base Configuration

```bash
oshell-config configure --step 1
```

Applies:
- OpenCode installation instructions
- Basic security defaults (restrictive permissions)
- Auto-update disabled by default
- External directory access restricted
- Destructive commands require confirmation

### Step 2: GitLab Integration

```bash
oshell-config configure --step 2
```

Applies:
- GitLab OAuth token configuration
- Network access to gitlab.com
- Credential storage setup

### Step 3: Skills & Sandbox

```bash
oshell-config configure --step 3
```

Applies:
- OpenCode skills configuration
- MCP server setup
- Websearch/Webfetch enablement

### Step 4: Verification & Polish

```bash
oshell-config configure --step 4
```

Applies:
- Status display functionality
- Dry-run verification mode
- Export/Import configuration

## Security Features

The GSD CLI implements these security best practices:

1. **Secure Defaults**: All configurations start restrictive
2. **Auto-Update Disabled**: Updates must be explicitly enabled
3. **Directory Restrictions**: External access limited to configured paths
4. **Destructive Command Protection**: `rm -rf` requires confirmation
5. **Permission Wildcards**: Never expand to full home directory
6. **State File Permissions**: Set to 600 (owner only)

## Troubleshooting

### Command Not Found

Make sure the bin directory is in your PATH:
```bash
echo $PATH | grep -q "openshell-config/bin" && echo "PATH OK" || echo "Add to PATH"
```

### Permission Denied

Make the scripts executable:
```bash
chmod +x bin/oshell-config bin/oshell-config-configure
```

### State File Issues

If you need to reset:
```bash
rm .oshell-config/state.json
oshell-config configure --step 1
```

## Further Reading

- See `.planning/PROJECT.md` for project overview
- See `.planning/ROADMAP.md` for phase roadmap
- See `.planning/REQUIREMENTS.md` for requirements specification

## Support

For issues or questions:
- Check the project documentation
- Review the CLI help: `oshell-config --help`
- Run with `--verify` flag to test without applying changes