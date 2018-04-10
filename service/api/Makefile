# Show available make commands.
api.usage:
	@echo "\tapi.tpl:\n\t\t- Generate the API templates."

# Generate the API templates.
api.tpl:
	@./bin/tpl -t=service/api/tpl/app.tpl -d=service/api/app.yaml -e=.env.gke
	@./bin/tpl -t=service/api/tpl/cron.tpl -d=service/api/cron.yaml -e=.env.gke
	@./bin/tpl -t=service/api/tpl/wp-config.tpl -d=service/api/wp-config.php -e=.env.gke