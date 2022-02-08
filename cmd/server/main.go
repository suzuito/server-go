package main

import (
	"context"
	"net/http"
	"os"

	"github.com/suzuito/server-go/internal/inject"
	"github.com/suzuito/server-go/internal/server"
	"github.com/suzuito/server-go/internal/setting"
	"github.com/suzuito/server-go/internal/usecase"
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

	pxy := usecase.NewProxyBlog(
		gdep.HealthCheckBotMatcher,
		gdep.ExternalAppBotMatcher,
		gdep.ReverseProxyFront,
		gdep.ReverseProxyPrerender,
		env.Env,
	)

	http.HandleFunc("/", server.Handler(env, pxy))
	http.ListenAndServe(":8080", nil)
}
