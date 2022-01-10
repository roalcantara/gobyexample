.PHONY: hello-world
hello-world: ## Runs hello-world/main.go
	go run src/hello-world/main.go

.PHONY: values
values: ## Runs values/main.go
	go run src/values/main.go