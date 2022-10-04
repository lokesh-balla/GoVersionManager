# GoVersionManager

A simple Go Version Management Tool with zero dependencies

## Installing

```bash
curl -sfL https://github.com/Lokesh-Balla/GoVersionManager/blob/main/install.sh | sh - 
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
