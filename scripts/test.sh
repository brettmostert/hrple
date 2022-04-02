#!/usr/bin/env bash
source $(dirname "$0")/common.sh
getProjectDir

echo -e "==> ${CYAN}Running Tests...${NC}"

if [ ! -d $DIR/tmp ]; then
    echo "==> Creating ${DIR}/tmp directory..."
    mkdir -p $DIR/tmp
fi

OUTPUT=$(go test ./... -cover -race -covermode=atomic -coverprofile=./tmp/coverage.out)
status=$?

if [[ $status -eq 0 ]]; then
    echo -e "$OUTPUT"
    echo -e "==> ${YELLOW}Result... ${GREEN}ALL GOOD!!!${NC}"
    echo
    exit 0
else
    echo -e "==> ${YELLOW}Result... ${RED}BAD!!!${NC}"
    echo $OUTPUT
    echo
    exit 1
fi