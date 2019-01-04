#!/bin/sh

echo "## Cleaning containers"
CONTAINERS=`docker ps -a -q`
if [ -n "$CONTAINERS" ]; then
   docker rm $CONTAINERS
fi

echo "## Remove unused images"
UNUSED=`docker images | grep none | awk '{print $3}'`
if [ -n "$UNUSED" ]; then
   docker rmi $UNUSED
fi
