package inject

import (
	"context"

	"github.com/suzuito/blog1-server-go/matcher"
	"github.com/suzuito/blog1-server-go/setting"
	"github.com/suzuito/blog1-server-go/usecase"
)

type GlobalDepends struct {
	UserAgentMatcher             usecase.UserAgentMatcher
	ReverseProxyFactoryPrerender usecase.ReverseProxyFactory
	ReverseProxyFactoryFront     usecase.ReverseProxyFactory
}

func NewGlobalDepends(ctx context.Context, env *setting.Environment) (*GlobalDepends, func(), error) {
	closeFuncs := []func(){}
	closeFunc := func() {
		for _, cf := range closeFuncs {
			cf()
		}
	}
	r := GlobalDepends{
		UserAgentMatcher:             matcher.NewUserAgentMatcher(),
		ReverseProxyFactoryPrerender: &usecase.ReverseProxyFactoryImpl{Target: env.PrerenderURL},
		ReverseProxyFactoryFront:     &usecase.ReverseProxyFactoryImpl{Target: env.FrontURL},
	}
	return &r, closeFunc, nil
}
