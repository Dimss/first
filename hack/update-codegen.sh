#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
CODEGEN_PKG=${GOPATH}/src/k8s.io/code-generator
API_KNOWN_VIOLATIONS_DIR="${SCRIPT_ROOT}/hack"

source "${CODEGEN_PKG}/kube_codegen.sh"

kube::codegen::gen_helpers \
  --input-pkg-root github.com/Dimss/first/apis \
  --output-base "${GOPATH}/src" \
  --boilerplate "${SCRIPT_ROOT}/hack/boilerplate.go.txt"

kube::codegen::gen_client \
  --with-watch \
  --input-pkg-root github.com/Dimss/first/apis \
  --output-pkg-root github.com/Dimss/first/pkg/generated \
  --output-base "${GOPATH}/src" \
  --boilerplate "${SCRIPT_ROOT}/hack/boilerplate.go.txt"

if [[ -n "${API_KNOWN_VIOLATIONS_DIR:-}" ]]; then
  report_filename="${API_KNOWN_VIOLATIONS_DIR}/violation_exceptions.list"
  if [[ "${UPDATE_API_KNOWN_VIOLATIONS:-}" == "true" ]]; then
    update_report="--update-report"
  fi
fi

kube::codegen::gen_openapi \
  --input-pkg-root github.com/Dimss/first/apis \
  --output-pkg-root github.com/Dimss/first/pkg/generated \
  --output-base "${GOPATH}/src" \
  --report-filename "${report_filename:-"/dev/null"}" \
  ${update_report:+"${update_report}"} \
  --boilerplate "${SCRIPT_ROOT}/hack/boilerplate.go.txt"
