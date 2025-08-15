# gvm

<p align="center">
<img width="512" height="512" alt="icon" src="https://github.com/user-attachments/assets/d633b1b0-d0af-40e9-9d29-df16dfac3d39" />
</p>

A simple Go Version Management Tool with zero dependencies

[![Go Report Card](https://goreportcard.com/badge/github.com/lokesh-balla/gvm)](https://goreportcard.com/report/github.com/lokesh-balla/gvm)
<p>
<img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/lokesh-balla/gvm">
<a href="https://github.com/lokesh-balla/gvm/releases"><img alt="GitHub release (latest by date)" src="https://img.shields.io/github/v/release/lokesh-balla/gvm"></a>
<a href="https://pkg.go.dev/github.com/lokesh-balla/gvm?tab=doc"><img src="https://godoc.org/github.com/golang/gddo?status.svg" alt="GoDoc"></a>
<img alt="GitHub" src="https://img.shields.io/github/license/lokesh-balla/gvm">
<img alt="GitHub Workflow Status (with branch)" src="https://img.shields.io/github/actions/workflow/status/lokesh-balla/gvm/go.yml?branch=main">
<img alt="GitHub Workflow Status (with branch)" src="https://img.shields.io/github/actions/workflow/status/lokesh-balla/gvm/golangci-lint.yml?branch=main&label=golangci-lint">
<img alt="GitHub Workflow Status (with branch)" src="https://img.shields.io/github/actions/workflow/status/lokesh-balla/gvm/codeql.yml?branch=main&label=CodeQL">
<img alt="GitHub repo size" src="https://img.shields.io/github/repo-size/lokesh-balla/gvm">
<img alt="GitHub all releases" src="https://img.shields.io/github/downloads/lokesh-balla/gvm/total">
<img alt="GitHub commits since latest release" src="https://img.shields.io/github/commits-since/lokesh-balla/gvm/latest">
</p>


## Installing

```bash
curl -sL https://raw.githubusercontent.com/lokesh-balla/gvm/main/install.sh | sh - 
```
![Installation GIF](docs/demo.gif)


## Usage

To check how to use gvm you can follow the below doc
* [gvm](docs/gvm.md)	 - A golang version manager

or check the help command

```bash
gvm --help
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

```bash
rm -rf ~/.gvm
```

### Removing PATH

Check and remove any PATH set for $HOME/.gvm or $HOME/.gvm/bin in ~/.profile, ~/.bashrc and ~/.zshenv
