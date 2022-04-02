#!/usr/bin/env bash
source $(dirname "$0")/common.sh
getProjectDir

echo -e "==> ${CYAN}Clean Up...${NC}"

if [ -d $DIR/bin ]; then
    rm -rf ./tmp
fi

echo -e "==> ${YELLOW}Done... ${GREEN}ALL GOOD!!!${NC}"
echo
exit 0
