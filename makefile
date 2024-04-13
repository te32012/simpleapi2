DOCKER_COMPOSE:=./docker-compose.yml
DOCKER_ENV := ./config/.env

 .PHONY: compose-build
compose-build:
	docker-compose -f $(DOCKER_COMPOSE) --env-file ${DOCKER_ENV} build

 .PHONY: compose-up
compose-up:
	docker-compose -f $(DOCKER_COMPOSE) --env-file ${DOCKER_ENV} up

 .PHONY: compose-down
compose-down:
	docker-compose -f $(DOCKER_COMPOSE) --env-file ${DOCKER_ENV} down --remove-orphans

 .PHONY: compose-logs
compose-logs:
	docker-compose -f $(DOCKER_COMPOSE) --env-file ${DOCKER_ENV} logs -f