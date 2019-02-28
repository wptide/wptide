# Lighthouse Server Docker settings.
export GCR_LH_IMAGE_TAG=gcr.io/${GCP_PROJECT}/${GKE_LH_IMAGE}:${VERSION}

# Lighthouse Server GO settings.
LH_BIN_PATH=bin/lighthouse-server
LH_PKG_PATH=./cmd/lighthouse-server/...

# Show available make commands.
lighthouse.usage:
	@echo "\tlighthouse.build.bin:\n\t\t- Build the Lighthouse Server Go binary."
	@echo "\tlighthouse.clean.bin:\n\t\t- Clean the Lighthouse Server Go binary."
	@echo "\tlighthouse.build.image:\n\t\t- Build the Lighthouse Server Docker image."
	@echo "\tlighthouse.build.up:\n\t\t- Rebuild the Lighthouse Server Docker image and run the container in isolation with docker-compose up."
	@echo "\tlighthouse.up:\n\t\t- Run the Lighthouse Server Docker image in isolation with docker-compose up."
	@echo "\tlighthouse.down:\n\t\t- Take the isolated Lighthouse Server Docker container down."
	@echo "\tlighthouse.stop:\n\t\t- Stop the isolated Lighthouse Server Docker container with docker-compose stop."
	@echo "\tlighthouse.rm:\n\t\t- Remove the stopped Lighthouse Server Docker container with docker-compose rm."
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
	@echo "Building Docker image [${GCR_LH_IMAGE_TAG}] ..."
	@docker build -f service/lighthouse-server/docker/Dockerfile -t ${GCR_LH_IMAGE_TAG} --no-cache .
	@docker tag ${GCR_LH_IMAGE_TAG} ${GKE_LH_IMAGE}:latest

# Rebuild & run the Lighthouse Server Docker image in isolation with docker-compose up.
lighthouse.build.up: lighthouse.build.image lighthouse.up

# Run the Lighthouse Server Docker image in isolation with docker-compose up.
lighthouse.up:
	@docker-compose up lighthouse-server

# Take the isolated Lighthouse Server Docker image down.
lighthouse.down: lighthouse.stop lighthouse.rm

# Stop the isolated Lighthouse Server Docker image with docker-compose stop.
lighthouse.stop:
	@docker-compose stop lighthouse-server

# Remove the isolated Lighthouse Server Docker image with docker-compose rm.
lighthouse.rm:
	@docker-compose rm -f lighthouse-server

# Push the Lighthouse Server Docker image to GCR.
lighthouse.push.image:
	@gcloud docker -- push ${GCR_LH_IMAGE_TAG}

# Clean the Lighthouse Server Docker image from the host machine.
lighthouse.clean.image:
	@docker rmi ${GCR_LH_IMAGE_TAG}
	@docker rmi ${GKE_LH_IMAGE}:latest

# Build the Lighthouse Server GKE cluster.
lighthouse.build.cluster: config
	@gcloud container --project ${GCP_PROJECT} clusters create ${GKE_LH_CLUSTER} \
	--zone ${GCP_ZONE} \
	--cluster-version ${GKE_LH_CLUSTER_VERSION} \
	--machine-type ${GKE_LH_MACHINE_TYPE} \
	--image-type "COS" \
	--disk-size ${GKE_LH_DISK_SIZE} \
	--scopes "https://www.googleapis.com/auth/cloud-platform","https://www.googleapis.com/auth/compute","https://www.googleapis.com/auth/devstorage.full_control","https://www.googleapis.com/auth/logging.write","https://www.googleapis.com/auth/monitoring","https://www.googleapis.com/auth/servicecontrol","https://www.googleapis.com/auth/service.management.readonly","https://www.googleapis.com/auth/trace.append" \
	--preemptible \
	--num-nodes ${GKE_LH_NUM_NODES} \
	--enable-cloud-logging \
	--enable-cloud-monitoring \
	--enable-autoscaling \
	--min-nodes ${GKE_LH_MIN_NODES} \
	--max-nodes ${GKE_LH_MAX_NODES} \
	--addons HorizontalPodAutoscaling,KubernetesDashboard \
	--no-enable-autorepair

# Get the Lighthouse Server GKE cluster credentials.
lighthouse.creds:
	@gcloud container clusters get-credentials ${GKE_LH_CLUSTER}

# Generate the Lighthouse Server GKE YAML template.
lighthouse.tpl:
	@docker-compose run -e GCR_LH_IMAGE_TAG=${GCR_LH_IMAGE_TAG} --rm tpl \
		-t=/srv/service/lighthouse-server/kubernetes/deployment.tpl \
		-d=/srv/data/lighthouse-server/deployment.yaml \
	&& cp \
		data/tpl/lighthouse-server/deployment.yaml \
		service/lighthouse-server/kubernetes/deployment.yaml

# Deploy the Lighthouse Server GKE cluster.
lighthouse.deploy.cluster: config lighthouse.creds lighthouse.tpl
	@kubectl apply -f service/lighthouse-server/kubernetes/deployment.yaml

# Get the Lighthouse Server GKE cluster status.
lighthouse.get.cluster: config lighthouse.creds
	@kubectl get deployment ${GKE_LH_CLUSTER} -o yaml

# Clean the Lighthouse Server GKE cluster.
lighthouse.clean.cluster: config lighthouse.creds lighthouse.tpl
	@kubectl delete -f service/lighthouse-server/kubernetes/deployment.yaml
	@gcloud container clusters delete ${GKE_LH_CLUSTER} -q
