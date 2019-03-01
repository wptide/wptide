# Show available make commands.
api.usage:
	@echo "\tapi.up:\n\t\t- Run the API Docker containers in isolation with docker-compose up."
	@echo "\tapi.down:\n\t\t- Take the isolated API Docker containers down."
	@echo "\tapi.stop:\n\t\t- Stop the isolated API Docker containers with docker-compose stop."
	@echo "\tapi.rm:\n\t\t- Remove the stopped API Docker containers with docker-compose rm."
	@echo "\tapi.composer:\n\t\t- Install the Composer dependencies. Runs make api.tpl automatically before installing the dependencies."
	@echo "\tapi.setup:\n\t\t- Run the setup script; API containers must be running to exec into."
	@echo "\tapi.tpl:\n\t\t- Generate the API YAML & PHP templates."
	@echo "\tapi.deploy.sql:\n\t\t- Deploy the Cloud SQL database & setup the users."
	@echo "\tapi.deploy.app:\n\t\t- Deploy the API to Google App Engine."
	@echo "\tapi.deploy.redis:\n\t\t- Deploy the Google Cloud Memorystore Redis instance."
	@echo "\tapi.get.redis:\n\t\t- Get metadata, including the internal VPC IP address, for the Google Cloud Memorystore Redis instance"
	@echo "\tapi.clean.redis:\n\t\t- Delete the Google Cloud Memorystore Redis instance."

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
	@docker-compose run --rm tpl -t=/srv/service/api/tpl/app.tpl -d=/srv/data/api/app.yaml \
		&& cp data/tpl/api/app.yaml service/api/app.yaml
	@docker-compose run --rm tpl -t=/srv/service/api/tpl/cron.tpl -d=/srv/data/api/cron.yaml \
		&& cp data/tpl/api/cron.yaml service/api/cron.yaml
	@docker-compose run --rm tpl -t=/srv/service/api/tpl/wp-config.tpl -d=/srv/data/api/wp-config.php \
		&& cp data/tpl/api/wp-config.php service/api/wp-config.php

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
	@cp service-account.json service/api/service-account.json \
	@cd service/api && gcloud app deploy --promote --stop-previous-version app.yaml cron.yaml && cd -- \
	@rm service/api/service-account.json

# Deploy the Google Cloud Memorystore Redis instance.
api.deploy.redis: config
	@gcloud beta redis instances create ${GCM_INSTANCE_NAME} \
	--region=${GCP_REGION} \
	--zone=${GCP_ZONE} \
	--size=${GCM_INSTANCE_SIZE} \
	--tier=${GCM_INSTANCE_TIER} \
	--redis-config=maxmemory-policy=allkeys-lru

# Get metadata, including the internal VPC IP address, for the Google Cloud Memorystore Redis instance.
api.get.redis: config
	@gcloud beta redis instances describe ${GCM_INSTANCE_NAME} --region=${GCP_REGION}

# Delete the Google Cloud Memorystore Redis instance.
api.clean.redis: config
	@gcloud beta redis instances delete ${GCM_INSTANCE_NAME} --region=${GCP_REGION}