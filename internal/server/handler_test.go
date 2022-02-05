package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/suzuito/server-go/internal/inject"
	"github.com/suzuito/server-go/internal/setting"
	"github.com/suzuito/server-go/internal/usecase"
)

func TestHandler(t *testing.T) {
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
				dep.ReverseProxyFactoryFront.EXPECT().
					NewReverseProxy(gomock.Any()).
					Return(dep.ReverseProxy)
				dep.ReverseProxy.EXPECT().
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
				dep.ReverseProxyFactoryPrerender.EXPECT().
					NewReverseProxy(gomock.Any()).
					Return(dep.ReverseProxy)
				dep.ReverseProxy.EXPECT().
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
			h := Handler(&env, &inject.GlobalDepends{
				ExternalAppBotMatcher:        dep.ExternalAppBotMatcher,
				HealthCheckBotMatcher:        dep.HealthCheckBotMatcher,
				ReverseProxyFactoryFront:     dep.ReverseProxyFactoryFront,
				ReverseProxyFactoryPrerender: dep.ReverseProxyFactoryPrerender,
			})
			req, err := http.NewRequest(tC.inputHTTPMethod, tC.inputURL, nil)
			if err != nil {
				t.Error(err)
				return
			}
			h(recorder, req)
		})
	}
}
