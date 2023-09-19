#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
CODEGEN_PKG=${GOPATH}/src/k8s.io/code-generator

source "${CODEGEN_PKG}/kube_codegen.sh"

#kube::codegen::gen_helpers \
#    --input-pkg-root github.com/Dimss/first/apis \
#    --output-base "${GOPATH}/src" \
#    --boilerplate "${SCRIPT_ROOT}/hack/boilerplate.go.txt"

kube::codegen::gen_client \
    --with-watch \
    --input-pkg-root github.com/Dimss/first/apis \
    --output-pkg-root github.com/Dimss/first/pkg/generated \
    --output-base "${GOPATH}/src" \
    --boilerplate "${SCRIPT_ROOT}/hack/boilerplate.go.txt"

#CODE_GENERATOR_EXEC_DIR="${GOPATH}/src/github.com/code-generator"
#GENERATE_GROUPS="${CODE_GENERATOR_EXEC_DIR}/generate-groups.sh"
#GENERATE_INTERNAL_GROUPS="${CODE_GENERATOR_EXEC_DIR}/generate-internal-groups.sh"
#
#$GENERATE_GROUPS "applyconfiguration,client,deepcopy,informer,lister" \
# github.com/Dimss/first/pkg/generated \
# github.com/Dimss/first/apis "metafirst:v1alpha1" \
# --go-header-file "hack/boilerplate.go.txt"
#
#$GENERATE_INTERNAL_GROUPS "deepcopy,defaulter,conversion,openapi" \
# github.com/Dimss/first/pkg/generated \
# github.com/Dimss/first/apis \
# github.com/Dimss/first/apis "metafirst:v1alpha1" \
# --go-header-file "hack/boilerplate.go.txt"
