all: build

test:
	@go test ./...
build:
	@go build
clean:
	@go clean