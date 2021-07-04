# BUILD API-SERVER
.PHONY:
build:
				go build -v ./cmd/apiserver

.PHONY: test
test:
				go test -v -race -timeout 30s ./...

# DEFAULT MAKEFILE'S GOAL
.DEFAULT_GOAL := build

