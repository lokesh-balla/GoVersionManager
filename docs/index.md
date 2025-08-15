---
layout: default
title: GoVersionManager - Home
---

# GoVersionManager

A simple Go Version Management Tool with zero dependencies

[![Go Report Card](https://goreportcard.com/badge/github.com/Lokesh-Balla/GoVersionManager)](https://goreportcard.com/report/github.com/Lokesh-Balla/GoVersionManager)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Lokesh-Balla/GoVersionManager)](https://github.com/Lokesh-Balla/GoVersionManager)
[![GitHub release](https://img.shields.io/github/v/release/Lokesh-Balla/GoVersionManager)](https://github.com/Lokesh-Balla/GoVersionManager/releases)
[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://pkg.go.dev/github.com/Lokesh-Balla/GoVersionManager?tab=doc)
[![GitHub license](https://img.shields.io/github/license/Lokesh-Balla/GoVersionManager)](LICENSE)

## Quick Start

### Installing

```bash
curl -sL https://raw.githubusercontent.com/Lokesh-Balla/GoVersionManager/main/install.sh | sh - 
```

![Installation GIF](demo.gif)

### Basic Usage

Once installed, you can use `gvm` to manage Go versions:

```bash
# List available Go versions
gvm list

# Install a specific version
gvm install 1.21.0

# Use a specific version
gvm use 1.21.0

# Uninstall a version
gvm uninstall 1.21.0
```

## Available Commands

GoVersionManager provides the following commands:

- **[install](commands/install)** - Install Go versions
- **[list](commands/list)** - List available/installed versions
- **[use](commands/use)** - Switch between Go versions
- **[uninstall](commands/uninstall)** - Remove installed versions

For detailed information about each command, visit the [commands section](commands/).

## Getting Help

You can get help for any command using:

```bash
gvm --help
gvm [command] --help
```

## Compiling From Source

### Prerequisites

- Linux, Darwin, FreeBSD or any other unix based OS
- git
- GoLang version 1.19 or higher

### Compiling

```bash
git clone https://github.com/Lokesh-Balla/GoVersionManager.git && cd GoVersionManager
go build -o gvm main.go
```

## Uninstall

```bash
rm -rf ~/.gvm
```

### Removing PATH

Check and remove any PATH set for `$HOME/.gvm` or `$HOME/.gvm/bin` in `~/.profile`, `~/.bashrc` and `~/.zshenv`
