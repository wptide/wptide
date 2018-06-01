# Show available make commands.
api.usage:
	@echo "\tapi.up:\n\t\t- Run the API Docker images in isolation with docker-compose up."
	@echo "\tapi.down:\n\t\t- Take the isolated API Docker images down."
	@echo "\tapi.stop:\n\t\t- Stop the isolated API Docker images with docker-compose stop."
	@echo "\tapi.rm:\n\t\t- Remove the isolated API Docker images with docker-compose rm."
	@echo "\tapi.composer:\n\t\t- Install the Composer dependencies."
	@echo "\tapi.setup:\n\t\t- Run the setup script; API containers must be running to exec into."
	@echo "\tapi.tpl:\n\t\t- Generate the API templates."
	@echo "\tapi.deploy.sql:\n\t\t- Deploy the Cloud SQL database & setup the users."
	@echo "\tapi.deploy.app:\n\t\t- Deploy the API to Google App Engine."
	@echo "\tapi.build.disk:\n\t\t- Build the Redis GKE persistent disk."
	@echo "\tapi.build.cluster:\n\t\t- Build the Redis GKE cluster."
	@echo "\tapi.deploy.cluster:\n\t\t- Deploy the Redis GKE cluster."
	@echo "\tapi.get.cluster:\n\t\t- Get the Redis GKE cluster status."
	@echo "\tapi.clean.cluster:\n\t\t- Clean the Redis GKE cluster."
	@echo "\tapi.creds:\n\t\t- Get the Redis GKE cluster credentials."

# Run the API Docker images in isolation with docker-compose up.
api.up:
	@docker-compose up api-php

# Take the isolated API Docker images down.
api.down: api.stop api.rm

# Stop the isolated API Docker images with docker-compose stop.
api.stop:
	@docker-compose stop api-php

# Remove the isolated API Docker images with docker-compose rm.
api.rm:
	@docker-compose rm -f api-php

# Install the Composer dependencies.
api.composer: api.tpl
	@cd service/api && composer install --no-interaction && cd --
	@cd service/api/wp-content/plugins/wp-tide-api && composer install --no-interaction && cd --

# Run the setup script; API containers must be running to exec into.
api.setup:
	@docker-compose exec api-php /app/setup.sh

# Generate the API templates.
api.tpl:
	@./bin/tpl -t=service/api/tpl/app.tpl -d=service/api/app.yaml -e=.env.gcp
	@./bin/tpl -t=service/api/tpl/cron.tpl -d=service/api/cron.yaml -e=.env.gcp
	@./bin/tpl -t=service/api/tpl/redis.tpl -d=service/api/redis.yaml -e=.env.gcp
	@./bin/tpl -t=service/api/tpl/wp-config.tpl -d=service/api/wp-config.php -e=.env.gcp

# Deploy the Cloud SQL database & setup the users.
api.deploy.sql:
	@gcloud sql instances create ${GCSQL_API_INSTANCE} \
	--region ${GCP_REGION} \
	--tier ${GCSQL_API_TIER} \
	--storage-size ${GCSQL_API_STORAGE_SIZE} \
	--database-version ${GCSQL_API_DATABASE_VERSION} \
	--backup-start-time ${GCSQL_API_BACKUP_START_TIME} \
	--enable-bin-log \
	--failover-replica-name ${GCSQL_API_FAILOVER_REPLICA_NAME} \
	--maintenance-release-channel ${GCSQL_API_MAINTENANCE_RELEASE_CHANNEL} \
	--maintenance-window-day ${GCSQL_API_MAINTENANCE_WINDOW_DAY} \
	--maintenance-window-hour ${GCSQL_API_MAINTENANCE_WINDOW_HOUR}
	@gcloud sql users set-password root % --instance ${GCSQL_API_INSTANCE} --password "${GCSQL_API_DB_ROOT_PASSWORD}"
	@gcloud sql databases create ${GCSQL_API_DB_NAME} --instance ${GCSQL_API_INSTANCE} --charset utf8 --collation utf8_general_ci
	@gcloud sql users create ${GCSQL_API_DB_USER} % --instance ${GCSQL_API_INSTANCE} --password "${GCSQL_API_DB_PASSWORD}"

# Deploy the API to Google App Engine.
api.deploy.app: api.tpl
	@cd service/api && gcloud app deploy --promote --stop-previous-version app.yaml cron.yaml && cd --

# Build the Redis GKE persistent disk.
api.build.disk: config
	@gcloud compute --project ${GCP_PROJECT} disks create ${GKE_REDIS_CLUSTER}-disk \
	--zone ${GCP_ZONE} \
	--type pd-ssd \
	--size 100GB

# Build the Redis GKE cluster.
api.build.cluster: config
	@gcloud container --project ${GCP_PROJECT} clusters create ${GKE_REDIS_CLUSTER} \
	--zone ${GCP_ZONE} \
	--cluster-version ${GKE_REDIS_CLUSTER_VERSION} \
	--machine-type ${GKE_REDIS_MACHINE_TYPE} \
	--image-type "COS" \
	--disk-size ${GKE_REDIS_DISK_SIZE} \
	--scopes "https://www.googleapis.com/auth/compute","https://www.googleapis.com/auth/devstorage.full_control","https://www.googleapis.com/auth/logging.write","https://www.googleapis.com/auth/monitoring","https://www.googleapis.com/auth/pubsub","https://www.googleapis.com/auth/servicecontrol","https://www.googleapis.com/auth/service.management.readonly","https://www.googleapis.com/auth/trace.append" \
	--num-nodes 1 \
	--enable-cloud-logging \
	--enable-cloud-monitoring

# Deploy the Redis GKE cluster.
api.deploy.cluster: config api.creds api.tpl
	@kubectl apply -f service/api/redis.yaml

# Get the Redis GKE cluster status.
api.get.cluster: config api.creds
	@kubectl get deployment ${GKE_REDIS_CLUSTER} -o yaml

# Clean the Redis GKE cluster.
api.clean.cluster: config api.creds api.tpl
	@kubectl delete -f service/api/redis.yaml
	@gcloud compute --project ${GCP_PROJECT} disks delete ${GKE_REDIS_CLUSTER}-disk
	@gcloud container clusters delete ${GKE_REDIS_CLUSTER} -q

# Get the Redis GKE cluster credentials.
api.creds:
	@gcloud container clusters get-credentials ${GKE_REDIS_CLUSTER}