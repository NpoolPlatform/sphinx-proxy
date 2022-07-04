#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

PLATFORMS=(
    linux/amd64
    windows/amd64
    darwin/amd64
)
OUTPUT=./output

pkg=github.com/NpoolPlatform/go-service-framework/pkg/version

for PLATFORM in "${PLATFORMS[@]}"; do
    OS="${PLATFORM%/*}"
    ARCH=$(basename "$PLATFORM")

    if git_status=$(git status --porcelain --untracked=no 2>/dev/null) && [[ -z "${git_status}" ]]; then
        git_tree_state=clean
    fi

    git_branch=`git rev-parse --abbrev-ref HEAD`
    set +e
    version=`git describe --tags --abbrev=0`
    if [ ! $? -eq 0 ]; then
        version=$git_branch
    fi
    set -e

    compile_date=`date -u +'%Y-%m-%dT%H:%M:%SZ'`
    git_revision=`git rev-parse HEAD 2>/dev/null || echo unknow`

    service_name=$1

    echo "Building project for $PLATFORM -- $version $compile_date $git_revision"
    CGO_ENABLED=0 GOOS=${OS} GOARCH=${ARCH} go build -v -ldflags "-s -w \
        -X $pkg.buildDate=${compile_date} \
        -X $pkg.gitCommit=${git_revision} \
        -X $pkg.gitVersion=${version}     \
        -X $pkg.gitBranch=${git_branch}"  \
        -o "${OUTPUT}/${OS}/${ARCH}/" "$(pwd)/cmd/$service_name" \
        || exit 1
done
