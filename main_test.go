package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRun(t *testing.T) {
	files, err := filepath.Glob("_testdata/*.tmpl")
	if err != nil {
		t.Fatal(err)
	}
	os.Setenv("TMPL_TEST", "14") // for env test
	for _, f := range files {
		name := strings.TrimSuffix(f, ".tmpl")
		t.Run(name, func(t *testing.T) {
			goldenTest(t, name)
		})
	}
}

func goldenTest(t *testing.T, name string) {
	data, format, err := getInput(t, name)
	outFname := fmt.Sprintf("%s.out", name)
	out, err := os.Create(outFname)
	if err != nil {
		t.Error(err)
		return
	}
	files := []string{fmt.Sprintf("%s.tmpl", name)}
	helpers, err := filepath.Glob(fmt.Sprintf("%s.tmpl.helper*", name))
	if err != nil {
		t.Error(err)
		return
	}
	wantErr := strings.Contains(name, "error")
	files = append(files, helpers...)
	err = run(data, out, files, option{format: format})
	if err != nil && !wantErr {
		t.Error(err)
	} else if err == nil && wantErr {
		t.Error("got nil, but want error")
	}
	out.Close()
	d, err := diff(outFname, fmt.Sprintf("%s.ok", name))
	if err != nil {
		t.Error(err)
	}
	if d != "" {
		t.Errorf("Diff Found:\n%s", d)
	}
}

func getInput(t *testing.T, name string) (file *os.File, format Format, err error) {
	jsonlFile, err := os.Open(fmt.Sprintf("%s.jsonl", name))
	if err == nil {
		t.Cleanup(func() { jsonlFile.Close() })
		return jsonlFile, FormatJSONL, nil
	}
	data, err := os.Open(fmt.Sprintf("%s.json", name))
	if err != nil {
		return nil, FormatJSON, err
	}
	t.Cleanup(func() { data.Close() })
	return data, FormatJSON, nil
}

func diff(f1, f2 string) (string, error) {
	b1, err := ioutil.ReadFile(f1)
	if err != nil {
		return "", err
	}
	b2, err := ioutil.ReadFile(f2)
	if err != nil {
		return "", err
	}
	return cmp.Diff(strings.Split(string(b1), "\n"),
		strings.Split(string(b2), "\n")), nil
}
