#!/usr/bin/env bash

# Get the parent directory of where this script is.
source $(dirname "$0")/common.sh
getProjectDir

# Set module download mode to readonly to not implicitly update go.mod
export GOFLAGS="-mod=readonly"

# In release mode we don't want debug information in the binary
if [[ -n "${TRPLE_RELEASE}" ]]; then
    LD_FLAGS="-s -w"
fi

# Create bin if it's doesn't exists
if [ ! -d $DIR/bin ]; then
    echo "==> Creating ${DIR}/bin directory..."
    mkdir -p $DIR/bin
fi

# Ensure all remote modules are downloaded and cached before build so that
# gox is not used yet
# the concurrent builds if used... launched by gox won't race to redundantly download them.
echo -e "==> ${CYAN}Installing remote modules...${NC}"
go mod download

# Build!
echo -e "==> ${CYAN}Building...${NC}"
CGO_ENABLED=0 GOOS=linux go build -o $DIR/bin -ldflags "${LD_FLAGS}" ./...
GOPATH=${GOPATH:-$(go env GOPATH)}
MAIN_GOPATH=($GOPATH)

# Create GOPATH/bin if it's doesn't exists
if [ ! -d $MAIN_GOPATH/bin ]; then
    echo "==> Creating GOPATH/bin directory..."
    mkdir -p $MAIN_GOPATH/bin
fi

for F in $(find $DIR/bin -mindepth 1 -maxdepth 1 -type f); do
    cp ${F} ${MAIN_GOPATH}/bin/
done

# Done!
echo
echo -e "==> ${YELLOW}Results:${NC}"
ls -hl $DIR/bin
echo
