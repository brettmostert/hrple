#!/usr/bin/env bash
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[0;93m'
CYAN='\033[0;36m'
MAGENTA='\033[0;35m'
GRAY='\033[0;37m'
WHITE='\033[0;97m'

BOLD_GREEN='\033[1;32m'
BOLD_RED='\033[1;31m'
BOLD_YELLOW='\033[1;93m'
BOLD_CYAN='\033[1;36m'
BOLD_MAGENTA='\033[1;35m'
BOLD_GRAY='\033[1;37m'
BOLD_WHITE='\033[1;97m'

NC='\033[0m' # No Color

function getProjectDir() {
    SOURCE="${BASH_SOURCE[0]}"
    while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
    DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"     
}