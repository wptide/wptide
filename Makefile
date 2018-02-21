# Copy .env.dist to .env and update.
-include .env

# Production environment variables **only** get applied to specific commands.
# Copy .env.dist to .env.prod and update.
ifeq ($(MAKECMDGOALS),$(filter $(MAKECMDGOALS),lighthouse.deploy.cluster))
    ifneq ($(strip $(wildcard .env.prod)),)
        include .env.prod
    endif
endif

# Binary settings.
VERSION=0.1.0
BUILD=`git rev-parse HEAD`

# Docker settings.
REPO=gcr.io/${PROJECT}

# Lighthouse Docker settings.
LH_TAG=${LH_IMAGE}:${VERSION}

# GO settings.
GOOS=linux
GO=`which go`
GOBUILD=CGO_ENABLED=0 GOOS=${GOOS} ${GO} build
GOTEST=${GO} list -f '{{if len .TestGoFiles}}"go test -cover {{.ImportPath}}"{{end}}' ./... | xargs -L 1 sh -c

# Lighthouse GO settings.
LH_BIN_PATH=bin/lighthouse-server
LH_PKG_PATH=./cmd/lighthouse-server/...

# Setup -ldflags for go build.
# Allows setting some global vars before compilation.
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

# Show available make commands.
usage:
	@echo "Please supply one of:"
	@echo "\tdeps:\n\t\t- Install dependencies."
	@echo "\tconfig:\n\t\t- Set GCP configurations."
	@echo "\tbuild.bins\n\t\t- Build all the GO binaries."
	@echo "\tclean.bins\n\t\t- Clean all the GO binaries."
	@echo "\tbuild.images\n\t\t- Build all the Docker images."
	@echo "\tbuild.up\n\t\t- Rebuild & run the Docker images with docker-compose up."
	@echo "\tup\n\t\t- Run the Docker images with docker-compose up."
	@echo "\tdown\n\t\t- Stop the Docker images with docker-compose down."
	@echo "\ttest:\n\t\t- Run the GO test suite."
	@make lighthouse.usage

# Include Makefiles.
include docker/lighthouse-server/Makefile

# Install dependencies.
deps:
	@echo "Installing dependencies ..."
	@glide install

# Set GCP configurations.
config:
	@gcloud config set project ${PROJECT}
	@gcloud config set compute/zone ${ZONE}
	@gcloud config set container/new_scopes_behavior true

# Build all the GO binaries.
build.bins: lighthouse.build.bin

# Clean all the GO binaries.
clean.bins: lighthouse.clean.bin

# Build all the Docker images.
build.images: deps clean.bins build.bins lighthouse.build.image

# Rebuild & run the Docker images with docker-compose up.
build.up: build.images up

# Run the Docker images with docker-compose up.
up:
	@docker-compose up

# Stop the Docker images with docker-compose down.
down:
	@docker-compose down

# Run the GO test suite.
test:
	@echo "Running tests ..."
	@${GOTEST}
