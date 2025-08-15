---
layout: default
title: Uninstall Command - gvm
parent: Commands
nav_order: 4
---

# Uninstall Command

Remove installed Go versions from your system.

## Usage

```bash
gvm uninstall [version]
gvm remove [version]
```

## Description

The uninstall command (also available as `remove`) allows you to remove specific Go versions that are no longer needed. This helps free up disk space and keeps your Go installation directory clean.

## Examples

### Remove a specific version
```bash
gvm uninstall go1.19
```

### Use the alias command
```bash
gvm remove go1.18
```

## Safety Checks

The uninstall command includes important safety checks:

1. **Version Validation**: Verifies the version exists and is installed
2. **Default Version Protection**: Prevents removal of the currently active default version
3. **Clean Removal**: Completely removes the version directory and all associated files

## Error Handling

The command provides clear error messages for:
- Non-existent versions
- Attempting to remove the default version
- Permission issues during file removal

## Important Notes

- **Cannot remove default version**: You must switch to a different version using `gvm use` before removing the current default
- **Permanent removal**: Once uninstalled, the version must be re-downloaded if needed again
- **No confirmation prompt**: The command executes immediately without confirmation

## Workflow Example

To safely remove a version:

1. Check installed versions:
   ```bash
   gvm list
   ```

2. If the version is default, switch to another:
   ```bash
   gvm use go1.20
   ```

3. Remove the version:
   ```bash
   gvm uninstall go1.19
   ```

## Related Commands

- [install](install) - Install a new version
- [list](list) - See installed versions
- [use](use) - Switch to a different version
