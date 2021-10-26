GO_SOURCES := $(shell find . -name '*.go')

server.exe: ${GO_SOURCES}
	go build -o server.exe cmd/server/main.go
