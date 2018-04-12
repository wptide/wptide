# Show available make commands.
api.usage:
	@echo "\tapi.up:\n\t\t- Run the API Docker images in isolation with docker-compose up."
	@echo "\tapi.composer:\n\t\t- Install the Composer dependencies."
	@echo "\tapi.setup:\n\t\t- Run the setup script; API containers must be running to exec into."
	@echo "\tapi.tpl:\n\t\t- Generate the API templates."
	@echo "\tapi.deploy.sql:\n\t\t- Deploy the Cloud SQL database & setup the users."
	@echo "\tapi.deploy.app:\n\t\t- Deploy the API to Google App Engine."

# Run the API Docker images in isolation with docker-compose up.
api.up:
	@docker-compose up api-php

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
	@./bin/tpl -t=service/api/tpl/wp-config.tpl -d=service/api/wp-config.php -e=.env.gcp

# Deploy the Cloud SQL database & setup the users.
api.deploy.sql:
	@gcloud sql instances create ${API_GAE_SQL_INSTANCE_NAME} \
	--region ${GCP_REGION} \
	--tier ${API_GAE_SQL_MACHINE_TYPE} \
	--storage-size ${API_GAE_SQL_STORAGE_SIZE} \
	--database-version ${API_GAE_SQL_DB_VERSION} \
	--backup-start-time ${API_GAE_SQL_BACKUP_START_TIME} \
	--enable-bin-log \
	--failover-replica-name ${API_GAE_SQL_INSTANCE_FAILOVER_NAME} \
	--maintenance-release-channel ${API_GAE_SQL_MAINTENANCE_RELEASE_CHANNEL} \
	--maintenance-window-day ${API_GAE_SQL_MAINTENANCE_WINDOW_DAY} \
	--maintenance-window-hour ${API_GAE_SQL_MAINTENANCE_WINDOW_HOUR}
	@gcloud sql users set-password root % --instance ${API_GAE_SQL_INSTANCE_NAME} --password "${API_GAE_SQL_ROOT_PASSWORD}"
	@gcloud sql databases create ${API_GAE_SQL_DB} --instance ${API_GAE_SQL_INSTANCE_NAME} --charset utf8 --collation utf8_general_ci
	@gcloud sql users create ${API_GAE_SQL_USER} % --instance ${API_GAE_SQL_INSTANCE_NAME} --password "${API_GAE_SQL_PASSWORD}"

# Deploy the API to Google App Engine.
api.deploy.app:
	@cd service/api && gcloud app deploy --promote --stop-previous-version app.yaml cron.yaml && cd --