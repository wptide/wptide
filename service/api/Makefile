# Show available make commands.
api.usage:
	@echo "\tapi.tpl:\n\t\t- Generate the API templates."

# Run the API Docker images in isolation with docker-compose up.
api.up:
	@docker-compose up api-php

# Run the setup script; API containers must be running to exec into.
api.setup:
	@docker-compose exec api-php /app/setup.sh

# Generate the API templates.
api.tpl:
	@./bin/tpl -t=service/api/tpl/app.tpl -d=service/api/app.yaml -e=.env.gke
	@./bin/tpl -t=service/api/tpl/cron.tpl -d=service/api/cron.yaml -e=.env.gke
	@./bin/tpl -t=service/api/tpl/wp-config.tpl -d=service/api/wp-config.php -e=.env.gke