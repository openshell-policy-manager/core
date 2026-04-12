# Architecture Patterns: OpenShell Security Configuration

**Domain:** macOS Security Configuration System
**Researched:** 2026-04-05
**Confidence:** MEDIUM

## Executive Summary

OpenShell security configuration on macOS follows a layered architecture pattern seen in tools like **stronghold** (CLI-based), **NIST mSCP** (compliance-focused), and **MACE** (visual editor). The recommended architecture separates configuration definition from execution, enabling the "step-by-step" deployment model required by this project. Three layers dominate: **Definition** (YAML/JSON configs), **Engine** (CLI + execution logic), and **Enforcement** (system APIs: defaults, TCC, pfctl).

---

## Recommended Architecture

### Component Boundaries

```
┌─────────────────────────────────────────────────────────────────┐
│                     USER INTERFACE                              │
│  CLI: gsd configure [--step N] [--list] [--verify]              │
└─────────────────────────────┬───────────────────────────────────┘
                              │
┌─────────────────────────────▼───────────────────────────────────┐
│                  CONFIGURATION LAYER                            │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────────────┐   │
│  │  Definitions │  │   Phase      │  │   Validation        │   │
│  │  (YAML/JSON) │  │   Manifest   │  │   Rules             │   │
│  │              │  │              │  │                      │   │
│  │  - Network   │  │  - Step 1    │  │   - Syntax          │   │
│  │  - Skills    │  │  - Step 2    │  │   - Consistency      │   │
│  │  - TCC       │  │  - Step 3   │  │   - Safety          │   │
│  └──────────────┘  └──────────────┘  └──────────────────────┘   │
└─────────────────────────────┬───────────────────────────────────┘
                              │
┌─────────────────────────────▼───────────────────────────────────┐
│                    EXECUTION ENGINE                             │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────────────┐   │
│  │   Network    │  │    Skill     │  │    Permission        │   │
│  │   Manager    │  │   Registry   │  │    Controller        │   │
│  │              │  │              │  │                      │   │
│  │  - pfctl     │  │  - Config    │  │  - TCC.db            │   │
│  │  - hosts     │  │    parsing   │  │  - Sandbox          │   │
│  │  - DNS       │  │  - Enable    │  │  - Entitlements     │   │
│  └──────────────┘  └──────────────┘  └──────────────────────┘   │
└─────────────────────────────┬───────────────────────────────────┘
                              │
┌─────────────────────────────▼───────────────────────────────────┐
│                   SYSTEM INTERFACE                              │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────────────┐   │
│  │   defaults   │  │   tccutil    │  │    security         │   │
│  │   command    │  │              │  │    command          │   │
│  └──────────────┘  └──────────────┘  └──────────────────────┘   │
└─────────────────────────────────────────────────────────────────┘
```

### Component Responsibilities

| Component | Responsibility | Communicates With |
|-----------|---------------|-------------------|
| **CLI Interface** | User-facing command parser, step orchestration | Configuration Layer |
| **Definitions** | YAML/JSON schemas for network, skills, permissions | Execution Engine (read) |
| **Phase Manifest** | Ordered list of configuration steps | Execution Engine (read) |
| **Validation Rules** | Syntax checks, dependency validation | Configuration Layer |
| **Network Manager** | Firewall rules, hosts file, allowed domains | System Interface |
| **Skill Registry** | Enable/disable OpenCode capabilities | Configuration Layer |
| **Permission Controller** | TCC database, sandbox settings, entitlements | System Interface |
| **System Interface** | macOS APIs: defaults, tccutil, security | OS Kernel |

---

## Data Flow

### Forward Flow (Configuration Application)

```
1. USER
   │
   ▼
2. CLI Parse: gsd configure --step 2
   │
   ▼
3. Validation Rules: Verify step 2 is valid, dependencies met
   │
   ├──▶ [VALID]   → 4. Phase Manifest: Load step 2 config
   │
   └──▶ [INVALID] → ERROR: Missing step 1 dependencies
   │
   ▼
4. Execution Engine: For each setting in step 2:
   │
   ├──▶ Network Manager → pfctl/hosts → System Interface
   ├──▶ Skill Registry  → OpenCode config  → Configuration File
   └──▶ Permission Ctrl → TCC database    → System Interface
   │
   ▼
5. Verification: Confirm settings applied, report status
   │
   ▼
6. USER: "Step 2 complete: GitLab access enabled"
```

### Reverse Flow (Status Verification)

```
USER: gsd configure --status
       │
       ▼
  Read all config files + current system state
       │
       ▼
  Compare: Defined vs. Applied
       │
       ▼
  Report:
  - Step 1: COMPLETE (3/3 settings verified)
  - Step 2: COMPLETE (5/5 settings verified)
  - Step 3: PENDING (waiting for user)
```

---

## Key Architectural Patterns

### Pattern 1: Declarative Configuration

**What:** Define desired state in YAML/JSON, let engine calculate differences.

**When:** Security configurations where auditability and idempotency matter.

**Example:**

```yaml
# network.yaml
version: 1
step: 1
name: "GitLab Access"
allowed_domains:
  - gitlab.com
  - gitlab.mycompany.com
firewall:
  default: deny
  allow_outgoing:
    - port: 443
      to: gitlab.com
    - port: 443
      to: gitlab.mycompany.com
```

```typescript
// Execution: Apply firewall rules
async function applyNetworkConfig(config: NetworkConfig): Promise<void> {
  // 1. Update pf.conf (packet filter)
  await exec('pfctl -f /etc/pf.conf');
  
  // 2. Update hosts file for domain resolution
  await updateHostsFile(config.allowed_domains);
  
  // 3. Enable pf if not already running
  await exec('pfctl -e');
}
```

### Pattern 2: Phased Deployment

**What:** Configuration applied in ordered steps, each building on the previous.

**When:** Security-critical systems where gradual trust establishment is required.

**Example from stronghold (similar approach):**

```python
# Phased approach: Each step adds permissions
# Step 0: Base hardening (always applied)
# Step 1: Git read access
# Step 2: Git write access  
# Step 3: Full network access
```

**Phase Dependency Graph:**

```
┌─────────┐     ┌─────────┐     ┌─────────┐
│ Step 0  │────▶│ Step 1  │────▶│ Step 2  │
│  Base   │     │  GitLab │     │ Network │
│         │     │   Read  │     │  Access │
└─────────┘     └─────────┘     └─────────┘
   │               │               │
   ▼               ▼               ▼
[Always]    [Requires 0]    [Requires 1]
```

### Pattern 3: Permission Least Privilege

**What:** Request minimal permissions, expand only when required.

**When:** macOS TCC (Transparency, Consent, Control) framework interactions.

**Example:**

```typescript
// Instead of requesting all permissions upfront:
// BAD: entitlements = ["*"]

// GOOD: Request only what's needed for current phase
const entitlements = {
  step1: ["com.apple.security.app-sandbox"],
  step2: ["com.apple.security.network.client"], 
  step3: ["com.apple.security.files.user-selected.read-write"]
};
```

---

## Anti-Patterns to Avoid

### Anti-Pattern 1: Mutable Global State

**What:** Modifying system settings directly without tracking current state.

**Why bad:** No rollback capability, impossible to verify configuration, potential for drift.

**Instead:** Use a configuration state file:

```typescript
// .gsd/state.json
{
  "applied_steps": [0, 1],
  "last_applied": "2026-04-05T10:30:00Z",
  "verification_hash": "abc123..."
}
```

### Anti-Pattern 2: Blind Execution

**What:** Applying settings without validation or verification.

**Why bad:** Silent failures, partial configurations, security gaps.

**Instead:** Always verify:

```typescript
async function applyAndVerify(config: Config): Promise<boolean> {
  await apply(config);
  const verified = await verify(config);
  if (!verified) {
    throw new Error(`Verification failed for ${config.step}`);
  }
  return true;
}
```

### Anti-Pattern 3: Tight Coupling

**What:** Hardcoding execution logic for each setting type.

**Why bad:** Adding new configuration types requires code changes.

**Instead:** Plugin architecture:

```typescript
interface ConfigApplicator {
  type: string;
  apply(config: unknown): Promise<void>;
  verify(config: unknown): Promise<boolean>;
}

// Registry of applicators
const applicators: Map<string, ConfigApplicator> = new Map([
  ['network', new NetworkApplicator()],
  ['skill', new SkillApplicator()],
  ['permission', new PermissionApplicator()],
]);
```

---

## Scalability Considerations

| Concern | At 1 User | At 10 Users | At 100+ Users |
|---------|-----------|-------------|---------------|
| **Config Storage** | Local YAML | Git-backed YAML | Config server + Git |
| **Verification** | CLI check | Cron job + alerting | MDM integration |
| **Rollback** | Manual + state file | Automated via state | Snapshot-based |
| **Network Rules** | Simple hosts file | pf rules per-user | Central firewall |

---

## Build Order (Suggested Phase Structure)

Based on component dependencies, recommended build order:

### Phase 1: Foundation (CLI + State Management)

```
1. CLI argument parser
2. Configuration loader (YAML)
3. State file manager
4. Basic validation
```

**Rationale:** Everything depends on CLI and loading. No execution without these.

### Phase 2: Definition Layer

```
5. Network config schema
6. Skill config schema  
7. Permission config schema
8. Phase manifest definition
```

**Rationale:** Defines what can be configured. Execution engine needs these schemas.

### Phase 3: Execution Engine (Read-Only)

```
9. Network manager (read-only)
10. Skill registry (read-only)
11. Permission controller (read-only)
```

**Rationale:** Read operations first. Verify current state before making changes.

### Phase 4: Execution Engine (Write)

```
12. Network manager (apply)
13. Skill registry (apply)
14. Permission controller (apply)
```

**Rationale:** Now that we can read state, we can safely write changes.

### Phase 5: Verification

```
15. Verification reporter
16. Diff generator (defined vs. applied)
17. Health check CLI
```

**Rationale:** Users need to know what was applied and if it succeeded.

### Phase 6: Polish

```
18. Rollback support
19. Export/import configs
20. Dry-run mode
```

**Rationale:** Nice-to-have features that make the tool production-ready.

---

## Sources

- **stronghold** — CLI macOS security configuration (GitHub, 1.2k stars)
  <https://github.com/alichtman/stronghold>
  
- **NIST macOS Security Compliance Project (mSCP)** — Automated security guidance
  <https://github.com/usnistgov/macos_security>
  
- **MACE (Mac Advanced Compliance Editor)** — Visual compliance baseline management
  <https://getmace.com/>
  
- **macos_hardening** — Shell-based hardening scripts
  <https://github.com/ataumo/macos_hardening>

---

## Confidence Assessment

| Area | Level | Notes |
|------|-------|-------|
| Component Boundaries | MEDIUM | Based on tool patterns, some inference |
| Data Flow | HIGH | Standard Unix pipeline + config tool patterns |
| Build Order | HIGH | Clear dependencies between layers |
| Anti-Patterns | MEDIUM | Common patterns, some macOS-specific unknowns |
