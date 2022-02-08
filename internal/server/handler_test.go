package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/suzuito/server-go/internal/setting"
	"github.com/suzuito/server-go/internal/usecase"
)

func TestHandlerBlog(t *testing.T) {
	testCases := []struct {
		desc            string
		inputHTTPMethod string
		inputURL        string
		setUp           func(dep *usecase.DependsMock)
	}{
		{
			desc:            `成功 ヘルスチェックリクエスト`,
			inputHTTPMethod: http.MethodGet,
			inputURL:        "http://foo.co.jp",
			setUp: func(dep *usecase.DependsMock) {
				dep.HealthCheckBotMatcher.EXPECT().
					IsMatched(gomock.Any()).
					Return(true)
			},
		},
		{
			desc:            `成功 ボットではないリクエストはFrontへ`,
			inputHTTPMethod: http.MethodGet,
			inputURL:        "http://foo.co.jp",
			setUp: func(dep *usecase.DependsMock) {
				dep.HealthCheckBotMatcher.EXPECT().
					IsMatched(gomock.Any()).
					Return(false)
				dep.ExternalAppBotMatcher.EXPECT().
					IsMatched(gomock.Any()).
					Return(false)
				dep.ReverseProxyFront.EXPECT().
					ServeHTTP(gomock.Any(), gomock.Any())
			},
		},
		{
			desc:            `成功 ボットではないリクエストはFrontへ`,
			inputHTTPMethod: http.MethodGet,
			inputURL:        "http://foo.co.jp/hoge.js",
			setUp: func(dep *usecase.DependsMock) {
				dep.HealthCheckBotMatcher.EXPECT().
					IsMatched(gomock.Any()).
					Return(false)
				dep.ExternalAppBotMatcher.EXPECT().
					IsMatched(gomock.Any()).
					Return(false)
				dep.ReverseProxyFront.EXPECT().
					ServeHTTP(gomock.Any(), gomock.Any())
			},
		},
		{
			desc:            `成功 ボットであるリクエストはPrerenderへ`,
			inputHTTPMethod: http.MethodGet,
			inputURL:        "http://foo.co.jp",
			setUp: func(dep *usecase.DependsMock) {
				dep.HealthCheckBotMatcher.EXPECT().
					IsMatched(gomock.Any()).
					Return(false)
				dep.ExternalAppBotMatcher.EXPECT().
					IsMatched(gomock.Any()).
					Return(true)
				dep.ReverseProxyPrerendering.EXPECT().
					ServeHTTP(gomock.Any(), gomock.Any())
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			dep, closeFunc := usecase.NewDependsMock(t)
			defer closeFunc()
			tC.setUp(dep)
			env := setting.Environment{
				Env: "dev",
			}
			recorder := httptest.NewRecorder()
			h := Handler(&env, usecase.NewProxyBlog(
				dep.HealthCheckBotMatcher,
				dep.ExternalAppBotMatcher,
				dep.ReverseProxyFront,
				dep.ReverseProxyPrerendering,
				env.Env,
			))
			req, err := http.NewRequest(tC.inputHTTPMethod, tC.inputURL, nil)
			if err != nil {
				t.Error(err)
				return
			}
			h(recorder, req)
		})
	}
}

func newRoutes(t *testing.T, c int) ([]*usecase.MockReverseProxyRoute, *usecase.MockReverseProxyRoute, func()) {
	ctrls := []*gomock.Controller{}
	routes := []*usecase.MockReverseProxyRoute{}
	for i := 0; i < c; i++ {
		ctrl := gomock.NewController(t)
		ctrls = append(ctrls, ctrl)
		route := usecase.NewMockReverseProxyRoute(ctrl)
		routes = append(routes, route)
	}
	ctrlDefault := gomock.NewController(t)
	routeDefault := usecase.NewMockReverseProxyRoute(ctrlDefault)
	return routes, routeDefault, func() {
		for _, ctrl := range ctrls {
			ctrl.Finish()
		}
		ctrlDefault.Finish()
	}
}

func TestNoRoutesHandler(t *testing.T) {
	pxy := usecase.ReverseProxyRoutes{
		Routes:       []usecase.ReverseProxyRoute{},
		DefaultRoute: nil,
	}
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	if err != nil {
		t.Error(err)
		return
	}
	recorder := httptest.NewRecorder()
	Handler(&setting.Environment{Env: "dev"}, &pxy)(recorder, req)
	assert.Equal(t, http.StatusNotFound, recorder.Result().StatusCode)
}

func TestHandlerV2(t *testing.T) {
	testCases := []struct {
		desc            string
		inputHTTPMethod string
		inputURL        string
		inputC          int
		setUp           func(routes []*usecase.MockReverseProxyRoute, routeDefault *usecase.MockReverseProxyRoute)
	}{
		{
			desc:            `Routes:0, DefaultRoute:not nil`,
			inputHTTPMethod: http.MethodGet,
			inputURL:        "http://localhost:8080",
			inputC:          0,
			setUp: func(routes []*usecase.MockReverseProxyRoute, routeDefault *usecase.MockReverseProxyRoute) {
				routeDefault.EXPECT().ServeHTTP(gomock.Any(), gomock.Any())
			},
		},
		{
			desc:            `Routes:true,false, DefaultRoute:not nil`,
			inputHTTPMethod: http.MethodGet,
			inputURL:        "http://localhost:8080",
			inputC:          2,
			setUp: func(routes []*usecase.MockReverseProxyRoute, routeDefault *usecase.MockReverseProxyRoute) {
				routes[0].EXPECT().
					Check(gomock.Any()).Return(true)
				routes[0].EXPECT().
					ServeHTTP(gomock.Any(), gomock.Any())
			},
		},
		{
			desc:            `Routes:false,true, DefaultRoute:not nil`,
			inputHTTPMethod: http.MethodGet,
			inputURL:        "http://localhost:8080",
			inputC:          2,
			setUp: func(routes []*usecase.MockReverseProxyRoute, routeDefault *usecase.MockReverseProxyRoute) {
				routes[0].EXPECT().
					Check(gomock.Any()).Return(false)
				routes[1].EXPECT().
					Check(gomock.Any()).Return(true)
				routes[1].EXPECT().
					ServeHTTP(gomock.Any(), gomock.Any())
			},
		},
		{
			desc:            `Routes:false,false, DefaultRoute:not nil`,
			inputHTTPMethod: http.MethodGet,
			inputURL:        "http://localhost:8080",
			inputC:          2,
			setUp: func(routes []*usecase.MockReverseProxyRoute, routeDefault *usecase.MockReverseProxyRoute) {
				routes[0].EXPECT().
					Check(gomock.Any()).Return(false)
				routes[1].EXPECT().
					Check(gomock.Any()).Return(false)
				routeDefault.EXPECT().
					ServeHTTP(gomock.Any(), gomock.Any())
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			routes, routeDefault, closeFunc := newRoutes(t, tC.inputC)
			defer closeFunc()
			tC.setUp(routes, routeDefault)
			env := setting.Environment{
				Env: "dev",
			}
			tempRoutes := []usecase.ReverseProxyRoute{}
			for _, r := range routes {
				tempRoutes = append(tempRoutes, r)
			}
			recorder := httptest.NewRecorder()
			h := Handler(
				&env,
				&usecase.ReverseProxyRoutes{
					Routes:       tempRoutes,
					DefaultRoute: routeDefault,
				},
			)
			req, err := http.NewRequest(tC.inputHTTPMethod, tC.inputURL, nil)
			if err != nil {
				t.Error(err)
				return
			}
			h(recorder, req)
		})
	}
}
