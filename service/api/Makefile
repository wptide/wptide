# Show available make commands.
api.usage:
	@echo "\tapi.up:\n\t\t- Run the API Docker images in isolation with docker-compose up."
	@echo "\tapi.composer:\n\t\t- Install the Composer dependencies."
	@echo "\tapi.setup:\n\t\t- Run the setup script; API containers must be running to exec into."
	@echo "\tapi.tpl:\n\t\t- Generate the API templates."

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
	@./bin/tpl -t=service/api/tpl/app.tpl -d=service/api/app.yaml -e=.env.gke
	@./bin/tpl -t=service/api/tpl/cron.tpl -d=service/api/cron.yaml -e=.env.gke
	@./bin/tpl -t=service/api/tpl/wp-config.tpl -d=service/api/wp-config.php -e=.env.gke