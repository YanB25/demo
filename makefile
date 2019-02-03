all: fmt vet build

test: vet
	@go test -v ./...
build:
	@go build
clean:
	@go clean -i
deploy:
	@go install

fmt:
	@go fmt ./...
vet:
	@go vet ./...