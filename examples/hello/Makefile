.DEFAULT_GOAL := help

build: ## Build this package
	go build -o ./build/hello ./...

clean: ## Clean build artifacts
	rm -r ./build

help: ## Show help
	@echo
	@echo ">> Available commands:"
	@cat $(MAKEFILE_LIST) | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'