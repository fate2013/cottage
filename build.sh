#!/bin/bash -e

if [[ $1 = "-loc" ]]; then
    find . -name '*.go' -or -name '*.js' -or -name '*.html' | xargs wc -l | sort -n
    exit
fi

VER=0.1a
#ID=unknown
ID=$(git rev-parse HEAD | cut -c1-7)

if [[ $1 = "-mac" ]]; then
    CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-X github.com/nicholaskh/golib/server.VERSION=$VER -X github.com/nicholaskh/golib/server.BuildID=$ID -w"
    mv cottage bin/cottage.mac
else
    go build -race -ldflags "-X github.com/nicholaskh/golib/server.VERSION=$VER -X github.com/nicholaskh/golib/server.BuildID=$ID -w"
    mv cottage bin/cottage.linux
    bin/cottage.linux -v
fi
