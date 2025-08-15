---
layout: default
title: Use Command
parent: Commands
nav_order: 3
---

# Use Command

Switch between installed Go versions by setting a specific version as the default.

## Usage

```bash
gvm use [version]
```

## Description

The use command allows you to switch between different installed Go versions by creating a symbolic link to the specified version. This makes the selected version the default Go version for your system.

## Examples

### Set a specific version as default
```bash
gvm use go1.19
```

## How It Works

When you run `gvm use`, the following happens:

1. **Validation**: Checks if the specified version is installed
2. **Symlink Creation**: Creates a symbolic link from `~/.gvm/go` to the selected version
3. **Database Update**: Updates the metadata to mark the selected version as default
4. **PATH Integration**: The symlink ensures your PATH points to the correct Go installation

## Requirements

- The specified version must be installed (use `gvm list` to see installed versions)
- The version must be specified exactly as shown in the installed versions list

## Error Handling

The use command provides clear error messages for:
- Non-existent versions
- Versions that aren't installed
- Permission issues when creating symlinks

## Verification

After using a version, you can verify it's active by running:
```bash
go version
```

This should display the version you just set as default.

## Related Commands

- [install](install) - Install a new version
- [list](list) - See installed versions
- [uninstall](uninstall) - Remove a version
