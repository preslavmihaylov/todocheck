#!/bin/bash

DIR=$1
VERSION=$2

if [ -z "$DIR" ] || [ -z "$VERSION" ]; then 
    echo "Usage: release.sh <target-dir> <version>"
    exit 1
fi

function create_build {
    GOOS=$1
    GOARCH=$2
    EXT=$3
    if [ -z $EXT ]; then
        EXT=$GOARCH
    fi

    BINARY=todocheck-$VERSION-$GOOS-$EXT
    GOOS=$GOOS GOARCH=$GOARCH go build -ldflags="-X main.version=$VERSION" -o $DIR/$BINARY
    cd $DIR && shasum -a 256 $BINARY > $BINARY.sha256 && cd ..
}

mkdir -p $DIR
create_build windows amd64 x86_64.exe
create_build darwin amd64 x86_64
create_build linux amd64 x86_64
create_build linux arm64
# foo
