#!/bin/sh

PROJECT=github.com/DavidSantia/tag_api

if ! [ -d $GOPATH/src/$PROJECT ]; then
    echo "Project $PROJECT not found"
    exit
fi

# Build for Linux, statically linked
build () {
    cd $GOPATH/src/$PROJECT/apps/$1
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 govvv build -o $GOPATH/src/$PROJECT/docker/$1/$1
    ls -l $GOPATH/src/$PROJECT/docker/$1
}

# Build apps
cd $GOPATH/src/$PROJECT/apps
APPS=`ls -d */ | sed 's+/$++'`
for i in $APPS
  do echo "## Building app $i"
  build $i
done

# Build docker images
cd $GOPATH/src/$PROJECT/docker
IMAGES=`ls -d */ | sed 's+/$++'`
for i in $IMAGES
  do echo "## Building docker image tagdemo/$i"
  docker build -t tagdemo/$i ./$i
done
