package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"text/template"

	funcs "github.com/haya14busa/tmpl/funcs"
)

// Populated during build.
var version = "master"

// Format represents input format.
type Format int

const (
	FormatJSON Format = iota
	FormatJSONL
)

// String implements the flag.Value interface
func (f *Format) String() string {
	names := [...]string{
		"json",
		"jsonl",
	}
	if *f < FormatJSON || *f > FormatJSONL {
		return "Unknown mode"
	}

	return names[*f]
}

// Set implements the flag.Value interface
func (f *Format) Set(value string) error {
	switch value {
	case "", "json":
		*f = FormatJSON
	case "jsonl":
		*f = FormatJSONL
	default:
		return fmt.Errorf("invalid format name: %s", value)
	}
	return nil
}

type option struct {
	version bool
	tmpl    string
	format  Format
}

func setupFlags(opt *option) {
	flag.BoolVar(&opt.version, "version", false, "print version")
	flag.StringVar(&opt.tmpl, "t", "", "Go text/template text template")
	flag.Var(&opt.format, "f", "input format. Available format: [json (default), jsonl (http://jsonlines.org/)]")
	flag.Usage = usage
	flag.Parse()
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: tmpl [FLAGS] [Template Files]")
	fmt.Fprintln(os.Stderr, "\tGenerate textual output using Go text/template from given data in STDIN")
	fmt.Fprintln(os.Stderr, "Flags:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Syntax: https://golang.org/pkg/text/template")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintf(os.Stderr, `Functions:
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
`)
	fmt.Fprintln(os.Stderr, `Examples:
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
  Nest: ok`)
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "GitHub: https://github.com/haya14busa/tmpl")
	os.Exit(2)
}

func main() {
	var opt option
	setupFlags(&opt)
	if opt.version {
		fmt.Fprintln(os.Stdout, version)
		os.Exit(0)
	}
	if err := run(os.Stdin, os.Stdout, flag.Args(), opt); err != nil {
		fmt.Fprintf(os.Stderr, "tmpl: %v\n", err)
		os.Exit(1)
	}
}

func run(r io.Reader, w io.Writer, args []string, opt option) error {
	var d interface{}
	var errCh chan (error)
	switch opt.format {
	case FormatJSON:
		if err := json.NewDecoder(r).Decode(&d); err != nil {
			return err
		}
	case FormatJSONL:
		errCh = make(chan error, 1)
		jsonlinesCh := make(chan interface{})
		go jsonl(r, jsonlinesCh, errCh)
		d = jsonlinesCh
	}
	t, err := buildTemplate(args, opt)
	if err != nil {
		return fmt.Errorf("failed to build template: %v", err)
	}
	if err := t.Execute(w, d); err != nil {
		return err
	}
	select {
	case err := <-errCh:
		return err
	default:
	}
	return nil
}

func buildTemplate(files []string, opt option) (*template.Template, error) {
	funcMap := template.FuncMap{
		"env":     os.Getenv,
		"strings": funcs.Strings,
		"pos":     funcs.Pos,
		"sp":      func() string { return " " },
		"tab":     func() string { return "\t" },
		"nl":      func() string { return "\n" },
	}
	name := "tmpl"
	if len(files) > 0 {
		name = filepath.Base(files[0])
	}
	t := template.New(name).Funcs(funcMap)
	if len(files) > 0 {
		return t.ParseFiles(files...)
	}
	if opt.tmpl != "" {
		var err error
		t, err = t.Parse(opt.tmpl)
		if err != nil {
			return nil, err
		}
		return t, nil
	}
	return nil, errors.New("No templates specified. Use -t or pass template files as arguments.")
}

func jsonl(r io.Reader, jsonlines chan<- interface{}, errCh chan<- error) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		var l interface{}
		if err := json.Unmarshal(s.Bytes(), &l); err != nil {
			errCh <- err
			break
		}
		jsonlines <- l
	}
	close(jsonlines)
	if err := s.Err(); err != nil {
		errCh <- err
	}
}
