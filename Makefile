GO_ENV=CGO_ENABLED=0 GO111MODULE=on
GO=$(GO_ENV) $(shell which go)
DIR_SRC=./cmd/vfs
BIN=vfs

build: generate
	@$(GO) build -o $(BIN) $(DIR_SRC)

generate:
	@$(GO) mod tidy