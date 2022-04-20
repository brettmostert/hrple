#!/usr/bin/env bash
COMPONENT="/components/hrple-cli"
# Get the parent directory of where this script is.
source ../../scripts/common.sh
getProjectDir

cd ${DIR}

# Build a image with src code and debug
docker build --rm --force-rm --target build -t hrple:latest-dev -f ${DIR}${COMPONENT}/docker/Dockerfile .

# Done!
echo
echo -e "==> ${YELLOW}Results:${NC}"
docker image ls -f label="app=hrple"
echo