#!/usr/bin/env bash

source "$(dirname "$0")/common.sh"
getProjectDir

echo -e "==> ${CYAN}Bootstrap...${NC}"

if [[ -d ${DIR}/bin ]]; then
	rm -rf ./tmp
fi

echo -e "==> Clean Up"
if [[ -d ${DIR}/tools/build ]]; then
	rm -rf ./tools/build
fi

echo -e "==> Building & Installing 'bob' the builder"
cd "${DIR}/go/tools/build" || exit
go install ./cmd/bob
cd "${DIR}" || exit

echo -e "==> ${YELLOW}Done... ${GREEN}ALL GOOD!!!${NC}"
echo
exit 0
