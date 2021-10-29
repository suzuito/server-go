package usecase

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
)

type DependsMock struct {
	UserAgentMatcher             *MockUserAgentMatcher
	ReverseProxy                 *MockReverseProxy
	ReverseProxyFactoryFront     *MockReverseProxyFactory
	ReverseProxyFactoryPrerender *MockReverseProxyFactory
}

func NewDependsMock(t *testing.T) (*DependsMock, func()) {
	r := DependsMock{}
	ctrlUserAgentMatcher := gomock.NewController(t)
	ctrlReverseProxy := gomock.NewController(t)
	ctrlReverseProxyFactoryFront := gomock.NewController(t)
	ctrlReverseProxyFactoryPrerender := gomock.NewController(t)
	r.UserAgentMatcher = NewMockUserAgentMatcher(ctrlUserAgentMatcher)
	r.ReverseProxy = NewMockReverseProxy(ctrlReverseProxy)
	r.ReverseProxyFactoryFront = NewMockReverseProxyFactory(ctrlReverseProxyFactoryFront)
	r.ReverseProxyFactoryPrerender = NewMockReverseProxyFactory(ctrlReverseProxyFactoryPrerender)
	return &r, func() {
		ctrlUserAgentMatcher.Finish()
		ctrlReverseProxy.Finish()
		ctrlReverseProxyFactoryFront.Finish()
		ctrlReverseProxyFactoryPrerender.Finish()
	}
}
