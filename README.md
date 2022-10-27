# GoVersionManager

A simple Go Version Management Tool with zero dependencies

[![Go Report Card](https://goreportcard.com/badge/github.com/Lokesh-Balla/GoVersionManager)](https://goreportcard.com/report/github.com/Lokesh-Balla/GoVersionManager)
<p>
<img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/Lokesh-Balla/GoVersionManager">
<a href="https://github.com/Lokesh-Balla/GoVersionManager/releases"><img alt="GitHub release (latest by date)" src="https://img.shields.io/github/v/release/Lokesh-Balla/GoVersionManager"></a>
<a href="https://pkg.go.dev/github.com/Lokesh-Balla/GoVersionManager?tab=doc"><img src="https://godoc.org/github.com/golang/gddo?status.svg" alt="GoDoc"></a>
<a href="https://github.com/Lokesh-Balla/GoVersionManager/actions"><img alt="GitHub Workflow Status" src="https://img.shields.io/github/workflow/status/Lokesh-Balla/GoVersionManager/Go"></a>
</p>

## Installing

```bash
curl -sL https://raw.githubusercontent.com/Lokesh-Balla/GoVersionManager/main/install.sh | sh - 
```

## Usage

To check how to use gvm you can just use the help command

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
git clone https://github.com/Lokesh-Balla/GoVersionManager.git && cd GoVersionManager
go build -o gvm main.go
```

## Uninstall

```bash
rm -rf ~/.gvm
```

### Removing PATH

Check and remove any PATH set for $HOME/.gvm or $HOME/.gvm/bin in ~/.profile, ~/.bashrc and ~/.zshenv
