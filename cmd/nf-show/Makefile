SHELL := $(shell which bash)

BUILD_FLAGS = -trimpath -ldflags "-w -s"

.ONESHELL:
.NOTPARALLEL:
.EXPORT_ALL_VARIABLES:

.DEFAULT_GOAL := help

build: ## Build binary
	go build $(BUILD_FLAGS)

help: ## Show this help
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-16s\033[0m %s\n", $$1, $$2}'

memprofile: ## Run benchmarks with memory profiling
