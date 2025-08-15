---
layout: default
title: List Command
parent: Commands
nav_order: 2
---

# List Command

List installed Go versions and available versions from the official Go download server.

## Usage

```bash
gvm list
gvm list --all
```

## Description

The list command displays all Go versions that are currently installed on your system. When used with the `--all` flag, it also shows available versions that can be installed.

## Examples

### List installed versions only
```bash
gvm list
```

**Output:**
```
OS: linux ARCH: amd64

Installed Versions
go1.21.0 ← default
go1.20.5
go1.19.10
```

### List installed and available versions
```bash
gvm list --all
```

**Output:**
```
OS: linux ARCH: amd64

Installed Versions
go1.21.0 ← default
go1.20.5
go1.19.10

Available Versions
go1.21.1
go1.21.0
go1.20.10
go1.20.9
...
```

## Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--all` | `-a` | List all available versions from the Go download server |

## Output Format

The list command displays:

- **OS and Architecture**: Shows your current operating system and architecture
- **Installed Versions**: Lists all locally installed Go versions
- **Default Indicator**: Shows which version is currently set as default with `← default`
- **Available Versions**: Shows versions available for download (when using `--all`)

## Filtering

The command automatically filters available versions based on:
- Your current operating system
- Your current architecture
- Only archive packages (not source or installer packages)

## Related Commands

- [install](install) - Install a specific version
- [use](use) - Switch to a different installed version
- [uninstall](uninstall) - Remove an installed version
