package inject

import (
	"context"
	"net/http/httputil"

	"github.com/suzuito/blog1-server-go/matcher"
	"github.com/suzuito/blog1-server-go/setting"
	"github.com/suzuito/blog1-server-go/usecase"
)

type GlobalDepends struct {
	UserAgentMatcher      usecase.UserAgentMatcher
	ReverseProxyPrerender usecase.ReverseProxy
	ReverseProxyFront     usecase.ReverseProxy
}

func NewGlobalDepends(ctx context.Context, env *setting.Environment) (*GlobalDepends, func(), error) {
	closeFuncs := []func(){}
	closeFunc := func() {
		for _, cf := range closeFuncs {
			cf()
		}
	}
	r := GlobalDepends{
		UserAgentMatcher:      matcher.NewUserAgentMatcher(),
		ReverseProxyPrerender: httputil.NewSingleHostReverseProxy(env.PrerenderURL),
		ReverseProxyFront:     httputil.NewSingleHostReverseProxy(env.FrontURL),
	}
	return &r, closeFunc, nil
}
