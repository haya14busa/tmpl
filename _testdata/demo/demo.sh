#!/bin/bash
DIR=$(dirname "${BASH_SOURCE[0]}")
bash "${DIR}/jsonl.sh" | tmpl -jsonl -t='{{ range $_, $e := . }}{{ .n }}: {{ .mes }}{{nl}}{{end}}'
