package usecase

import (
	"fmt"
	"net/http"
	"regexp"
)

func NewProxyBlog(
	healthCheckBotMatcher UserAgentMatcher,
	externalAppBotMatcher UserAgentMatcher,
	reverseProxyFront ReverseProxy,
	reverseProxyPrerendering ReverseProxy,
	env string,
) *ReverseProxyRoutes {
	return &ReverseProxyRoutes{
		Routes: []ReverseProxyRoute{
			&ProxyHealthCheck{
				healthCheckBotMatcher: healthCheckBotMatcher,
			},
			&ReverseProxyRouteBlogPrerendering{
				externalAppBotMatcher: externalAppBotMatcher,
				pxy:                   reverseProxyPrerendering,
				env:                   env,
			},
			&ReverseProxyRouteBlogGCS1{
				pxy: reverseProxyFront,
			},
			&ReverseProxyRouteBlogGCS2{
				pxy: reverseProxyFront,
			},
		},
	}
}

type ReverseProxyRouteBlogPrerendering struct {
	pxy ReverseProxy

	externalAppBotMatcher UserAgentMatcher

	env string
}

func (p *ReverseProxyRouteBlogPrerendering) Check(r *http.Request) bool {
	ua := r.Header.Get("user-agent")
	if p.externalAppBotMatcher.IsMatched(ua) {
		return true
	}
	return false
}

func (p *ReverseProxyRouteBlogPrerendering) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	scheme := "https"
	if p.env == "dev" {
		scheme = "http"
	}
	r.URL.Path = fmt.Sprintf("/render/%s://%s%s", scheme, r.Host, r.URL.Path)
	p.pxy.ServeHTTP(w, r)
}

type ReverseProxyRouteBlogGCS1 struct {
	pxy ReverseProxy
}

func (p *ReverseProxyRouteBlogGCS1) Check(r *http.Request) bool {
	result, _ := regexp.MatchString(`\.txt$|\.png$|\.html$|\.js$|\.xml$|\.css$`, r.URL.Path)
	return result
}

func (p *ReverseProxyRouteBlogGCS1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.pxy.ServeHTTP(w, r)
}

type ReverseProxyRouteBlogGCS2 struct {
	pxy ReverseProxy
}

func (p *ReverseProxyRouteBlogGCS2) Check(r *http.Request) bool {
	return true
}

func (p *ReverseProxyRouteBlogGCS2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = "/index.html"
	p.pxy.ServeHTTP(w, r)
}
