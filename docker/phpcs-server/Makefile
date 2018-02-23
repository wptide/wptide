# PHPCS Server Docker settings.
PHPCS_TAG=${PHPCS_GKE_IMAGE}:${VERSION}

# PHPCS Server GO settings.
PHPCS_BIN_PATH=bin/phpcs-server
PHPCS_PKG_PATH=./cmd/phpcs-server/...

# Show available make commands.
phpcs.usage:
	@echo "\tphpcs.build.bin:\n\t\t- Build the PHPCS Server GO binary."
	@echo "\tphpcs.clean.bin:\n\t\t- Clean the PHPCS Server GO binary."
	@echo "\tphpcs.build.image:\n\t\t- Build the PHPCS Server Docker image."
	@echo "\tphpcs.build.up:\n\t\t- Rebuild & run the PHPCS Server Docker image in isolation with docker-compose up."
	@echo "\tphpcs.up:\n\t\t- Run the PHPCS Server Docker image in isolation with docker-compose up."
	@echo "\tphpcs.down:\n\t\t- Stop the isolated PHPCS Server Docker image with docker-compose down."
	@echo "\tphpcs.push.image:\n\t\t- Push the PHPCS Server Docker image to GCR."
	@echo "\tphpcs.clean.image:\n\t\t- Clean the PHPCS Server Docker image from the host machine."
	@echo "\tphpcs.build.cluster:\n\t\t- Build the PHPCS Server GKE cluster."
	@echo "\tphpcs.creds:\n\t\t- Get the PHPCS Server GKE cluster credentials."
	@echo "\tphpcs.deploy.cluster:\n\t\t- Deploy the PHPCS Server GKE cluster."
	@echo "\tphpcs.get.cluster:\n\t\t- Get the PHPCS Server GKE cluster status."
	@echo "\tphpcs.clean.cluster:\n\t\t- Clean the PHPCS Server GKE cluster."

# Build the PHPCS Server GO binary.
phpcs.build.bin:
	@echo "Building PHPCS Server GO binary '${PHPCS_BIN_PATH}' ..."
	@${GOBUILD} ${LDFLAGS} -o ${PHPCS_BIN_PATH} ${PHPCS_PKG_PATH}

# Clean the PHPCS Server GO binary.
phpcs.clean.bin:
	@echo "Cleaning up PHPCS Server GO binary '${PHPCS_BIN_PATH}' ..."
	@if [ -f ${PHPCS_BIN_PATH} ]; then rm ${PHPCS_BIN_PATH} ; fi

# Build the PHPCS Server Docker image.
phpcs.build.image: deps phpcs.clean.bin phpcs.build.bin
	@echo "Building Docker image [${REPO}/${PHPCS_TAG}] ..."
	@docker build -f docker/phpcs-server/Dockerfile -t ${REPO}/${PHPCS_TAG} --no-cache .
	@docker tag ${REPO}/${PHPCS_TAG} ${PHPCS_GKE_IMAGE}:latest

# Rebuild & run the PHPCS Server Docker image in isolation with docker-compose up.
phpcs.build.up: phpcs.build.image phpcs.up

# Run the PHPCS Server Docker image in isolation with docker-compose up.
phpcs.up:
	@docker-compose up phpcs-server

# Stop the isolated PHPCS Server Docker image with docker-compose down.
phpcs.down:
	@docker-compose down phpcs-server

# Push the PHPCS Server Docker image to GCR.
phpcs.push.image:
	@gcloud docker -- push ${REPO}/${PHPCS_TAG}

# Clean the PHPCS Server Docker image from the host machine.
phpcs.clean.image:
	@docker rmi ${REPO}/${PHPCS_TAG}
	@docker rmi ${PHPCS_GKE_IMAGE}:latest

# Build the PHPCS Server GKE cluster.
phpcs.build.cluster: config
	@gcloud beta container --project ${GCP_PROJECT} clusters create ${PHPCS_GKE_CLUSTER} \
	--zone ${GCP_ZONE} \
	--cluster-version ${PHPCS_GKE_CLUSTER_VERSION} \
	--machine-type ${PHPCS_GKE_MACHINE_TYPE} \
	--image-type "COS" \
	--disk-size ${PHPCS_GKE_DISK_SIZE} \
	--scopes "https://www.googleapis.com/auth/compute","https://www.googleapis.com/auth/devstorage.full_control","https://www.googleapis.com/auth/logging.write","https://www.googleapis.com/auth/monitoring","https://www.googleapis.com/auth/pubsub","https://www.googleapis.com/auth/servicecontrol","https://www.googleapis.com/auth/service.management.readonly","https://www.googleapis.com/auth/trace.append" \
	--preemptible \
	--num-nodes ${PHPCS_GKE_NUM_NODES} \
	--enable-cloud-logging \
	--enable-cloud-monitoring \
	--enable-autoscaling \
	--min-nodes ${PHPCS_GKE_MIN_NODES} \
	--max-nodes ${PHPCS_GKE_MAX_NODES} \
	--enable-legacy-authorization \
	--disable-addons HttpLoadBalancing

# Get the PHPCS Server GKE cluster credentials.
phpcs.creds:
	@gcloud container clusters get-credentials ${PHPCS_GKE_CLUSTER}

# Deploy the PHPCS Server GKE cluster.
phpcs.deploy.cluster: config phpcs.creds
	@kubectl run ${PHPCS_GKE_CLUSTER} --image=${REPO}/${PHPCS_TAG} --replicas=${PHPCS_GKE_REPLICAS} \
	--env="TIDE_API_AUTH_URL=${TIDE_API_AUTH_URL}" \
	--env="TIDE_API_HOST=${TIDE_API_HOST}" \
	--env="TIDE_API_PROTOCOL=${TIDE_API_PROTOCOL}" \
	--env="TIDE_API_KEY=${TIDE_API_KEY}" \
	--env="TIDE_API_SECRET=${TIDE_API_SECRET}" \
	--env="TIDE_API_VERSION=${TIDE_API_VERSION}" \
 	--env="PHPCS_SQS_VERSION=${PHPCS_SQS_VERSION}" \
	--env="PHPCS_SQS_REGION=${PHPCS_SQS_REGION}" \
	--env="PHPCS_SQS_KEY=${PHPCS_SQS_KEY}" \
	--env="PHPCS_SQS_SECRET=${PHPCS_SQS_SECRET}" \
	--env="PHPCS_SQS_QUEUE_NAME=${PHPCS_SQS_QUEUE_NAME}" \
	--env="LH_S3_REGION=${LH_S3_REGION}" \
	--env="LH_S3_KEY=${LH_S3_KEY}" \
	--env="LH_S3_SECRET=${LH_S3_SECRET}" \
	--env="LH_S3_BUCKET_NAME=${LH_S3_BUCKET_NAME}"
	@kubectl autoscale deployment ${PHPCS_GKE_CLUSTER} \
	--cpu-percent=${PHPCS_GKE_CPU_PERCENT} \
	--min=${PHPCS_GKE_MIN_PODS} \
	--max=${PHPCS_GKE_MAX_PODS}

# Get the PHPCS Server GKE cluster status.
phpcs.get.cluster:
	@kubectl get deployment ${PHPCS_GKE_CLUSTER} -o yaml

# Clean the PHPCS Server GKE cluster.
phpcs.clean.cluster: config
	@gcloud container clusters delete ${PHPCS_GKE_CLUSTER} -q
