#!/bin/bash

## Output styles.
BOLD=$(tput bold)
RED=$(tput setaf 1)
YELLOW=$(tput setaf 3)
CYAN=$(tput setaf 6)
RESET=$(tput sgr0)

## CLI messages.
USAGE="usage: $0 [<options>]

    -h, --help          help with script usage
    -t, --template      path to the template file (${RED}required${RESET})
    -d, --destination   path to the destination file (${RED}required${RESET})

All paths are relative to the project ${RESET}${BOLD}root${RESET}${CYAN}.${RESET}"

## Handle cli parameters.
for p in "$@";
do
    case $p in
    --help|-h)
        echo "${USAGE}"
        exit 0
        ;;
    -t=*|--template=*)
        TEMP_FILE="${p#*=}"
        ;;
    -d=*|--destination=*)
        DEST_FILE="${p#*=}"
        ;;
    esac
done

## Validate cli arguments.
if [ "$#" -eq 0 ]; then
    echo "${USAGE}"
    exit 1
elif [ -z ${TEMP_FILE} ]; then
    echo "${RED}error${RESET}: Missing required [template] argument. Example [-t=<dir-path/file.tpl>] or [--template=<dir-path/file.tpl>]."
    exit 1
elif [ -z ${DEST_FILE} ]; then
    echo "${RED}error${RESET}: Missing required [destination] argument. Example [-d=<dir-path/file.yaml>] or [--destination=<dir-path/file.yaml>]."
    exit 1
elif [ ! -e ${TEMP_FILE} ]; then
    echo "${RED}error${RESET}: The template file [${CYAN}${TEMP_FILE}${RESET}] does not exist at the location provided."
    exit 1
fi

## WordPress unique Keys and Salts.
uniqueKeys=(
    AUTH_KEY
    SECURE_AUTH_KEY
    LOGGED_IN_KEY
    NONCE_KEY
    AUTH_SALT
    SECURE_AUTH_SALT
    LOGGED_IN_SALT
    NONCE_SALT
)

## Export the unique Keys and Salts.
for unique in "${uniqueKeys[@]}"; do
    if [ -z ${!unique+x} ]; then
        export "${unique}=`LC_CTYPE=C; cat /dev/urandom | tr -dc [:print:] | tr -d '[:space:]\042\047\134' | head -c 64`"
    fi
done

## Substitute environment variables and generate destination file.
envsubst "`printf '${%s} ' $(sh -c "env|cut -d'=' -f1")`" < "$TEMP_FILE" > "$DEST_FILE"