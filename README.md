# kde-tool

Small KDE/Plasma automation utilities written in Go.

This repository currently provides:
- A CLI for KDE activity switching and Night Light toggling over D-Bus
- A `qdbus` package with typed clients for KDE D-Bus interfaces
- A `kscreen` package to build and apply `kscreen-doctor` output configs

## Features

### CLI actions
The CLI currently supports:
- Switch to next activity
- Switch to previous activity
- Toggle Night Light

### D-Bus clients
The `qdbus` package includes:
- `ActivityClient` for `org.kde.ActivityManager`
- `NightLightClient` for `org.kde.KWin.NightLight`
- A reusable typed D-Bus interface wrapper (`QDBusInterfaceClient`)

### Screen configuration helpers
The `kscreen` package includes:
- Data model for outputs, mode, size, and position
- Argument generation for `kscreen-doctor`
- `EnableOnly(...)` helper that enables selected outputs and disables others

Note: the `kscreen` package is available in the repo but is not currently exposed through the CLI entrypoint.

## Requirements

- Linux desktop session running KDE Plasma
- D-Bus session bus available (`DBUS_SESSION_BUS_ADDRESS` set by session)
- `kscreen-doctor` installed (required only when using the `kscreen` package)
- Go `1.26.3` or newer

## Build

### Using Make

```bash
make build
```

This produces the binary:
- `./kde-tool`

### Using Go directly

```bash
go build -o kde-tool ./cmd
```

## Usage

The root command name is currently `kde`.

```bash
./kde-tool --help
```

Available flags:

```text
--next-activity      switch to the next activity
--prev-activity      switch to the previous activity
--toggle-nightlight  toggle night light
```

Examples:

```bash
# Go to the next activity
./kde-tool --next-activity

# Go to the previous activity
./kde-tool --prev-activity

# Toggle Night Light
./kde-tool --toggle-nightlight
```

## Release artifacts

Build multi-arch Linux release binaries:

```bash
make release
```

Output format:
- `dist/kde-tool-<version>-linux-amd64`
- `dist/kde-tool-<version>-linux-arm64`

Version is derived from:

```bash
git describe --tags --always --dirty
```

## Project layout

```text
cmd/main.go           CLI entrypoint (cobra)
qdbus/                KDE D-Bus clients and typed wrappers
kscreen/              Screen config modeling and kscreen-doctor integration
Makefile              build/release targets
```
