# tgbot

Minimal CLI for the Telegram Bot API written in Go.

## Build

- The binary must be small. Always use `build-small` (stripped, static, CGO_ENABLED=0, `-trimpath -ldflags="-s -w"`) or `build-tiny` (+ UPX) as the default build target.
- Do not add dependencies. This project has zero external dependencies — keep it that way.
- Cross-compilation targets: linux/amd64, linux/arm64, darwin/amd64, darwin/arm64.

## Releases

- Tags follow semver: `v0.1.0`, `v0.2.0`, etc.
- Pushing a `v*` tag triggers the GitHub Actions workflow that builds all platform binaries and creates a GitHub release.

## Code style

- Keep it simple and minimal. No frameworks, no unnecessary abstractions.
- All commands output JSON.
