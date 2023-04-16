.PHONY: help
help: ## show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: wire
wire: ## Generate wire
	@wire ./internal/provider/.

.PHONY: build
build: ## Build binary
	@go build -o bin/aquafarm-rest ./cmd/aquafarm-rest

.PHONY: run
run: build ## Build & run the app
	@./bin/aquafarm-rest

.PHONY: test
test: ## Run unit test with cover
	@go test ./... -v --cover

.PHONY: test-report
test-report: ## Run unit test and show coverage report
	@go test ./... -v --cover -coverprofile=coverage.out
	@go tool cover -html=coverage.out

.PHONY: lint
lint: ## Run linters
	@golangci-lint run

.PHONY: docker-rebuild-app
docker-rebuild-app: ## Rebuild docker app
	@docker-compose up --build --force-recreate --no-deps -d app
