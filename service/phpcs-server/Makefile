# PHPCS Server Docker settings.
export GCR_PHPCS_IMAGE_TAG=gcr.io/${GCP_PROJECT}/${GKE_PHPCS_IMAGE}:${VERSION}

# PHPCS Server GO settings.
PHPCS_BIN_PATH=bin/phpcs-server
PHPCS_PKG_PATH=./cmd/phpcs-server/...

# Show available make commands.
phpcs.usage:
	@echo "\tphpcs.build.bin:\n\t\t- Build the PHPCS Server Go binary."
	@echo "\tphpcs.clean.bin:\n\t\t- Clean the PHPCS Server Go binary."
	@echo "\tphpcs.build.image:\n\t\t- Build the PHPCS Server Docker image."
	@echo "\tphpcs.build.up:\n\t\t- Rebuild the PHPCS Server Docker image and run the container in isolation with docker-compose up."
	@echo "\tphpcs.up:\n\t\t- Run the PHPCS Server Docker container in isolation with docker-compose up."
	@echo "\tphpcs.down:\n\t\t- Take the isolated PHPCS Server Docker container down."
	@echo "\tphpcs.stop:\n\t\t- Stop the isolated PHPCS Server Docker container with docker-compose stop."
	@echo "\tphpcs.rm:\n\t\t- Remove the stopped PHPCS Server Docker container with docker-compose rm."
	@echo "\tphpcs.push.image:\n\t\t- Push the PHPCS Server Docker image to GCR."
	@echo "\tphpcs.clean.image:\n\t\t- Clean the PHPCS Server Docker image from the host machine."
	@echo "\tphpcs.build.cluster:\n\t\t- Build the PHPCS Server GKE cluster."
	@echo "\tphpcs.creds:\n\t\t- Get the PHPCS Server GKE cluster credentials."
	@echo "\tphpcs.tpl:\n\t\t- Generate the PHPCS Server GKE YAML template."
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
	@echo "Building Docker image [${GCR_PHPCS_IMAGE_TAG}] ..."
	@docker build -f service/phpcs-server/docker/Dockerfile -t ${GCR_PHPCS_IMAGE_TAG} --no-cache .
	@docker tag ${GCR_PHPCS_IMAGE_TAG} ${GKE_PHPCS_IMAGE}:latest

# Rebuild & run the PHPCS Server Docker image in isolation with docker-compose up.
phpcs.build.up: phpcs.build.image phpcs.up

# Run the PHPCS Server Docker image in isolation with docker-compose up.
phpcs.up:
	@docker-compose up phpcs-server

# Take the isolated PHPCS Server Docker image down.
phpcs.down: phpcs.stop phpcs.rm

# Stop the isolated PHPCS Server Docker image with docker-compose stop.
phpcs.stop:
	@docker-compose stop phpcs-server

# Remove the isolated PHPCS Server Docker image with docker-compose rm.
phpcs.rm:
	@docker-compose rm -f phpcs-server

# Push the PHPCS Server Docker image to GCR.
phpcs.push.image:
	@gcloud docker -- push ${GCR_PHPCS_IMAGE_TAG}

# Clean the PHPCS Server Docker image from the host machine.
phpcs.clean.image:
	@docker rmi ${GCR_PHPCS_IMAGE_TAG}
	@docker rmi ${GKE_PHPCS_IMAGE}:latest

# Build the PHPCS Server GKE cluster.
phpcs.build.cluster: config
	@gcloud container --project ${GCP_PROJECT} clusters create ${GKE_PHPCS_CLUSTER} \
	--zone ${GCP_ZONE} \
	--cluster-version ${GKE_PHPCS_CLUSTER_VERSION} \
	--machine-type ${GKE_PHPCS_MACHINE_TYPE} \
	--image-type "COS" \
	--disk-size ${GKE_PHPCS_DISK_SIZE} \
	--scopes "https://www.googleapis.com/auth/cloud-platform","https://www.googleapis.com/auth/compute","https://www.googleapis.com/auth/devstorage.full_control","https://www.googleapis.com/auth/logging.write","https://www.googleapis.com/auth/monitoring","https://www.googleapis.com/auth/servicecontrol","https://www.googleapis.com/auth/service.management.readonly","https://www.googleapis.com/auth/trace.append" \
	--preemptible \
	--num-nodes ${GKE_PHPCS_NUM_NODES} \
	--enable-cloud-logging \
	--enable-cloud-monitoring \
	--enable-autoscaling \
	--min-nodes ${GKE_PHPCS_MIN_NODES} \
	--max-nodes ${GKE_PHPCS_MAX_NODES} \
	--addons HorizontalPodAutoscaling,KubernetesDashboard \
	--no-enable-autorepair

# Get the PHPCS Server GKE cluster credentials.
phpcs.creds:
	@gcloud container clusters get-credentials ${GKE_PHPCS_CLUSTER}

# Generate the PHPCS Server GKE YAML template.
phpcs.tpl:
	@docker-compose run -e GCR_PHPCS_IMAGE_TAG=${GCR_PHPCS_IMAGE_TAG} --rm tpl \
		-t=/srv/service/phpcs-server/kubernetes/deployment.tpl \
		-d=/srv/data/phpcs-server/deployment.yaml \
	&& cp \
		data/tpl/phpcs-server/deployment.yaml \
		service/phpcs-server/kubernetes/deployment.yaml

# Deploy the PHPCS Server GKE cluster.
phpcs.deploy.cluster: config phpcs.creds phpcs.tpl
	@kubectl apply -f service/phpcs-server/kubernetes/deployment.yaml

# Get the PHPCS Server GKE cluster status.
phpcs.get.cluster: config phpcs.creds
	@kubectl get deployment ${GKE_PHPCS_CLUSTER} -o yaml

# Clean the PHPCS Server GKE cluster.
phpcs.clean.cluster: config phpcs.creds phpcs.tpl
	@kubectl delete -f service/phpcs-server/kubernetes/deployment.yaml
	@gcloud container clusters delete ${GKE_PHPCS_CLUSTER} -q
