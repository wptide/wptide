#!/bin/bash

TEMPLATE="$1"
DESTINATION="$2"

if [ -e .env ]; then
    export $(cat .env | grep -v ^# | xargs)
fi

if [ -e .env.prod ]; then
    export $(cat .env.prod | grep -v ^# | xargs)
fi

if ! command -v envsubst >/dev/null 2>&1; then
    echo "Variable Interpolation failure: envsubst is not installed"
    exit 1
fi

if [ -z ${1} ]; then
    echo "The template file is required."
    exit 1
elif [ ! -e ${1} ]; then
    echo "The template file does not exist."
    exit 1
fi

if [ -z ${2} ]; then
    echo "The destination file is required."
    exit 1
fi

envsubst < "$1" > "$2"