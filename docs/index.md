---
layout: default
title: Home
nav_order: 1
---

# gvm
{: .fs-9 }

A simple Go Version Management Tool with zero dependencies
{: .fs-6 .fw-300 }

[Get Started](#quick-start){: .btn .btn-primary .fs-5 .mb-4 .mb-md-0 .mr-2 }
[View on GitHub](https://github.com/lokesh-balla/gvm){: .btn .fs-5 .mb-4 .mb-md-0 }

---

[![Go Report Card](https://goreportcard.com/badge/github.com/lokesh-balla/gvm)](https://goreportcard.com/report/github.com/lokesh-balla/gvm)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/lokesh-balla/gvm)](https://github.com/lokesh-balla/gvm)
[![GitHub release](https://img.shields.io/github/v/release/lokesh-balla/gvm)](https://github.com/lokesh-balla/gvm/releases)
[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://pkg.go.dev/github.com/lokesh-balla/gvm?tab=doc)
[![GitHub license](https://img.shields.io/github/license/lokesh-balla/gvm)](LICENSE)

## Quick Start

### Installing

```bash
curl -sL https://raw.githubusercontent.com/lokesh-balla/gvm/main/install.sh | sh
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

gvm provides the following commands:

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
git clone https://github.com/lokesh-balla/gvm.git && cd gvm
go build -o gvm main.go
```

## Uninstall

To uninstall `gvm`, run the following command:

```bash
rm -rf ~/.gvm
