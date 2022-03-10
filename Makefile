GO_SOURCES := $(shell find . -name '*.go')

test:
	CGO_ENABLED=0 go test -coverpkg=$(go list ./... | grep -v mock | tr '\n' ',') -coverprofile=coverage.txt -covermode=atomic ./...
	go tool cover -func=coverage.txt
	go tool cover -html=coverage.txt -o coverage.html