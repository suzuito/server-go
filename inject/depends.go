package inject

import (
	"context"

	"github.com/suzuito/server-go/matcher"
	"github.com/suzuito/server-go/setting"
	"github.com/suzuito/server-go/usecase"
	"golang.org/x/xerrors"
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
	mat, err := matcher.NewUserAgentMatcher([]string{
		"^.*googlebot.*$",
		"^.*twitterbot.*$",
		"^.*facebookexternalhit.*$",
	})
	if err != nil {
		closeFunc()
		return nil, nil, xerrors.Errorf(": %w")
	}
	r := GlobalDepends{
		UserAgentMatcher:             mat,
		ReverseProxyFactoryPrerender: &usecase.ReverseProxyFactoryImpl{Target: env.PrerenderURL},
		ReverseProxyFactoryFront:     &usecase.ReverseProxyFactoryImpl{Target: env.FrontURL},
	}
	return &r, closeFunc, nil
}
