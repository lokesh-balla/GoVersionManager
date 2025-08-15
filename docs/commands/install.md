---
layout: default
title: Install Command - gvm
parent: Commands
nav_order: 1
---

# Install Command

Install Go versions using the `install` command.

## Usage

```bash
gvm install [version]
gvm install --latest
```

## Description

The install command allows you to download and install specific Go versions or the latest stable release. It automatically handles:
- Download verification via SHA256 checksum
- Progress tracking during download
- Automatic extraction and cleanup
- Metadata tracking of installed versions

## Examples

### Install a specific version
```bash
gvm install go1.19
```

### Install the latest stable version
```bash
gvm install --latest
```

## Flags

| Flag | Description |
|------|-------------|
| `--latest` | Install the latest stable Go version |

## Installation Process

When you run `gvm install`, the following happens:

1. **Download**: The specified Go version is downloaded from the official Go download server
2. **Verification**: SHA256 checksum is verified to ensure integrity
3. **Extraction**: The downloaded archive is extracted to `~/.gvm/[version]`
4. **Cleanup**: The downloaded archive is automatically removed
5. **Registration**: The version is registered in the local metadata database

## Error Handling

The install command provides detailed error messages for common issues:
- Invalid version format
- Network connectivity problems
- Checksum verification failures
- Insufficient disk space

## Related Commands

- [list](list) - List available/installed versions
- [use](use) - Switch to an installed version
- [uninstall](uninstall) - Remove an installed version
