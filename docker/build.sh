#!/bin/sh

PROJECT=github.com/DavidSantia/tag_api

if ! [ -d $GOPATH/src/$PROJECT ]; then
    echo "Project $PROJECT not found"
    exit
fi

APPS=`cd $GOPATH/src/$PROJECT/apps; ls -1`

# Build for Linux, statically linked
build () {
    echo "## Building $1"
    cd $GOPATH/src/$PROJECT/apps/$1
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 govvv build -o $GOPATH/src/$PROJECT/docker/$1/$1
    ls -l $GOPATH/src/$PROJECT/docker/$1
}

for i in $APPS
  do build api-server
done
