#!/usr/bin/env bash

source "$(dirname "$0")/common.sh"
getProjectDir

echo -e "==> ${CYAN}Bootstrap...${NC}"

if [[ -d ${DIR}/bin ]]; then
	rm -rf ./tmp
fi

echo -e "==> Building & Installing 'bob' the builder"
cd "${DIR}/tools/build" || exit
go install ./cmd/bob
go install github.com/jackc/tern@latest
cd "${DIR}" || exit

echo -e "==> ${YELLOW}Done... ${GREEN}ALL GOOD!!!${NC}"
echo
exit 0
