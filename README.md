<div align="center">

<img src="https://user-images.githubusercontent.com/30717818/190334331-e1aea304-eb4e-4848-9333-2a8156be26ba.svg" width="150">

# CompVer

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gowizzard/compver.svg)](https://golang.org/) [![Go](https://github.com/gowizzard/compver/actions/workflows/go.yml/badge.svg)](https://github.com/gowizzard/compver/actions/workflows/go.yml) [![CodeQL](https://github.com/gowizzard/compver/actions/workflows/codeql.yml/badge.svg)](https://github.com/gowizzard/compver/actions/workflows/codeql.yml) [![CompVer](https://github.com/gowizzard/compver/actions/workflows/compver.yml/badge.svg)](https://github.com/gowizzard/compver/actions/workflows/compver.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/gowizzard/compver/v3)](https://goreportcard.com/report/github.com/gowizzard/compver/v3) [![GitHub issues](https://img.shields.io/github/issues/gowizzard/compver)](https://github.com/gowizzard/compver/issues) [![GitHub forks](https://img.shields.io/github/forks/gowizzard/compver)](https://github.com/gowizzard/compver/network) [![GitHub stars](https://img.shields.io/github/stars/gowizzard/compver)](https://github.com/gowizzard/compver/stargazers) [![GitHub license](https://img.shields.io/github/license/gowizzard/compver)](https://github.com/gowizzard/compver/blob/master/LICENSE)

</div>

With this small cli tool, you can compare two versions with each other. **Currently, it is important to know that we only compare the version core `x.y.z`.** This means that currently the pre-releases and the meta information are not taken into account.

Our version specifications are based on semantic versioning. [Here](https://semver.org/) you can find the corresponding definition. So that major, minor and patch can be read correctly, we use regex. You can find [here](https://regex101.com/r/un81dE/4) at regex101 once the used expression. We made sure that we can read versions with a preceding `v` as well, even though we know it's not part of semantic versioning. [Here](https://semver.org/#is-v123-a-semantic-version) you can find an explanation in the faq on the [semver.org](https://semver.org/) website.

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

Here you can find the different statements you can use.

### Compare versions

Actually, you don't need to know much to execute the command. You actually only compare the newest version with the older version. We assume that `VERSION1` is the newer version and `VERSION2` is the older one. But downgrades of versions can also be recorded.

```bash
compver -compare -version1 <VERSION1> -version2 <VERSION2>
```

If this command is executed now, we get an information back, this can contain the following information: `no changes`, `major update`, `major downgrade`, `minor update`, `minor downgrade`, `patch update` & `patch downgrade`.

### Get version core block

If you want to read a block of the version core, you can do this with the following command. You will receive the number of the block as an answer. You can read and return the blocks `major`, `minor` and `patch`.

```bash
compver -core -block major -version1 <VERSION1>
```

## Using the GitHub Action

Here you can find an example if you want to use CompVer as a GitHub Action. In the example the action is only triggered when a new release is created.

I use this example a lot when I need to maintain major branches. Especially when developing Golang libraries this is very important. The action takes the release version and determines the major block of the version core. Then the version branch, for example `v3`, is merged with the default branch, so that the version branch is always up-to-date. So you only have to create the branch once and don't have to worry about it not being maintained.

```yaml
name: CompVer

on:
  push:
    tags:
      - "v*.*.*"

env:
  USER_NAME: "GitHub Action"
  USER_EMAIL: "actions@github.com"
  DEFAULT_BRANCH: ${{ github.event.repository.default_branch }}
  COMMIT_MESSAGE: "ci: The data of the master branch was merged automatically."

jobs:
  version:
    runs-on: ubuntu-latest
    steps:

      - name: Clone repository
        uses: actions/checkout@v3
        with:
          fetch-depth: 0

      - name: Set repository tag information
        id: information
        run: echo ::set-output name=tag::${GITHUB_REF#refs/*/}

      - name: Get the major version
        id: compver
        uses: gowizzard/compver@v3.3.0
        with:
          args: "-core -block major -version1 ${{ steps.information.outputs.tag }}"

      - name: Set git config
        run: |
          git config --local user.name "$USER_NAME"
          git config --local user.email "$USER_EMAIL"

      - name: Merge data from default branch
        run: |
          git fetch
          git checkout v${{ steps.compver.outputs.core_number }}
          git pull
          git merge --no-ff "origin/$DEFAULT_BRANCH" -m "$COMMIT_MESSAGE"
          git push
```

## Special thanks

Thanks to [JetBrains](https://github.com/JetBrains) for supporting me with this and other [open source projects](https://www.jetbrains.com/community/opensource/#support).
