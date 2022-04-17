#!/usr/bin/env bash

# Get the parent directory of where this script is.
source $(dirname "$0")/common.sh
getProjectDir

# Build a image with src code and debug
docker build . --rm --force-rm --target dev -t hrple:latest-dev

# Done!
echo
echo -e "==> ${YELLOW}Results:${NC}"
docker image ls -f label="app=hrple"
echo