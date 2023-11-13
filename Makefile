.PHONY: run
validator:
	go run cmd/validator/main.go

repository:
	go run cmd/repository/main.go

bot:
	go run cmd/bot/main.go

config:
	cp internal/config/config.json.example internal/config/config.json

build-bot:
	go build -o bin/bot cmd/bot/main.go


.PHONY: test
test:
	go test -v ./...

integration-test:
	go test -v -tags=integration ./tests/... -parallel 1

LOCAL_BIN:=$(CURDIR)/bin
.PHONY: .deps
.deps:
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway && \
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 && \
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go && \
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

MIGRATIONS_DIR=./migrations
.PHONY: migration
migration:
	goose -dir=${MIGRATIONS_DIR} create $(NAME) sql