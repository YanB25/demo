all: build

test:
	@go test -v ./...
build:
	@go build
clean:
	@go clean -i
deploy:
	@go install