package main

import (
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

type option struct {
	version bool
	tmpl    string
}

func setupFlags(opt *option) {
	flag.BoolVar(&opt.version, "version", false, "print version")
	flag.StringVar(&opt.tmpl, "t", "", "Go text/template text template")

	flag.Usage = usage
	flag.Parse()
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: tmpl [FLAGS] [Template Files]")
	fmt.Fprintln(os.Stderr, "\tGenerate textual output using Go text/template from given data in STDIN")
	fmt.Fprintln(os.Stderr, "Flags:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "GitHub: https://github.com/haya14busa/tmpl")
	fmt.Fprintln(os.Stderr, "Syntax: https://golang.org/pkg/text/template")
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
	var d map[string]interface{}
	if err := json.NewDecoder(r).Decode(&d); err != nil {
		return err
	}
	t, err := buildTemplate(args, opt)
	if err != nil {
		return fmt.Errorf("failed to build template: %v", err)
	}
	return t.Execute(w, d)
}

func buildTemplate(files []string, opt option) (*template.Template, error) {
	funcMap := template.FuncMap{
		"env":     os.Getenv,
		"strings": funcs.Strings,
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
