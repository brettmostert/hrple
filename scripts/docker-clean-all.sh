#!/usr/bin/env bash

# Removes ALL images for "hrple"
echo
echo "==> Checking hrple images to remove:"
docker image ls -f label="app=hrple"

IMAGES=$(docker image ls -q -f label="app=hrple")

if [[ -n ${IMAGES} ]]; then
	docker image rm -f "${IMAGES}"
fi

# Done!
echo
echo "==> Results:"
docker image ls -f label="app=hrple"
