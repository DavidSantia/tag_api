#!/bin/sh

# Clean previous
echo "## Cleaning images"
CONTAINERS="`docker ps -a -f ancestor=tagdemo/content-server -f ancestor=tagdemo/auth-server -f ancestor=tagdemo/mysql -q`"
if [ -n "$CONTAINERS" ]; then
   docker stop $CONTAINERS
   docker rm $CONTAINERS
fi

IMAGES="`docker images tagdemo/content-server -q; docker images tagdemo/auth-server -q; docker images tagdemo/mysql -q`"
if [ -n "$IMAGES" ]; then
   docker rmi $IMAGES
fi
