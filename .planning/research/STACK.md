# Stack Research: OpenShell Security Configuration macOS

**Domain:** Security configuration and sandbox for OpenShell on macOS
**Researched:** 2026-04-05
**Confidence:** MEDIUM

## Recommended Technology Stack

### Operating System Security Components

| Technology | Version | Purpose | Why Recommended |
|------------|---------|---------|-----------------|
| TCC (Transparency, Consent, and Control) | System-integrated | Access control for sensitive data (contacts, calendar, microphone, camera) | Apple's standard mechanism for privacy permissions; checked on every app launch |
| sandbox-exec | System-integrated (macOS 10.7+) | Command-line sandboxing | Built-in macOS feature for process isolation without third-party tools; uses Seatbelt sandbox profiles |
| SIP (System Integrity Protection) | System-integrated | Protection of critical system files | Prevents modification of system directories; essential for macOS security |
| Gatekeeper | System-integrated | App signature verification | Prevents execution of unsigned apps; first line of defense against malware |

### Shell and Terminal

| Technology | Version | Purpose | Why Recommended |
|------------|---------|---------|-----------------|
| zsh | 5.9+ (macOS Default) | Standard shell on macOS | Default since macOS Catalina; better plugin architecture than bash |
| Starship Prompt | 1.21+ | Cross-shell prompt optimization | 2026 Recommended alternative to Oh My Zsh; faster, less overhead, Rust-based |
| Ghostty | 1.3+ | Terminal emulator | Faster than iTerm2, Rust-based, modern architecture |

### Security Tools

| Technology | Version | Purpose | Why Recommended |
|------------|---------|---------|-----------------|
| tccutil | System-integrated | TCC database management | Apple's official CLI tool for resetting permissions; reset function only |
| tccplus (Fork) | Latest version | Extended TCC management | Fork of tccutil with Grant/Remove functions; enables finer control |
| stronghold | Latest | Automated macOS security settings | Community tool with 1.2k stars; configures firewall, Gatekeeper, SIP automatically |

### Network Security

| Technology | Version | Purpose | Why Recommended |
|------------|---------|---------|-----------------|
| Little Snitch | 6.x | Network monitoring and filtering | Industry standard for macOS network monitoring; blocks unwanted outbound connections |
| Murena / OpenSnitch | Latest | Linux-macOS network filtering | Open-source alternative to Little Snitch; less polished but functional |

## Installation

```bash
# Shell extensions
brew install starship
brew install ghostty

# Security tools
brew install tccplus  # or build from source
brew install stronghold

# Network monitoring (optional)
# Little Snitch is commercial, no Homebrew option
# OpenSnitch: brew install opensnitch
```

## Alternatives Considered

| Category | Recommended | Alternative | Why Not Recommended |
|----------|-------------|-------------|---------------------|
| Shell | zsh + Starship | Oh My Zsh + Powerlevel10k | Too slow for daily use in 2026; overhead from Ruby plugins |
| Terminal | Ghostty | iTerm2 | iTerm2 has tech debt; Rust-based tools are superior in performance |
| Prompt | Starship | Fig + Powerline | Starship is shell-agnostic and faster |
| TCC Management | tccplus | Manual SQL on TCC.db | Risky on updates; tccplus uses official APIs where possible |
| Sandbox | sandbox-exec | Docker containers | sandbox-exec is lighter and closer to the system; Docker overhead for simple isolation |

## What NOT to Use

| Avoid | Why | Use Instead |
|-------|-----|-------------|
| Direct TCC.db manipulation via SQL | Apple changes schema on updates; leads to corruption | tccplus or official Apple methods |
| Completely disabling SIP | Security risk; makes system vulnerable | Only temporarily disable for specific development tasks |
| Full root access without least-privilege principle | Violates Minimal Necessary Rights | Only unlock required rights step by step |
| Old macOS versions without current security patches | Known CVE exploits | Always use current macOS |

## Stack Variants by Use Case

**For developers with OpenCode/OpenShell:**
- Starship for fast prompts
- sandbox-exec for isolation of individual commands
- TCCplus for permission management
- No complete firewall blocking (networking must work for GitLab)

**For maximum security (Enterprise):**
- Little Snitch for full network control
- Strong sandbox profiles via sandbox-exec
- Complete TCC restrictions
- Regular security audits via scripts

## Version Compatibility

| Package | Compatible with | Notes |
|---------|-----------------|-------|
| Starship 1.21+ | zsh 5.0+, bash 4+, fish 3+ | Works with all common shells |
| Ghostty 1.3+ | macOS 13+ | Requires Apple Silicon or modern Intel |
| tccplus | macOS 11+ | Works with Big Sur and newer |
| stronghold | macOS 10.15+ | Catalina and newer supported |

## Confidence Assessment

| Area | Confidence | Rationale |
|------|------------|-----------|
| Shell/Terminal Tools | HIGH | Starship/Ghostty are current 2026 recommendations; verified through multiple sources |
| TCC Management | MEDIUM | tccplus is actively maintained fork; direct manipulation is risky |
| Sandbox Technologies | MEDIUM | sandbox-exec is system-integrated; however Apple changes APIs regularly |
| Network Security | MEDIUM | Little Snitch is proven; OpenSnitch less stable |

## Sources

- Apple Developer Documentation — sandbox-exec and entitlements
- HackTricks macOS TCC — current security mechanisms
- starship.rs — Cross-shell prompt (official docs)
- Ghostty GitHub — Terminal emulator
- Community/forums — tccplus and stronghold usage

---

*Stack research for OpenShell security configuration*
*Researched: 2026-04-05*
