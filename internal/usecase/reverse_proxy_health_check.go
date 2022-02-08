package usecase

import (
	"fmt"
	"net/http"
)

type ProxyHealthCheck struct {
	healthCheckBotMatcher UserAgentMatcher
}

func (p *ProxyHealthCheck) Check(r *http.Request) bool {
	ua := r.Header.Get("user-agent")
	if p.healthCheckBotMatcher.IsMatched(ua) {
		return true
	}
	return false
}

func (p *ProxyHealthCheck) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok\n")
}
