package inject

import (
	"context"

	"github.com/suzuito/server-go/internal/matcher"
	"github.com/suzuito/server-go/internal/setting"
	"github.com/suzuito/server-go/internal/usecase"
	"golang.org/x/xerrors"
)

type GlobalDepends struct {
	ExternalAppBotMatcher usecase.UserAgentMatcher
	HealthCheckBotMatcher usecase.UserAgentMatcher
	ReverseProxyPrerender usecase.ReverseProxy
	ReverseProxyFront     usecase.ReverseProxy
	ReverseProxySitemap   usecase.ReverseProxy
}

func NewGlobalDepends(ctx context.Context, env *setting.Environment) (*GlobalDepends, func(), error) {
	closeFuncs := []func(){}
	closeFunc := func() {
		for _, cf := range closeFuncs {
			cf()
		}
	}
	appMat, err := matcher.NewExternalBotMatcherDefault()
	if err != nil {
		closeFunc()
		return nil, nil, xerrors.Errorf(": %w")
	}
	healthMat, err := matcher.NewHealthCheckBotMatcherDefault()
	if err != nil {
		closeFunc()
		return nil, nil, xerrors.Errorf(": %w")
	}
	r := GlobalDepends{
		ExternalAppBotMatcher: appMat,
		HealthCheckBotMatcher: healthMat,
		ReverseProxyPrerender: usecase.NewSingleHostReverseProxy(env.PrerenderURL),
		ReverseProxyFront:     usecase.NewSingleHostReverseProxy(env.FrontURL),
		ReverseProxySitemap:   usecase.NewSingleHostReverseProxy(env.SitemapURL),
	}
	return &r, closeFunc, nil
}
