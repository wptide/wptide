#!/bin/bash

## Output styles.
BOLD=$(tput bold)
RED=$(tput setaf 1)
YELLOW=$(tput setaf 3)
CYAN=$(tput setaf 6)
RESET=$(tput sgr0)

## CLI messages.
HELP="Pass [${CYAN}--help${RESET}|${CYAN}-h${RESET}] as a cli argument for help using this script."
ERROR="${RED}${BOLD}Error${RESET}:"
USAGE="${BOLD}Usage${RESET}:
    $0 {${RED}template${RESET}} {${RED}destination${RESET}} {${YELLOW}environment${RESET}}
${BOLD}Notes${RESET}:${CYAN}
    1. The default ${RESET}${BOLD}.env${RESET}${CYAN} file is always sourced.
    2. All file paths are relative to the project ${RESET}${BOLD}root${RESET}${CYAN}.${RESET}
${BOLD}Arguments${RESET}:
    ${CYAN}template${RESET}:
        The template file used to generate the destination file. (${RED}required${RESET})
    ${CYAN}destination${RESET}:
        The destination file that is generated from the template file. (${RED}required${RESET})
    ${CYAN}environment${RESET}:
        The environment file used to source custom environment variables. (${YELLOW}optional${RESET})"

## Handle help parameters.
for p in "$@";
do
    case $p in
    --help|-h)
        echo "${USAGE}"
        exit 0;
        ;;
    esac
done

## Validate cli arguments.
if [ $# -lt 1 ]; then
    echo "${ERROR}
    Missing required {${RED}template${RESET}} and {${RED}destination${RESET}} arguments.
    ${HELP}"
    exit 1
elif [ $# -lt 2 ]; then
    echo "${ERROR}
    Missing required {${RED}destination${RESET}} argument.
    ${HELP}"
    exit 1
fi

## Check for envsubst command.
if ! command -v envsubst >/dev/null 2>&1; then
    echo "${ERROR}
    Variable Interpolation failure. The ${RED}envsubst${RESET} command is not installed on your OS."
    exit 1
fi

## Validate template file.
if [ ! -e ${1} ]; then
    echo "${ERROR}
    The template file (${RED}${BOLD}${1}${RESET}) does not exist at the location provided.
    ${HELP}"
    exit 1
fi

# Source .env file.
if [ -e .env ]; then
    export $(cat .env | grep -v ^# | xargs)
fi

## Source custom environment file.
if [ ! -z ${3} ] && [ -e ${3} ]; then
    export $(cat ${3} | grep -v ^# | xargs)
fi

## Substitute environment variables and generate destination file.
envsubst < "$1" > "$2"