BIN:=core-api

.PHONY: test
test:
	go test -race -cover ./...

.PHONY: test-ci
test-ci:
	make test

.PHONY: build
build:
	go build -o $(BIN) ./cmd/api/main.go

# ensures that dependencies have been tidied and vendored
.PHONY: ensure-deps
ensure-deps:
	@go mod download
	@go mod tidy
	@go mod vendor
	@git diff --exit-code