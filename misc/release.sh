#!/usr/bin/env bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

VERSION=$( cat ${DIR}/../VERSION )

echo "Generating release for proxyhop $VERSION"

export GOPATH=$( cd "${DIR}/../../.." && pwd )
echo "Using GOPATH: ${GOPATH}"

RELEASE_DIR="${GOPATH}/release"
mkdir -p ${RELEASE_DIR}

function build() {
	local os="$1"
	local arch="$2"
	filename="proxyhop_${VERSION}_${os}_${arch}"
	echo "${filename}"
	GOOS=${os} GOARCH=${arch} go build -o ${RELEASE_DIR}/${filename} -ldflags "-X main.version=${VERSION}" proxyhop/cmd/proxyhop
	chmod +x ${RELEASE_DIR}/${filename}
}

build linux amd64
build linux 386
build darwin amd64