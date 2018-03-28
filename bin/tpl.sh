#!/bin/bash

## Output styles.
BOLD=$(tput bold)
RED=$(tput setaf 1)
YELLOW=$(tput setaf 3)
CYAN=$(tput setaf 6)
RESET=$(tput sgr0)

## CLI messages.
USAGE="usage: $0 [<options>]

    -h, --help      help with script usage
    -p, --path      path to the template file (${RED}required${RESET})
    -f, --file      name of the .tpl file (without extension) (${RED}required${RESET})
    -e, --env       environment file used to source custom environment variables (${YELLOW}optional${RESET})

${CYAN}The default ${RESET}${BOLD}.env${RESET}${CYAN} file is always sourced and all file paths are relative to the project ${RESET}${BOLD}root${RESET}${CYAN}.${RESET}"


## Handle help parameters.
for p in "$@";
do
    case $p in
    --help|-h)
        echo "${USAGE}"
        exit 0
        ;;
    -p=*|--path=*)
        TEMP_PATH="${p#*=}"
        ;;
    -f=*|--filename=*)
        TEMP_FILENAME="${p#*=}"
        ;;
    -e=*|--env=*)
        ENV_FILE="${p#*=}"
        ;;
    esac
done

## Check for envsubst command.
if ! command -v envsubst >/dev/null 2>&1; then
    echo "${RED}error${RESET}: Variable Interpolation failure. The ${CYAN}envsubst${RESET} command is not installed on your OS."
    exit 1
fi

TEMP_FILE="${TEMP_PATH}${TEMP_FILENAME}.tpl"
DEST_FILE="${TEMP_PATH}${TEMP_FILENAME}.yaml"

## Validate cli arguments.
if [ "$#" -eq 0 ]; then
    echo "${USAGE}"
    exit 1
elif [ -z ${TEMP_PATH} ]; then
    echo "${RED}error${RESET}: Missing required [path] argument. Example [-p=<dir-path/>] or [--path=<dir-path/>]."
    exit 1
elif [ -z ${TEMP_FILENAME} ]; then
    echo "${RED}error${RESET}: Missing required [filename] argument. Example [-f=<filename>] or [--filename=<filename>]."
    exit 1
elif [ ! -e ${TEMP_FILE} ]; then
    echo "${RED}error${RESET}: The template file [${CYAN}${TEMP_FILE}${RESET}] does not exist at the location provided."
    exit 1
fi

# Source .env file.
if [ -e .env ]; then
    export $(cat .env | grep -v ^# | xargs)
fi

## Source custom environment file.
if [ ! -z ${ENV_FILE} ]; then
    if [ -e ${ENV_FILE} ]; then
        export $(cat ${ENV_FILE} | grep -v ^# | xargs)
    else
        echo "${RED}error${RESET}: The environment file [${CYAN}${ENV_FILE}${RESET}] does not exist at the location provided."
    fi
fi

## Substitute environment variables and generate destination file.
envsubst < "$TEMP_FILE" > "$DEST_FILE"