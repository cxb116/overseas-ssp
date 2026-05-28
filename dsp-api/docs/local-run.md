# Local Run

This project needs Go 1.25+.

- `config.json` is the online Linux release config.
- `config.json1` is the local/test config and is used by `scripts/run-local.sh` by default.

## Why `go run` may fail

If your shell still exports an old `GOROOT` such as `/usr/local/go`, the `go` binary may be new while the standard library is old. In that case startup fails with errors like:

- `package slices is not in std`
- `package maps is not in std`
- `package log/slog is not in std`

## Recommended start command

Run:

```bash
./scripts/run-local.sh
```

If you want a different config file:

```bash
./scripts/run-local.sh ./config.json
```

The script tries to detect Homebrew's Go installation and exports a usable `GOROOT` before starting the service.
