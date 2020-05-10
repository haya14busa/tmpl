#!/bin/bash
DIR=$(dirname "${BASH_SOURCE[0]}")
bash "${DIR}/demo.sh" | tee /dev/stderr > "${DIR}/demo.out"
diff -u "${DIR}/demo.out" "${DIR}/demo.ok"
