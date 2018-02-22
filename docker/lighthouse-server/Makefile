# Lighthouse Server Docker settings.
LH_TAG=${LH_IMAGE}:${VERSION}

# Lighthouse Server GO settings.
LH_BIN_PATH=bin/lighthouse-server
LH_PKG_PATH=./cmd/lighthouse-server/...

# Show available make commands.
lighthouse.usage:
	@echo "\tlighthouse.build.bin:\n\t\t- Build the Lighthouse Server GO binary."
	@echo "\tlighthouse.clean.bin:\n\t\t- Clean the Lighthouse Server GO binary."
	@echo "\tlighthouse.build.image:\n\t\t- Build the Lighthouse Server Docker image."
	@echo "\tlighthouse.build.up:\n\t\t- Rebuild & run the Lighthouse Server Docker image in isolation with docker-compose up."
	@echo "\tlighthouse.up:\n\t\t- Run the Lighthouse Server Docker image in isolation with docker-compose up."
	@echo "\tlighthouse.down:\n\t\t- Stop the isolated Lighthouse Server Docker image with docker-compose down."
	@echo "\tlighthouse.push.image:\n\t\t- Push the Lighthouse Server Docker image to GCR."
	@echo "\tlighthouse.clean.image:\n\t\t- Clean the Lighthouse Server Docker image from the host machine."
	@echo "\tlighthouse.build.cluster:\n\t\t- Build the Lighthouse Server GKE cluster."
	@echo "\tlighthouse.creds:\n\t\t- Get the Lighthouse Server GKE cluster credentials."
	@echo "\tlighthouse.deploy.cluster:\n\t\t- Deploy the Lighthouse Server GKE cluster."
	@echo "\tlighthouse.get.cluster:\n\t\t- Get the Lighthouse Server GKE cluster status."
	@echo "\tlighthouse.clean.cluster:\n\t\t- Clean the Lighthouse Server GKE cluster."

# Build the Lighthouse Server GO binary.
lighthouse.build.bin:
	@echo "Building Lighthouse Server GO binary '${LH_BIN_PATH}' ..."
	@${GOBUILD} ${LDFLAGS} -o ${LH_BIN_PATH} ${LH_PKG_PATH}

# Clean the Lighthouse Server GO binary.
lighthouse.clean.bin:
	@echo "Cleaning up Lighthouse Server GO binary '${LH_BIN_PATH}' ..."
	@if [ -f ${LH_BIN_PATH} ]; then rm ${LH_BIN_PATH} ; fi

# Build the Lighthouse Server Docker image.
lighthouse.build.image: deps lighthouse.clean.bin lighthouse.build.bin
	@echo "Building Docker image [${REPO}/${LH_TAG}] ..."
	@docker build -f docker/lighthouse-server/Dockerfile -t ${REPO}/${LH_TAG} --no-cache .
	@docker tag ${REPO}/${LH_TAG} ${LH_IMAGE}:latest

# Rebuild & run the Lighthouse Server Docker image in isolation with docker-compose up.
lighthouse.build.up: lighthouse.build.image lighthouse.up

# Run the Lighthouse Server Docker image in isolation with docker-compose up.
lighthouse.up:
	@docker-compose up lighthouse-server

# Stop the isolated Lighthouse Server Docker image with docker-compose down.
lighthouse.down:
	@docker-compose down lighthouse-server

# Push the Lighthouse Server Docker image to GCR.
lighthouse.push.image:
	@gcloud docker -- push ${REPO}/${LH_TAG}

# Clean the Lighthouse Server Docker image from the host machine.
lighthouse.clean.image:
	@docker rmi ${REPO}/${LH_TAG}
	@docker rmi ${LH_IMAGE}:latest

# Build the Lighthouse Server GKE cluster.
lighthouse.build.cluster: config
	@gcloud beta container --project ${PROJECT} clusters create ${LH_CLUSTER} \
	--zone ${ZONE} \
	--cluster-version ${LH_CLUSTER_VERSION} \
	--machine-type ${LH_MACHINE_TYPE} \
	--image-type "COS" \
	--disk-size ${LH_DISK_SIZE} \
	--scopes "https://www.googleapis.com/auth/compute","https://www.googleapis.com/auth/devstorage.full_control","https://www.googleapis.com/auth/logging.write","https://www.googleapis.com/auth/monitoring","https://www.googleapis.com/auth/pubsub","https://www.googleapis.com/auth/servicecontrol","https://www.googleapis.com/auth/service.management.readonly","https://www.googleapis.com/auth/trace.append" \
	--preemptible \
	--num-nodes ${LH_NUM_NODES} \
	--enable-cloud-logging \
	--enable-cloud-monitoring \
	--enable-autoscaling \
	--min-nodes ${LH_MIN_NODES} \
	--max-nodes ${LH_MAX_NODES} \
	--enable-legacy-authorization \
	--disable-addons HttpLoadBalancing

# Get the Lighthouse Server GKE cluster credentials.
lighthouse.creds:
	@gcloud container clusters get-credentials ${LH_CLUSTER}

# Deploy the Lighthouse Server GKE cluster.
lighthouse.deploy.cluster: config lighthouse.creds
	@kubectl run ${LH_CLUSTER} --image=${REPO}/${LH_TAG} --replicas=${LH_REPLICAS} \
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
	--env="AWS_S3_BUCKET_NAME=${AWS_S3_BUCKET_NAME}"
	@kubectl autoscale deployment ${LH_CLUSTER} \
	--cpu-percent=${LH_CPU_PERCENT} \
	--min=${LH_MIN_PODS} \
	--max=${LH_MAX_PODS}

# Get the Lighthouse Server GKE cluster status.
lighthouse.get.cluster:
	@kubectl get deployment ${LH_CLUSTER} -o yaml

# Clean the Lighthouse Server GKE cluster.
lighthouse.clean.cluster: config
	@gcloud container clusters delete ${LH_CLUSTER} -q
