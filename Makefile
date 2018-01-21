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

# Show available make subcommands
default:
	@echo "Please supply one of:"
	@echo "\tdep\t- Dependencies"
	@echo "\tclean\t- Cleanup builds"
	@echo "\tbuild\t- Build binaries"
	@echo "\tlocal\t- Build binaries to run locally"
	@echo "\tpackage\t- Build the Docker image"
	@echo "\ttest\t- Run test suite"
	@echo "\tall\t- Do everything"

# Install dependencies. Do this first!
dep:
	@echo "Getting dependencies ..."
	@glide install

# Clean all built files.
clean:
	@echo "Cleaning up ..."
	@if [ -f ${BINARY} ]; then rm ${BINARY} ; fi
	@if [ -f ${LOCALBIN} ]; then rm ${LOCALBIN} ; fi

# Build the binary.
build:
	@echo "Building ${BINARY} ..."
	@${GOBUILD} ${LDFLAGS} -o ${BINARY} ${PACKAGEPATH}

# Build a local binary that can run on local machine.
local:
	@echo "Building ${LOCALBIN} ..."
	@${GOBUILDLOCAL} ${LDFLAGS} -o ${LOCALBIN} ${PACKAGEPATH}

# Package and build the Docker image.
package: clean build
	@echo "Building Docker image [${TAG}] ..."
	@docker build -t ${TAG} --no-cache .
	@docker tag ${TAG} ${DOCKERNAME}:latest

# Run the image with docker-compose.
run:
	@docker-compose up

# Force rebuilding of image before running.
up: package run

# Run the go test suite.
test:
	@echo "Running tests ..."
	@${GOTEST}

# Do it all!
all: clean dep build local up