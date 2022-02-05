package usecase

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
)

type DependsMock struct {
	ExternalAppBotMatcher        *MockUserAgentMatcher
	HealthCheckBotMatcher        *MockUserAgentMatcher
	ReverseProxy                 *MockReverseProxy
	ReverseProxyFactoryFront     *MockReverseProxyFactory
	ReverseProxyFactoryPrerender *MockReverseProxyFactory
}

func NewDependsMock(t *testing.T) (*DependsMock, func()) {
	r := DependsMock{}
	ctrlExternalAppBotMatcher := gomock.NewController(t)
	ctrlHealthCheckBotMatcher := gomock.NewController(t)
	ctrlReverseProxy := gomock.NewController(t)
	ctrlReverseProxyFactoryFront := gomock.NewController(t)
	ctrlReverseProxyFactoryPrerender := gomock.NewController(t)
	r.ExternalAppBotMatcher = NewMockUserAgentMatcher(ctrlExternalAppBotMatcher)
	r.HealthCheckBotMatcher = NewMockUserAgentMatcher(ctrlHealthCheckBotMatcher)
	r.ReverseProxy = NewMockReverseProxy(ctrlReverseProxy)
	r.ReverseProxyFactoryFront = NewMockReverseProxyFactory(ctrlReverseProxyFactoryFront)
	r.ReverseProxyFactoryPrerender = NewMockReverseProxyFactory(ctrlReverseProxyFactoryPrerender)
	return &r, func() {
		ctrlExternalAppBotMatcher.Finish()
		ctrlHealthCheckBotMatcher.Finish()
		ctrlReverseProxy.Finish()
		ctrlReverseProxyFactoryFront.Finish()
		ctrlReverseProxyFactoryPrerender.Finish()
	}
}
