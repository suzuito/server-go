package main

import (
	"context"
	"net/http"
	"os"

	"github.com/suzuito/blog1-server-go/inject"
	"github.com/suzuito/blog1-server-go/server"
	"github.com/suzuito/blog1-server-go/setting"
)

func main() {
	ctx := context.Background()
	env, err := setting.NewEnvironment()
	if err != nil {
		os.Exit(1)
	}
	gdep, closeFunc, err := inject.NewGlobalDepends(ctx, env)
	if err != nil {
		os.Exit(1)
	}
	defer closeFunc()
	http.HandleFunc("/", server.Handler(env, gdep))
	http.ListenAndServe(":8080", nil)
}
