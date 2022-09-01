# compver

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gowizzard/compver.svg)](https://golang.org/) [![Go](https://github.com/gowizzard/compver/actions/workflows/go.yml/badge.svg)](https://github.com/gowizzard/compver/actions/workflows/go.yml) [![CodeQL](https://github.com/gowizzard/compver/actions/workflows/codeql.yml/badge.svg)](https://github.com/gowizzard/compver/actions/workflows/codeql.yml) [![VMerge](https://github.com/gowizzard/compver/actions/workflows/vmerge.yml/badge.svg)](https://github.com/gowizzard/compver/actions/workflows/vmerge.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/gowizzard/compver/v2)](https://goreportcard.com/report/github.com/gowizzard/compver/v2) [![GitHub issues](https://img.shields.io/github/issues/gowizzard/compver)](https://github.com/gowizzard/compver/issues) [![GitHub forks](https://img.shields.io/github/forks/gowizzard/compver)](https://github.com/gowizzard/compver/network) [![GitHub stars](https://img.shields.io/github/stars/gowizzard/compver)](https://github.com/gowizzard/compver/stargazers) [![GitHub license](https://img.shields.io/github/license/gowizzard/compver)](https://github.com/gowizzard/compver/blob/master/LICENSE)

With this small cli tool, you can compare two versions with each other. **Currently, it is important to know that we only compare the version core (x.y.z).** This means that currently the pre-releases and the meta information are not taken into account.

Our version specifications are based on semantic versioning. [Here](https://semver.org/) you can find the corresponding definition.

## Installation

In order to install the tool correctly, you need to download it [here](https://github.com/gowizzard/compver/releases) and make it executable. After that you can move the binary into the application folder and use it via the terminal.

```bash
# Download the binary
curl -L -o compver <URL>

# Make the binary executable 
chmod +x compver

# Move to application directory
sudo mv compver /usr/local/bin
```

## How to use

Actually, you don't need to know much to execute the command. You actually only compare the newest version with the older version. We assume that `VERSION1` is the newer version and `VERSION2` is the older one. But downgrades of versions can also be recorded.

```bash
compver -version1 <VERSION1> -version2 <VERSION2>
```

If this command is executed now, we get an information back, this can contain the following information: `no changes`, `major update`, `major downgrade`, `minor update`, `minor downgrade`, `patch update` & `patch downgrade`.

## Special thanks

Thanks to [JetBrains](https://github.com/JetBrains) for supporting me with this and other [open source projects](https://www.jetbrains.com/community/opensource/#support).