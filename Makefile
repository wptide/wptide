# Copy .env.dist to .env and update.
include .env

# Production environment variables **only** get applied to specific commands.
# Copy .env.dist to .env.prod and update.
ifeq ($(MAKECMDGOALS),$(filter $(MAKECMDGOALS),deploy.cluster))
    ifneq ($(strip $(wildcard .env.prod)),)
        include .env.prod
    endif
endif

# Binary settings.
VERSION=0.1.0
BUILD=`git rev-parse HEAD`

# Docker settings.
REPO=gcr.io/${PROJECT}
TAG=${IMAGE}:${VERSION}

# GO settings.
EXECNAME=lh-server
BINARY=bin/${EXECNAME}
GOOS=linux
GO=`which go`
GOBUILD=CGO_ENABLED=0 GOOS=${GOOS} ${GO} build
GOTEST=${GO} list -f '{{if len .TestGoFiles}}"go test -cover {{.ImportPath}}"{{end}}' ./... | xargs -L 1 sh -c
PACKAGEPATH=./cmd/lh-server/...

# Setup -ldflags for go build.
# Allows setting some global vars before compilation.
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

# Show available make commands.
default:
	@echo "Please supply one of:"
	@echo "\tconfig\t\t- Set GCP configurations."
	@echo "\tcreds\t\t- Get GKE credentials."
	@echo "\tdeps\t\t- Install dependencies."
	@echo "\tbuild.bin\t- Build the GO binary."
	@echo "\tclean.bin\t- Clean the GO binary."
	@echo "\tbuild.image\t- Build the Docker image."
	@echo "\tbuild.run.image\t- Rebuild & run the Docker image with docker-compose up."
	@echo "\trun.image\t- Run the Docker image with docker-compose up."
	@echo "\tstop.image\t- Stop the Docker image with docker-compose down."
	@echo "\tpush.image\t- Push the Docker image to GCR."
	@echo "\tclean.image\t- Clean the Docker image from the host machine."
	@echo "\tbuild.cluster\t- Build the GKE cluster."
	@echo "\tdeploy.cluster\t- Deploy the GKE cluster."
	@echo "\tclean.cluster\t- Clean the GKE cluster."
	@echo "\ttest\t\t- Run the go test suite."

# Set GCP configurations.
config:
	@gcloud config set project ${PROJECT}
	@gcloud config set compute/zone ${ZONE}
	@gcloud config set container/new_scopes_behavior true

# Get GKE credentials.
creds:
	@gcloud container clusters get-credentials ${CLUSTER}

# Install dependencies.
deps:
	@echo "Installing dependencies ..."
	@glide install

# Build the GO binary.
build.bin:
	@echo "Building ${BINARY} ..."
	@${GOBUILD} ${LDFLAGS} -o ${BINARY} ${PACKAGEPATH}

# Clean the GO binary.
clean.bin:
	@echo "Cleaning up binary ..."
	@if [ -f ${BINARY} ]; then rm ${BINARY} ; fi

# Build the Docker image.
build.image: deps clean.bin build.bin
	@echo "Building Docker image [${REPO}/${TAG}] ..."
	@docker build -t ${REPO}/${TAG} --no-cache .
	@docker tag ${REPO}/${TAG} ${IMAGE}:latest

# Rebuild & run the Docker image with docker-compose up.
build.run.image: build.image run.image

# Run the Docker image with docker-compose up.
run.image:
	@docker-compose up

# Stop the Docker image with docker-compose down.
stop.image:
	@docker-compose down

# Push the Docker image to GCR.
push.image:
	@gcloud docker -- push ${REPO}/${TAG}

# Clean the Docker image from the host machine.
clean.image:
	@docker rmi ${REPO}/${TAG}
	@docker rmi ${IMAGE}:latest

# Build the GKE cluster.
build.cluster: config
	@gcloud beta container --project ${PROJECT} clusters create ${CLUSTER} \
	--zone ${ZONE} \
	--cluster-version ${CLUSTER_VERSION} \
	--machine-type ${MACHINE_TYPE} \
	--image-type "COS" \
	--disk-size ${DISK_SIZE} \
	--scopes "https://www.googleapis.com/auth/compute","https://www.googleapis.com/auth/devstorage.full_control","https://www.googleapis.com/auth/logging.write","https://www.googleapis.com/auth/monitoring","https://www.googleapis.com/auth/pubsub","https://www.googleapis.com/auth/servicecontrol","https://www.googleapis.com/auth/service.management.readonly","https://www.googleapis.com/auth/trace.append" \
	--preemptible \
	--num-nodes ${NUM_NODES} \
	--enable-cloud-logging \
	--enable-cloud-monitoring \
	--enable-autoscaling \
	--min-nodes ${MIN_NODES} \
	--max-nodes ${MAX_NODES} \
	--enable-legacy-authorization \
	--disable-addons HttpLoadBalancing

# Deploy the GKE cluster.
deploy.cluster: config creds
	@kubectl run ${CLUSTER} --image=${REPO}/${TAG} --replicas=${REPLICAS} \
	--env="TIDE_API_AUTH_URL=${TIDE_API_AUTH_URL}" \
    --env="TIDE_API_HOST=${TIDE_API_HOST}" \
    --env="TIDE_API_PROTOCOL=${TIDE_API_PROTOCOL}" \
    --env="TIDE_API_KEY=${TIDE_API_KEY}" \
    --env="TIDE_API_SECRET=${TIDE_API_SECRET}" \
    --env="TIDE_API_VERSION=${TIDE_API_VERSION}" \
    --env="AWS_SQS_LH_VERSION=${AWS_SQS_LH_VERSION}" \
    --env="AWS_SQS_LH_REGION=${AWS_SQS_LH_REGION}" \
    --env="AWS_SQS_LH_KEY=${AWS_SQS_LH_KEY}" \
    --env="AWS_SQS_LH_SECRET=${AWS_SQS_LH_SECRET}" \
    --env="AWS_SQS_LH_QUEUE_NAME=${AWS_SQS_LH_QUEUE_NAME}" \
    --env="AWS_S3_REGION=${AWS_S3_REGION}" \
    --env="AWS_S3_KEY=${AWS_S3_KEY}" \
    --env="AWS_S3_SECRET=${AWS_S3_SECRET}" \
    --env="AWS_S3_BUCKET_NAME=${AWS_S3_BUCKET_NAME}" \
    @kubectl autoscale deployment ${CLUSTER} \
    --cpu-percent=${CPU_PERCENT} \
    --min=${MIN_PODS} \
    --max=${MAX_PODS}

# Clean the GKE cluster.
clean.cluster: config
	@gcloud container clusters delete ${CLUSTER} -q

# Run the GO test suite.
test:
	@echo "Running tests ..."
	@${GOTEST}
