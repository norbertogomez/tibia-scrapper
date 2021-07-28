help: ## This help dialog.
	@grep '^[a-zA-Z]' $(MAKEFILE_LIST) | \
        sort | \
        awk -F ':.*?## ' 'NF==2 {printf "\033[36m  %-25s\033[0m %s\n", $$1, $$2}'

logs: ## Displays scrapper logs
	@docker-compose logs -f tibia-scrapper

start: ## Starts tibia scrapper
	@docker-compose up --build -d

status: ## Shows container status
	@docker-compose ps

stop: ## Stops tibia scrapper
	@docker-compose down

