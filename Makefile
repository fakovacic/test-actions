.PHONY: build
build:
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o ./cmd/app ./cmd

.PHONY: test
test:
	@go test -race -cover ./internal/...