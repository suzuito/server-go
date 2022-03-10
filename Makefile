GO_SOURCES := $(shell find . -name '*.go')

mock:
	make internal/usecase/reverse_proxy_mock.go
	make internal/usecase/user_agent_matcher_mock.go

%_mock.go: %.go
	bash mockgen.sh $^

test:
	CGO_ENABLED=0 go test -coverpkg=github.com/suzuito/server-go/... -coverprofile=coverage.txt ./...
	go tool cover -func=coverage.txt
	go tool cover -html=coverage.txt -o coverage.html