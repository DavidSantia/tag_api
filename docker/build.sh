#!/bin/sh

PROJECT=github.com/DavidSantia/tag_api

if ! [ -d $GOPATH/src/$PROJECT ]; then
    echo "Project $PROJECT not found"
    exit
fi

# Clean previous
$GOPATH/src/$PROJECT/docker/clean.sh

IMAGES="`docker images tagdemo/content-server -q; docker images tagdemo/auth-server -q; docker images tagdemo/mysql -q`"
if [ -n "$IMAGES" ]; then
   docker rmi $IMAGES
fi

# Build for Linux, statically linked
build () {
    echo "## Building $1"
    cd $GOPATH/src/$PROJECT/apps/$1
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 govvv build
    ls -l $1
    tar -cf $1.tar $1
    mv $1.tar $GOPATH/src/$PROJECT/docker/$1/
}

build content-server

build auth-server
