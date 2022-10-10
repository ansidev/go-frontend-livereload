.PHONY: prepare

prepare:
	@echo "Installing golangci-lint"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "Installing Husky"
	@go install github.com/go-courier/husky/cmd/husky@latest && husky init
	@echo "Installing Taskfile"
	@go install github.com/go-task/task/v3/cmd/task@latest
	@echo "Installing Air"
	@go install github.com/cosmtrek/air@latest
