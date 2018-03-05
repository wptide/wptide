# Sync Server Docker settings.
SYNC_TAG=${SYNC_GKE_IMAGE}:${VERSION}

# Sync Server GO settings.
SYNC_BIN_PATH=bin/sync-server
SYNC_PKG_PATH=./cmd/sync-server/...

# Show available make commands.
sync.usage:
	@echo "\tsync.build.bin:\n\t\t- Build the Sync Server GO binary."
	@echo "\tsync.clean.bin:\n\t\t- Clean the Sync Server GO binary."
	@echo "\tsync.build.image:\n\t\t- Build the Sync Server Docker image."
	@echo "\tsync.build.up:\n\t\t- Rebuild & run the Sync Server Docker image in isolation with docker-compose up."
	@echo "\tsync.up:\n\t\t- Run the Sync Server Docker image in isolation with docker-compose up."
	@echo "\tsync.down:\n\t\t- Take the isolated Sync Server Docker image down."
	@echo "\tsync.stop:\n\t\t- Stop the isolated Sync Server Docker image with docker-compose stop."
	@echo "\tsync.rm:\n\t\t- Remove the isolated Sync Server Docker image with docker-compose rm."
	@echo "\tsync.push.image:\n\t\t- Push the Sync Server Docker image to GCR."
	@echo "\tsync.clean.image:\n\t\t- Clean the Sync Server Docker image from the host machine."
	@echo "\tsync.build.cluster:\n\t\t- Build the Sync Server GKE cluster."
	@echo "\tsync.creds:\n\t\t- Get the Sync Server GKE cluster credentials."
	@echo "\tsync.deploy.cluster:\n\t\t- Deploy the Sync Server GKE cluster."
	@echo "\tsync.get.cluster:\n\t\t- Get the Sync Server GKE cluster status."
	@echo "\tsync.clean.cluster:\n\t\t- Clean the Sync Server GKE cluster."
	@echo "\tsync.update.certs:\n\t\t- Get CA Root Certs from curl mirror."

# Build the Sync Server GO binary.
sync.build.bin:
	@echo "Building Sync Server GO binary '${SYNC_BIN_PATH}' ..."
	@${GOBUILD} ${LDFLAGS} -o ${SYNC_BIN_PATH} ${SYNC_PKG_PATH}

# Clean the Sync Server GO binary.
sync.clean.bin:
	@echo "Cleaning up Sync Server GO binary '${SYNC_BIN_PATH}' ..."
	@if [ -f ${SYNC_BIN_PATH} ]; then rm ${SYNC_BIN_PATH} ; fi

# Build the Sync Server Docker image.
sync.build.image: deps sync.clean.bin sync.build.bin
	@echo "Building Docker image [${REPO}/${SYNC_TAG}] ..."
	@docker build -f docker/sync-server/Dockerfile -t ${REPO}/${SYNC_TAG} --no-cache .
	@docker tag ${REPO}/${SYNC_TAG} ${SYNC_GKE_IMAGE}:latest

# Rebuild & run the Sync Server Docker image in isolation with docker-compose up.
sync.build.up: sync.build.image sync.up

# Run the Sync Server Docker image in isolation with docker-compose up.
sync.up:
	@docker-compose up sync-server

# Take the isolated Sync Server Docker image down.
sync.down: sync.stop sync.rm

# Stop the isolated Sync Server Docker image with docker-compose stop.
sync.stop:
	@docker-compose stop sync-server

# Remove the isolated Sync Server Docker image with docker-compose rm.
sync.rm:
	@docker-compose rm -f sync-server

# Push the Sync Server Docker image to GCR.
sync.push.image:
	@gcloud docker -- push ${REPO}/${SYNC_TAG}

# Clean the Sync Server Docker image from the host machine.
sync.clean.image:
	@docker rmi ${REPO}/${SYNC_TAG}
	@docker rmi ${SYNC_GKE_IMAGE}:latest

# Build the Sync Server GKE cluster.
sync.build.cluster: config
	@gcloud beta container --project ${GCP_PROJECT} clusters create ${SYNC_GKE_CLUSTER} \
	--zone ${GCP_ZONE} \
	--cluster-version ${SYNC_GKE_CLUSTER_VERSION} \
	--machine-type ${SYNC_GKE_MACHINE_TYPE} \
	--image-type "COS" \
	--disk-size ${SYNC_GKE_DISK_SIZE} \
	--scopes "https://www.googleapis.com/auth/compute","https://www.googleapis.com/auth/devstorage.full_control","https://www.googleapis.com/auth/logging.write","https://www.googleapis.com/auth/monitoring","https://www.googleapis.com/auth/pubsub","https://www.googleapis.com/auth/servicecontrol","https://www.googleapis.com/auth/service.management.readonly","https://www.googleapis.com/auth/trace.append" \
	--preemptible \
	--num-nodes 1 \
	--enable-cloud-logging \
	--enable-cloud-monitoring \
	--enable-legacy-authorization \
	--disable-addons HttpLoadBalancing

# Get the Sync Server GKE cluster credentials.
sync.creds:
	@gcloud container clusters get-credentials ${SYNC_GKE_CLUSTER}

# Deploy the Sync Server GKE cluster.
sync.deploy.cluster: config sync.creds
	@kubectl run ${SYNC_GKE_CLUSTER} --image=${REPO}/${SYNC_TAG} \
	--env="TIDE_API_HOST=${TIDE_API_HOST}" \
	--env="TIDE_API_PROTOCOL=${TIDE_API_PROTOCOL}" \
	--env="TIDE_API_VERSION=${TIDE_API_VERSION}" \
	--env="SYNC_ACTIVE=${SYNC_ACTIVE}" \
    --env="SYNC_API_BROWSE_CATEGORY=${SYNC_API_BROWSE_CATEGORY}" \
    --env="SYNC_DATA=${SYNC_DATA}" \
    --env="SYNC_DEFAULT_CLIENT=${SYNC_DEFAULT_CLIENT}" \
    --env="SYNC_DEFAULT_VISIBILITY=${SYNC_DEFAULT_VISIBILITY}" \
    --env="SYNC_FORCE_AUDITS=${SYNC_FORCE_AUDITS}" \
    --env="SYNC_ITEMS_PER_PAGE=${SYNC_ITEMS_PER_PAGE}" \
    --env="SYNC_POOL_DELAY=${SYNC_POOL_DELAY}" \
    --env="SYNC_POOL_WORKERS=${SYNC_POOL_WORKERS}" \
	--env="SYNC_SQS_KEY=${SYNC_SQS_KEY}" \
	--env="SYNC_SQS_QUEUE=${SYNC_SQS_QUEUE}" \
	--env="SYNC_SQS_QUEUE_NAME=${SYNC_SQS_QUEUE_NAME}" \
	--env="SYNC_SQS_REGION=${SYNC_SQS_REGION}" \
	--env="SYNC_SQS_SECRET=${SYNC_SQS_SECRET}" \
    --env="SYNC_SQS_THEME_KEY=${SYNC_SQS_THEME_KEY}" \
    --env="SYNC_SQS_THEME_QUEUE=${SYNC_SQS_THEME_QUEUE}" \
    --env="SYNC_SQS_THEME_QUEUE_NAME=${SYNC_SQS_THEME_QUEUE_NAME}" \
    --env="SYNC_SQS_THEME_REGION=${SYNC_SQS_THEME_REGION}" \
    --env="SYNC_SQS_THEME_SECRET=${SYNC_SQS_THEME_SECRET}" \
    --env="SYNC_SQS_PLUGIN_KEY=${SYNC_SQS_PLUGIN_KEY}" \
    --env="SYNC_SQS_PLUGIN_QUEUE=${SYNC_SQS_PLUGIN_QUEUE}" \
    --env="SYNC_SQS_PLUGIN_QUEUE_NAME=${SYNC_SQS_PLUGIN_QUEUE_NAME}" \
    --env="SYNC_SQS_PLUGIN_REGION=${SYNC_SQS_PLUGIN_REGION}" \
    --env="SYNC_SQS_PLUGIN_SECRET=${SYNC_SQS_PLUGIN_SECRET}"

# Get the Sync Server GKE cluster status.
sync.get.cluster:
	@kubectl get deployment ${SYNC_GKE_CLUSTER} -o yaml

# Clean the Sync Server GKE cluster.
sync.clean.cluster: config
	@gcloud container clusters delete ${SYNC_GKE_CLUSTER} -q

# Get CA Certs for scratch file.
sync.update.certs:
	@curl --remote-name --time-cond cacert.pem https://curl.haxx.se/ca/cacert.pem