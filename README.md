# tmpl

[![Go status](https://github.com/haya14busa/tmpl/workflows/Go/badge.svg)](https://github.com/haya14busa/tmpl/actions)
[![releases](https://img.shields.io/github/release/haya14busa/tmpl.svg)](https://github.com/haya14busa/tmpl/releases)

## Installation

```shell
# Install latest version. (Install it into ./bin/ by default).
$ curl -sfL https://raw.githubusercontent.com/haya14busa/tmpl/master/install.sh| sh -s

# Specify installation directory ($(go env GOPATH)/bin/) and version.
$ curl -sfL https://raw.githubusercontent.com/haya14busa/tmpl/master/install.sh| sh -s -- -b $(go env GOPATH)/bin [vX.Y.Z]

# In alpine linux (as it does not come with curl by default)
$ wget -O - -q https://raw.githubusercontent.com/haya14busa/tmpl/master/install.sh| sh -s [vX.Y.Z]

$ go get github.com/haya14busa/tmpl/cmd/tmpl

# homebrew / linuxbrew
$ brew install haya14busa/tap/tmpl
$ brew upgrade haya14busa/tap/tmpl

# Go
$ go get github.com/haya14busa/tmpl
```

## tmpl -h

```
$ tmpl -h
Usage: tmpl [FLAGS] [Template Files]
        Generate textual output using Go text/template from given data in STDIN
Flags:
  -t string
        Go text/template text template
  -version
        print version

Syntax: https://golang.org/pkg/text/template

Functions:
   text/template builtin:
        See https://golang.org/pkg/text/template/#hdr-Functions
   env:
        Return environment value of its argument.
        e.g. {{ env "HOME" }}
   strings:
        Go "strings" package functions. https://golang.org/pkg/strings/
        e.g. {{ strings.Title "test" }}

GitHub: https://github.com/haya14busa/tmpl
```
