# TemporalCI-test

A test repository for [TemporalCI](https://github.com/AndreKurait/TemporalCI) — demonstrating every feature of the CI/CD platform.

![Build](https://your-temporalci-host/badge/AndreKurait/TemporalCI-test)

## Features Demonstrated

### 6 Named Pipelines (`.temporalci/`)

| Pipeline | Trigger | Features Used |
|----------|---------|---------------|
| `ci.yaml` | push + PR | Matrix builds, service containers, DinD, conditionals, artifacts, gate, parameters, post hooks, path filtering |
| `security.yaml` | push + weekly cron | Scheduled triggers, parallel steps |
| `release.yaml` | tag push (`v*`) | Tag triggers, cross-compilation matrix (2×2), artifacts, parameters, gate |
| `auto-label.yaml` | issue opened | Issue event handling, utility workflow (no pod) |
| `branch-cleanup.yaml` | PR closed | PR lifecycle events, conditional (`startsWith`), utility workflow |

### CI Pipeline Features

- **Matrix builds** — Tests run across Go 1.22 and 1.23 in parallel
- **Service containers** — Redis sidecar with health check for integration tests
- **Docker-in-Docker** — Builds a container image inside CI
- **Conditional steps** — `coverage-check` only runs on PRs, `docker-build` only on push
- **Parameterized** — `COVERAGE_THRESHOLD` (choice) and `RUN_E2E` (boolean)
- **Artifacts** — Build binary and coverage reports uploaded to S3
- **Gate step** — `all-checks-pass` aggregates all results for branch protection
- **Post hooks** — `always` reports metrics, `on_failure` sends notification
- **Path filtering** — CI only triggers on `.go`, `go.mod`, or `.temporalci/` changes

### CI Dashboard

View builds at: `https://your-temporalci-host/ci/builds`

- Build list with repo/branch/status filters
- Build detail with step timeline, inline logs, DAG visualization
- Manual trigger UI at `/ci/triggers`
- Repo overview at `/ci/repos`
- Analytics at `/ci/analytics`

## Local Development

```bash
go test -v -race ./...
go build ./...
```
