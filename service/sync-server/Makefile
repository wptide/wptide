# Sync Server Docker settings.
export SYNC_TAG=${SYNC_GKE_IMAGE}:${VERSION}

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
	@echo "\tsync.tpl:\n\t\t- Generate the Sync Server GKE YAML template."
	@echo "\tsync.deploy.cluster:\n\t\t- Deploy the Sync Server GKE cluster."
	@echo "\tsync.get.cluster:\n\t\t- Get the Sync Server GKE cluster status."
	@echo "\tsync.clean.cluster:\n\t\t- Clean the Sync Server GKE cluster."

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
	@docker build -f service/sync-server/docker/Dockerfile -t ${REPO}/${SYNC_TAG} --no-cache .
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
	@gcloud container --project ${GCP_PROJECT} clusters create ${SYNC_GKE_CLUSTER} \
	--zone ${GCP_ZONE} \
	--cluster-version ${SYNC_GKE_CLUSTER_VERSION} \
	--machine-type ${SYNC_GKE_MACHINE_TYPE} \
	--image-type "COS" \
	--disk-size ${SYNC_GKE_DISK_SIZE} \
	--scopes "https://www.googleapis.com/auth/compute","https://www.googleapis.com/auth/devstorage.full_control","https://www.googleapis.com/auth/logging.write","https://www.googleapis.com/auth/monitoring","https://www.googleapis.com/auth/pubsub","https://www.googleapis.com/auth/servicecontrol","https://www.googleapis.com/auth/service.management.readonly","https://www.googleapis.com/auth/trace.append" \
	--preemptible \
	--num-nodes 1 \
	--enable-cloud-logging \
	--enable-cloud-monitoring

# Get the Sync Server GKE cluster credentials.
sync.creds:
	@gcloud container clusters get-credentials ${SYNC_GKE_CLUSTER}

# Generate the Sync Server GKE YAML template.
sync.tpl:
	@./bin/tpl -p=service/sync-server/kubernetes/ -f=deployment -e=.env.gke

# Deploy the Sync Server GKE cluster.
sync.deploy.cluster: config sync.creds sync.tpl
	@kubectl apply -f service/sync-server/kubernetes/deployment.yaml

# Get the Sync Server GKE cluster status.
sync.get.cluster: config sync.creds
	@kubectl get deployment ${SYNC_GKE_CLUSTER} -o yaml

# Clean the Sync Server GKE cluster.
sync.clean.cluster: config sync.creds sync.tpl
	@kubectl delete -f service/sync-server/kubernetes/deployment.yaml
	@gcloud container clusters delete ${SYNC_GKE_CLUSTER} -q
