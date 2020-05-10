# tmpl - Template engine for CLI to format JSON data

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
Usage: tmpl [FLAGS] [Template Files]
        Generate textual output using Go text/template from given data in STDIN
Flags:
  -f value
        input format. Available format: [json (default), jsonl (http://jsonlines.org/)]
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
   sp:
        Space
   nl:
        New line
   tab:
        Tab character
   pos.Offset:
        Return position[1] of file from a given offset.
        e.g. {{ pos.Offset "file.txt" 14 }}
        [1]: https://godoc.org/github.com/haya14busa/offset#Position

Examples:
  $ echo '{"data": 14}' | tmpl -t 'data={{.data}},user={{ strings.ToUpper (env "USER") }}'
  data=14,user=HAYA14BUSA
  $ cat _testdata/base.json
  {
    "str": "string test",
    "num": 14,
    "array": [1,2,3],
    "nested": { "value": "ok"}
  }
  $ cat _testdata/base.tmpl
  str: {{ .str }}
  num: {{ .num }}
  Loop array
  {{ range $idx,$ele := .array -}}
  - {{ $idx }}: element={{ $ele }}
  {{ end -}}

  Nest: {{ .nested.value }}
  $ tmpl _testdata/base.tmpl < _testdata/base.json
  str: string test
  num: 14
  Loop array
  - 0: element=1
  - 1: element=2
  - 2: element=3
  Nest: ok

GitHub: https://github.com/haya14busa/tmpl
```

### tmpl -f=jsonl demo
![tmpl jsonl demo](https://user-images.githubusercontent.com/3797062/81494706-e2e44c00-92e5-11ea-9f91-6a9d341c1638.gif)
