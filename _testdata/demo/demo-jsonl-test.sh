#!/bin/bash
DIR=$(dirname "${BASH_SOURCE[0]}")
bash "${DIR}/demo-jsonl.sh" | tee /dev/stderr > "${DIR}/demo-jsonl.out"
diff -u "${DIR}/demo-jsonl.out" "${DIR}/demo-jsonl.ok"
