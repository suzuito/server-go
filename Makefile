GO_SOURCES := $(shell find . -name '*.go')

server.exe: ${GO_SOURCES}
	go build -o server.exe cmd/server/main.go

mock:
	make usecase/reverse_proxy_mock.go
	make usecase/user_agent_matcher_mock.go

%_mock.go: %.go
	bash mockgen.sh $^

test:
	CGO_ENABLED=0 go test -coverpkg=github.com/suzuito/server-go/... -coverprofile=cov.out ./...
	go tool cover -func=cov.out
	go tool cover -html=cov.out -o cov.html