#!/bin/sh

PROJECT=github.com/DavidSantia/tag_api

# Build for Linux, statically linked

NAME=content-server

echo "## Building $NAME"
docker run --rm --name golang -v $GOPATH/src:/go/src golang:alpine /bin/sh -l -c \
    "cd /go/src/$PROJECT/$NAME; CGO_ENABLED=0 /usr/local/go/bin/go build -i; ls -l $NAME; tar -cf $NAME.tar $NAME"
mv $GOPATH/src/$PROJECT/$NAME/$NAME.tar $NAME

NAME=auth-server

echo "## Building $NAME"
docker run --rm --name golang -v $GOPATH/src:/go/src golang:alpine /bin/sh -l -c \
    "cd /go/src/$PROJECT/$NAME; CGO_ENABLED=0 /usr/local/go/bin/go build -i; ls -l $NAME; tar -cf $NAME.tar $NAME"
mv $GOPATH/src/$PROJECT/$NAME/$NAME.tar $NAME

echo "## Running docker-compose build"
docker-compose build
