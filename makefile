.DEFAULT_GOAL := lint

fmt:
	go fmt ./...
.PHONY:fmt

vet: fmt
	go vet ./...
	shadow ./...
.PHONY:vet

lint: vet
	golangci-lint run
.PHONY:lint

test: lint
	bash run-tests.sh
.PHONY:test
