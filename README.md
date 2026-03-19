# TemporalCI-test

A test repository for [TemporalCI](https://github.com/AndreKurait/TemporalCI) — a Kubernetes-native CI system built on Temporal.

## What Happens When You Push

Every push and PR triggers a **TemporalCI pipeline** that:

1. **Clones** this repo into an isolated K8s pod
2. **Runs steps as a DAG** — independent steps execute in parallel
3. **Reports results** directly on the PR with timing, logs, and a link to the Temporal Web UI

### Pipeline (``.temporalci.yaml``)

```
build ──┬── unit-test ── race-test
        ├── vet
        └── format-check
```

- `build` runs first
- `unit-test`, `vet`, and `format-check` run **in parallel** after build
- `race-test` runs after unit tests pass

### Helm Chart Testing

This repo also includes a Helm chart at `deploy/helm/`. TemporalCI can lease an **ephemeral EKS cluster** from a warm pool, install the chart, run `helm test`, and report results — all orchestrated as a durable Temporal workflow.

## Functions

| Function | Description |
|----------|-------------|
| `Greet(name)` | Returns a greeting |
| `Reverse(s)` | Reverses a string |
| `IsPalindrome(s)` | Checks if a string is a palindrome |
| `Anagram(a, b)` | Checks if two strings are anagrams |
