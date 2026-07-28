# go-standLib Makefile

GO  ?= go

.PHONY: help
help: ## 显示帮助
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-10s\033[0m %s\n", $$1, $$2}'

.PHONY: test
test: ## 跑测试
	$(GO) test ./...

.PHONY: vet
vet: ## 静态检查
	$(GO) vet ./...

.PHONY: build
build: ## 编译检查
	$(GO) build ./...

.PHONY: tidy
tidy: ## 整理依赖
	$(GO) mod tidy

.PHONY: run
run: ## 跑 main.go（演示 collections）
	$(GO) run main.go