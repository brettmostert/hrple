#!/usr/bin/env bash
source $(dirname "$0")/common.sh
getProjectDir

echo -e "==> ${CYAN}Checking that code complies with static analysis requirements...${NC}"

OUTPUT=$(golangci-lint run --color=always)

if [[ -z "${OUTPUT}" ]]; then
    echo -e "==> ${YELLOW}Result... ${GREEN}ALL GOOD!!!${NC}"
    echo
    exit 0
else
    echo -e "==> ${YELLOW}Result... ${RED}BAD!!!${NC}"
    echo $OUTPUT
    echo
    exit 1
fi