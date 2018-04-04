# Lighthouse Server Docker settings.
export LH_TAG=${LH_GKE_IMAGE}:${VERSION}

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
	@echo "\tlighthouse.tpl:\n\t\t- Generate the Lighthouse Server GKE YAML template."
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
	@docker build -f service/lighthouse-server/docker/Dockerfile -t ${REPO}/${LH_TAG} --no-cache .
	@docker tag ${REPO}/${LH_TAG} ${LH_GKE_IMAGE}:latest

# Rebuild & run the Lighthouse Server Docker image in isolation with docker-compose up.
lighthouse.build.up: lighthouse.build.image lighthouse.up

# Run the Lighthouse Server Docker image in isolation with docker-compose up.
lighthouse.up:
	@docker-compose up lighthouse-server

# Stop the isolated Lighthouse Server Docker image with docker-compose down.
lighthouse.down:
	@docker-compose down

# Push the Lighthouse Server Docker image to GCR.
lighthouse.push.image:
	@gcloud docker -- push ${REPO}/${LH_TAG}

# Clean the Lighthouse Server Docker image from the host machine.
lighthouse.clean.image:
	@docker rmi ${REPO}/${LH_TAG}
	@docker rmi ${LH_GKE_IMAGE}:latest

# Build the Lighthouse Server GKE cluster.
lighthouse.build.cluster: config
	@gcloud beta container --project ${GCP_PROJECT} clusters create ${LH_GKE_CLUSTER} \
	--zone ${GCP_ZONE} \
	--cluster-version ${LH_GKE_CLUSTER_VERSION} \
	--machine-type ${LH_GKE_MACHINE_TYPE} \
	--image-type "COS" \
	--disk-size ${LH_GKE_DISK_SIZE} \
	--scopes "https://www.googleapis.com/auth/compute","https://www.googleapis.com/auth/devstorage.full_control","https://www.googleapis.com/auth/logging.write","https://www.googleapis.com/auth/monitoring","https://www.googleapis.com/auth/pubsub","https://www.googleapis.com/auth/servicecontrol","https://www.googleapis.com/auth/service.management.readonly","https://www.googleapis.com/auth/trace.append" \
	--preemptible \
	--num-nodes ${LH_GKE_NUM_NODES} \
	--enable-cloud-logging \
	--enable-cloud-monitoring \
	--enable-autoscaling \
	--min-nodes ${LH_GKE_MIN_NODES} \
	--max-nodes ${LH_GKE_MAX_NODES}

# Get the Lighthouse Server GKE cluster credentials.
lighthouse.creds:
	@gcloud container clusters get-credentials ${LH_GKE_CLUSTER}

# Generate the Lighthouse Server GKE YAML template.
lighthouse.tpl:
	@./bin/tpl -p=service/lighthouse-server/kubernetes/ -f=deployment -e=.env.gke

# Deploy the Lighthouse Server GKE cluster.
lighthouse.deploy.cluster: config lighthouse.creds lighthouse.tpl
	@kubectl apply -f service/lighthouse-server/kubernetes/deployment.yaml

# Get the Lighthouse Server GKE cluster status.
lighthouse.get.cluster: config lighthouse.creds
	@kubectl get deployment ${LH_GKE_CLUSTER} -o yaml

# Clean the Lighthouse Server GKE cluster.
lighthouse.clean.cluster: config lighthouse.creds lighthouse.tpl
	@kubectl delete -f service/lighthouse-server/kubernetes/deployment.yaml
	@gcloud container clusters delete ${LH_GKE_CLUSTER} -q
