---
status: resolved
trigger: "npx get-shit-done-cc@latest hangs in sandbox"
created: 2026-04-05T00:00:00Z
updated: 2026-04-05T00:00:00Z
---

## Current Focus
hypothesis: FIXED - Missing tls: skip in npm_registry endpoints caused HTTPS proxy to hang
test: Added tls: skip to npm_registry endpoints and applied policy version 9
expecting: npx commands should now work without hanging
next_action: User verification needed - run npx get-shit-done-cc@latest in sandbox

## Symptoms
expected: npx get-shit-done-cc@latest should execute and download/run the package
actual: Command hangs indefinitely, no output
errors: None visible (hangs before any output)
reproduction: Run `npx get-shit-done-cc@latest` in sandbox
started: After adding /usr/bin/node to npm_registry binaries (previously worked)

## Eliminated
- hypothesis: Binary not allowed in policy
  evidence: Logs showed binary was allowed in later entries but still hung
  timestamp: 2026-04-05T00:00:00Z

- hypothesis: Endpoint not in policy
  evidence: Endpoint was in npm_registry policy
  timestamp: 2026-04-05T00:00:00Z

## Evidence
- timestamp: 2026-04-05T00:00:00Z
  checked: "openshell logs demo"
  found: "HTTP_REQUEST to /get-shit-done-cc logged, but NO HTTP_RESPONSE - 225 second gap between requests"
  implication: "Request goes through but response times out - proxy issue"

- timestamp: 2026-04-05T00:00:00Z
  checked: "demo-policy.yaml - npm_registry vs git_pull"
  found: "npm_registry had NO tls config, git_pull has 'tls: skip'"
  implication: "Missing tls: skip causes HTTPS proxy to hang on TLS handshake"

## Resolution
root_cause: "The npm_registry policy was missing 'tls: skip' on endpoint configurations. When the proxy connects to HTTPS endpoints without TLS termination bypass, the connection hangs waiting for TLS handshake that never completes through the proxy."
fix: "Added 'tls: skip' to both registry.npmjs.org and api.npmjs.org endpoints in demo-policy.yaml, then applied policy version 9 to sandbox."
verification: "Policy version 9 loaded successfully. User needs to verify npx command works."
files_changed:
- ".oshell-config/demo-policy.yaml": Added tls: skip to npm_registry endpoints