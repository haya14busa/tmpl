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
	data, err := os.Open(fmt.Sprintf("%s.json", name))
	if err != nil {
		t.Error(err)
		return
	}
	defer data.Close()
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
	files = append(files, helpers...)
	if err := run(data, out, files, option{}); err != nil {
		t.Error(err)
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
