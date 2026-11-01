include .env
export

PROJECT_ROOT := $(shell pwd)
export PROJECT_ROOT

.PHONY: help env-up env-down env-cleanup env-port-forward env-port-close migrate-create migrate-up migrate-down todoapp-run todoapp-deploy todoapp-undeploy

help: ## Показать список доступных команд
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

env-up: ## Запустить окружение
	@docker compose up -d todoapp-postgres

env-down: ## Остановить окружение
	@docker compose down todoapp-postgres

env-cleanup: ## Удалить данные окружения (volume)
	@read -p "Очистить volume? Опасность утери данных. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down todoapp-postgres port-forwarder && \
		rm -rf ${PROJECT_ROOT}/out/pgdata; \
	fi

env-port-forward: ## Открыть порт к базе
	@docker compose up -d port-forwarder

env-port-close: ## Закрыть порт
	@docker compose down port-forwarder

migrate-create: ## Создать новую версию схемы (make migrate-create seq=name)
ifndef seq
	$(error Параметр seq обязателен, например: make migrate-create seq=init)
endif
	@docker compose run --rm todoapp-postgres-migrate \
		create -ext sql -dir /migrations -seq "$(seq)"

migrate-up: ## Накатить миграции
	@docker compose run --rm todoapp-postgres-migrate \
		-path /migrations \
		-database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@todoapp-postgres:5432/$(POSTGRES_DB)?sslmode=disable&search_path=public" \
		up

migrate-down: ## Откатить миграции
	@docker compose run --rm todoapp-postgres-migrate \
		-path /migrations \
		-database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@todoapp-postgres:5432/$(POSTGRES_DB)?sslmode=disable&search_path=public" \
		down

todoapp-run: ## Запустить локально на хост-системе (нужен env-port-forward)
	@export LOGGER_FOLDER=${PROJECT_ROOT}/out/logs && \
	export POSTGRES_HOST=localhost && \
	go run ${PROJECT_ROOT}/cmd/app

todoapp-deploy: ## Собрать образ и запустить приложение в Docker Compose
	@docker compose up -d --build todoapp

todoapp-undeploy: ## Остановить сервис приложения
	@docker compose down todoapp
