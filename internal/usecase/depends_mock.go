package usecase

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
)

type DependsMock struct {
	ExternalAppBotMatcher    *MockUserAgentMatcher
	HealthCheckBotMatcher    *MockUserAgentMatcher
	ReverseProxySitemap      *MockReverseProxy
	ReverseProxyFront        *MockReverseProxy
	ReverseProxyPrerendering *MockReverseProxy
}

func NewDependsMock(t *testing.T) (*DependsMock, func()) {
	r := DependsMock{}
	ctrlExternalAppBotMatcher := gomock.NewController(t)
	ctrlHealthCheckBotMatcher := gomock.NewController(t)
	ctrlReverseProxyFront := gomock.NewController(t)
	ctrlReverseProxyPrerendering := gomock.NewController(t)
	ctrlReverseProxySitemap := gomock.NewController(t)
	r.ExternalAppBotMatcher = NewMockUserAgentMatcher(ctrlExternalAppBotMatcher)
	r.HealthCheckBotMatcher = NewMockUserAgentMatcher(ctrlHealthCheckBotMatcher)
	r.ReverseProxyFront = NewMockReverseProxy(ctrlReverseProxyFront)
	r.ReverseProxyPrerendering = NewMockReverseProxy(ctrlReverseProxyPrerendering)
	r.ReverseProxySitemap = NewMockReverseProxy(ctrlReverseProxySitemap)
	return &r, func() {
		ctrlExternalAppBotMatcher.Finish()
		ctrlHealthCheckBotMatcher.Finish()
		ctrlReverseProxyFront.Finish()
		ctrlReverseProxyPrerendering.Finish()
		ctrlReverseProxySitemap.Finish()
	}
}
