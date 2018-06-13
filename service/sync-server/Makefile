# Sync Server Docker settings.
export GCR_SYNC_IMAGE_TAG=gcr.io/${GCP_PROJECT}/${GKE_SYNC_IMAGE}:${VERSION}

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
	@echo "\tsync.build.disk:\n\t\t- Build the Sync Server GKE persistent disk."
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
	@echo "Building Docker image [${GCR_SYNC_IMAGE_TAG}] ..."
	@docker build -f service/sync-server/docker/Dockerfile -t ${GCR_SYNC_IMAGE_TAG} --no-cache .
	@docker tag ${GCR_SYNC_IMAGE_TAG} ${GKE_SYNC_IMAGE}:latest

# Rebuild & run the Sync Server Docker image in isolation with docker-compose up.
sync.build.up: sync.build.image sync.up

# Run the Sync Server Docker image in isolation with docker-compose up.
sync.up:
	@docker-compose up sync-server

# Start the Sync Server Docker daemon in isolation with docker-compose up.
sync.upd:
	@docker-compose up -d sync-server

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
	@gcloud docker -- push ${GCR_SYNC_IMAGE_TAG}

# Clean the Sync Server Docker image from the host machine.
sync.clean.image:
	@docker rmi ${GCR_SYNC_IMAGE_TAG}
	@docker rmi ${GKE_SYNC_IMAGE}:latest

# Build the Sync Server GKE persistent disk.
sync.build.disk: config
	@gcloud compute --project ${GCP_PROJECT} disks create ${GKE_SYNC_CLUSTER}-disk \
	--zone ${GCP_ZONE} \
	--type ${GKE_SYNC_PERSISTENT_DISK_TYPE} \
	--size ${GKE_SYNC_PERSISTENT_DISK_SIZE}

# Build the Sync Server GKE cluster.
sync.build.cluster: config
	@gcloud container --project ${GCP_PROJECT} clusters create ${GKE_SYNC_CLUSTER} \
	--zone ${GCP_ZONE} \
	--cluster-version ${GKE_SYNC_CLUSTER_VERSION} \
	--machine-type ${GKE_SYNC_MACHINE_TYPE} \
	--image-type "COS" \
	--disk-size ${GKE_SYNC_DISK_SIZE} \
	--scopes "https://www.googleapis.com/auth/compute","https://www.googleapis.com/auth/cloud-platform","https://www.googleapis.com/auth/devstorage.full_control","https://www.googleapis.com/auth/logging.write","https://www.googleapis.com/auth/monitoring","https://www.googleapis.com/auth/pubsub","https://www.googleapis.com/auth/servicecontrol","https://www.googleapis.com/auth/service.management.readonly","https://www.googleapis.com/auth/trace.append" \
	--preemptible \
	--num-nodes 1 \
	--enable-cloud-logging \
	--enable-cloud-monitoring

# Get the Sync Server GKE cluster credentials.
sync.creds:
	@gcloud container clusters get-credentials ${GKE_SYNC_CLUSTER}

# Generate the Sync Server GKE YAML template.
sync.tpl:
	@./bin/tpl \
	-t=service/sync-server/kubernetes/deployment.tpl \
	-d=service/sync-server/kubernetes/deployment.yaml \
	-e=.env.gcp

# Deploy the Sync Server GKE cluster.
sync.deploy.cluster: config sync.creds sync.tpl
	@kubectl apply -f service/sync-server/kubernetes/deployment.yaml

# Get the Sync Server GKE cluster status.
sync.get.cluster: config sync.creds
	@kubectl get deployment ${GKE_SYNC_CLUSTER} -o yaml

# Clean the Sync Server GKE cluster.
sync.clean.cluster: config sync.creds sync.tpl
	@kubectl delete -f service/sync-server/kubernetes/deployment.yaml
	@gcloud container clusters delete ${GKE_SYNC_CLUSTER} -q
	@gcloud compute --project ${GCP_PROJECT} disks delete ${GKE_SYNC_CLUSTER}-disk
