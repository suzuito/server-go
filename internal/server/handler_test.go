package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/suzuito/server-go/internal/inject"
	"github.com/suzuito/server-go/internal/setting"
	mock_usecase "github.com/suzuito/server-go/internal/usecase/mock"
)

func TestHandlerBlog(t *testing.T) {
	testCases := []struct {
		desc            string
		inputHTTPMethod string
		inputURL        string
		setUp           func(dep *mock_usecase.DependsMock)
	}{
		{
			desc:            `成功 ヘルスチェックリクエスト`,
			inputHTTPMethod: http.MethodGet,
			inputURL:        "http://foo.co.jp",
			setUp: func(dep *mock_usecase.DependsMock) {
				dep.HealthCheckBotMatcher.EXPECT().
					IsMatched(gomock.Any()).
					Return(true)
			},
		},
		{
			desc:            `成功 ボットではないリクエストはFrontへ`,
			inputHTTPMethod: http.MethodGet,
			inputURL:        "http://foo.co.jp",
			setUp: func(dep *mock_usecase.DependsMock) {
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
			setUp: func(dep *mock_usecase.DependsMock) {
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
			setUp: func(dep *mock_usecase.DependsMock) {
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
		{
			desc:            `成功 /sitemap.xmlはサイトマップへ`,
			inputHTTPMethod: http.MethodGet,
			inputURL:        "http://foo.co.jp/sitemap.xml",
			setUp: func(dep *mock_usecase.DependsMock) {
				dep.ReverseProxySitemap.EXPECT().
					ServeHTTP(gomock.Any(), gomock.Any())
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			dep, closeFunc := mock_usecase.NewDependsMock(t)
			defer closeFunc()
			tC.setUp(dep)
			env := setting.Environment{
				Env: "dev",
			}
			gdep := inject.GlobalDepends{
				ExternalAppBotMatcher: dep.ExternalAppBotMatcher,
				HealthCheckBotMatcher: dep.HealthCheckBotMatcher,
				ReverseProxyPrerender: dep.ReverseProxyPrerendering,
				ReverseProxyFront:     dep.ReverseProxyFront,
				ReverseProxySitemap:   dep.ReverseProxySitemap,
			}
			recorder := httptest.NewRecorder()
			h := HandlerBlog(&env, &gdep)
			req, err := http.NewRequest(tC.inputHTTPMethod, tC.inputURL, nil)
			if err != nil {
				t.Error(err)
				return
			}
			h(recorder, req)
		})
	}
}
