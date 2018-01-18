# @TODO: Perhaps move some of these to another Makefile.
EXECNAME=lh-server
BINARY=bin/${EXECNAME}
GOOS=linux
LOCAL=local
LOCALBIN=bin/${LOCAL}/${EXECNAME}
GO=`which go`
GOBUILD=CGO_ENABLED=0 GOOS=${GOOS} ${GO} build
GOTEST=${GO} test -cover ./src/...
GOBUILDLOCAL=${GO} build
PACKAGEPATH=./cmd/lh-server/...
CERTFILE=certs/cacert.pem

# Copy Sample.properties to Makefile.properties and update.
include .env

# Binary version.
VERSION=0.0.1
BUILD=`git rev-parse HEAD`

DOCKERNAME=wptide/lighthouse
TAG=${DOCKERNAME}:${VERSION}

# Setup -ldflags for go build.
# Allows setting some global vars before compilation.
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

default:
	@echo "Please supply one of:"
	@echo "\tdep\t- Dependencies"
	@echo "\tclean\t- Cleanup builds"
	@echo "\tcert\t- Get root certificates for SSL/TSL"
	@echo "\tbuild\t- Build binaries"
	@echo "\tlocal\t- Build binaries to run locally"
	@echo "\tpackage\t- Build the Docker image"
	@echo "\ttest\t- Run test suite"
	@echo "\tall\t- Do everything"

dep:
	@echo "Getting dependencies ..."
	@glide install

clean:
	@echo "Cleaning up ..."
	@if [ -f ${BINARY} ]; then rm ${BINARY} ; fi
	@if [ -f ${LOCALBIN} ]; then rm ${LOCALBIN} ; fi
	@if [ -f ${CERTFILE} ]; then rm ${CERTFILE} ; fi

cert:
	@echo "Getting root certs for SSL/TLS ..."
	@curl --output certs/cacert.pem https://curl.haxx.se/ca/cacert.pem -s

build:
	@echo "Building ${BINARY} ..."
	@${GOBUILD} ${LDFLAGS} -o ${BINARY} ${PACKAGEPATH}

local:
	@echo "Building ${LOCALBIN} ..."
	@${GOBUILDLOCAL} ${LDFLAGS} -o ${LOCALBIN} ${PACKAGEPATH}

package: clean cert build
	@echo "Building Docker image [${TAG}] ..."
	@docker build -t ${TAG} --no-cache .
	@docker tag ${TAG} ${DOCKERNAME}:latest

test:
	@echo "Running tests ..."
	@${GOTEST}

all: clean cert dep build local package