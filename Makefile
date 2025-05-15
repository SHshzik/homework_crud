include .env.example
export

LOCAL_BIN:=$(CURDIR)/bin
PATH:=$(LOCAL_BIN):$(PATH)
BASE_STACK = docker compose -f docker-compose.yml

# HELP =================================================================================================================
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

compose-up: ### Run docker compose (without backend and reverse proxy)
	$(BASE_STACK) up --build -d db
.PHONY: compose-up

compose-down: ### Down docker compose
	$(BASE_STACK) down --remove-orphans
.PHONY: compose-down

swag-v1: ### swag init
	swag init -g services/user-server/controller/http/router.go
.PHONY: swag-v1

deps: ### deps tidy + verify
	go mod tidy && go mod verify
.PHONY: deps

user-service-run: deps swag-v1 ### swag run for API v1
	go mod download && \
	CGO_ENABLED=0 go run ./services/user-server/cmd
.PHONY: user-service-run

user-client-run:
	go mod download && \
	CGO_ENABLED=0 go run ./services/user-client/cmd
.PHONY: user-client-run

deps-audit: ### check dependencies vulnerabilities
	govulncheck ./...
.PHONY: deps-audit

format: ### Run code formatter
	gofumpt -l -w .
	gci write . --skip-generated -s standard -s default
.PHONY: format

docker-rm-volume: ### remove docker volume
	docker volume rm homework_crud_db_data
.PHONY: docker-rm-volume

linter-golangci: ### check by golangci linter
	golangci-lint run
.PHONY: linter-golangci

linter-hadolint: ### check by hadolint linter
	git ls-files --exclude='Dockerfile*' -i -o | xargs hadolint
.PHONY: linter-hadolint

linter-dotenv: ### check by dotenv linter
	dotenv-linter
.PHONY: linter-dotenv

bin-deps: ### install tools
	GOBIN=$(LOCAL_BIN) go install github.com/golang/mock/mockgen@latest
.PHONY: bin-deps

psql:
	docker-compose exec db psql -d db -U user

evans:
	~/evans/evans --proto ./api/proto/users.proto --port 8082 repl
.PHONY: evans

protoc:
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	api/proto/users.proto
.PHONY: protoc

# compose-up-all: ### Run docker compose (with backend and reverse proxy)
# 	$(BASE_STACK) up --build -d
# .PHONY: compose-up-all
