# Copy .env.dist to .env and update.
-include .env

# Binary settings.
VERSION=1.0.0-beta
BUILD=`git rev-parse HEAD`

# GO settings.
GOOS=linux
GO=`which go`
GOBUILD=CGO_ENABLED=0 GOOS=${GOOS} ${GO} build
GOTEST=${GO} list -f '{{if len .TestGoFiles}}"go test -cover {{.ImportPath}}"{{end}}' ./... | xargs -L 1 sh -c

# Setup -ldflags for go build.
# Allows setting some global vars before compilation.
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

# Show available make commands.
usage:
	@echo "Please supply one of:"
	@echo "\tdeps:\n\t\t- Install dependencies."
	@echo "\tconfig:\n\t\t- Set GCP configurations."
	@echo "\tbuild.bins:\n\t\t- Build all the GO binaries."
	@echo "\tclean.bins:\n\t\t- Clean all the GO binaries."
	@echo "\tbuild.images:\n\t\t- Build all the Docker images."
	@echo "\tbuild.up:\n\t\t- Rebuild & run the Docker images with docker-compose up."
	@echo "\tup:\n\t\t- Run the Docker images with docker-compose up."
	@echo "\tdown:\n\t\t- Stop the Docker images with docker-compose down."
	@echo "\ttest:\n\t\t- Run the GO test suite."
	@make api.usage
	@make lighthouse.usage
	@make phpcs.usage
	@make sync.usage

# Include Makefiles.
include service/api/Makefile
include service/lighthouse-server/Makefile
include service/phpcs-server/Makefile
include service/sync-server/Makefile

# Install dependencies.
deps:
	@echo "Installing dependencies ..."
	@glide install

# Set GCP configurations.
config:
	@gcloud config set project ${GCP_PROJECT}
	@gcloud config set compute/zone ${GCP_ZONE}
	@gcloud config set container/new_scopes_behavior true

# Build all the GO binaries.
build.bins: lighthouse.build.bin phpcs.build.bin sync.build.bin

# Clean all the GO binaries.
clean.bins: lighthouse.clean.bin phpcs.clean.bin sync.clean.bin

# Build all the Docker images.
build.images: deps clean.bins build.bins lighthouse.build.image phpcs.build.image sync.build.image

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
