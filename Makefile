## Makefile - helpers for running example modules (modules = dirs containing main.go)
# Usage:
#   make list                # list detected modules (dirs with main.go)
#   make run DIR=hello-world # run a specific module
#   make run-all             # run every module found (stops on first failure)
#   make build-all           # build every module into ./bin
#   make test                # run `go test ./...`

.PHONY: help list run run-all build-all test

help: ## Show this help
	@sed -n 's/^\(.*\):.*##\s*\(.*\)/\1 - \2/p' $(MAKEFILE_LIST)

list: ## List detected modules (paths that contain a main.go)
	@find . -type f -name main.go 2>/dev/null | sed 's:/main.go::' | sed 's|^./||' | sort -u



run: ## Run a module. Usage: make run DIR=path/to/module
	@if [ -z "$(DIR)" ]; then echo "Provide DIR=<path-to-module> (example: hello-world)"; exit 1; fi
	@if [ ! -d "$(DIR)" ]; then echo "Directory '$(DIR)' not found"; exit 1; fi
	@cd "$(DIR)" && go run main.go


run-all: ## Run every module containing main.go (one by one)
	@for d in $(shell find . -type f -name main.go 2>/dev/null | sed 's:/main.go::' | sed 's|^./||' | sort -u); do \
		if [ -d "$$d" ]; then \
			echo "==> $$d"; \
			cd "$$d" && go run main.go || exit 1; \
		else \
			echo "skipping missing: $$d"; \
		fi; \
	done



build-all: ## Build all modules into bin/ (name = dir path with / -> -)
	@mkdir -p bin
	@for d in $(shell find . -type f -name main.go 2>/dev/null | sed 's:/main.go::' | sed 's|^./||' | sort -u); do \
		if [ -d "$$d" ]; then \
			name=$$(echo $$d | tr '/' '-'); \
			echo "building $$d -> bin/$$name"; \
			cd "$$d" && go build -o ../../bin/$$name . || exit 1; \
		else \
			echo "skipping missing: $$d"; \
		fi; \
	done

test: ## Run all tests in module
	@go test ./...
