package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/suzuito/server-go/entity"
	"github.com/suzuito/server-go/inject"
	"github.com/suzuito/server-go/setting"
	"github.com/suzuito/server-go/usecase"
)

func Handler(env *setting.Environment, gdep *inject.GlobalDepends) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		lentry := entity.LogEntry{}
		lentry.StartedAt = time.Now()
		defer func() {
			lentryJSON, _ := json.Marshal(&lentry)
			fmt.Println(string(lentryJSON))
		}()
		lentry.Method = r.Method
		lentry.UserAgent = r.UserAgent()
		lentry.URI = r.URL.String()
		lentry.RemoteAddr = r.RemoteAddr
		ua := r.Header.Get("user-agent")
		if gdep.UserAgentMatcher.IsBot(ua) {
			scheme := "https"
			if env.Env == "dev" {
				scheme = "http"
			}
			r.URL.Path = fmt.Sprintf("/render/%s://%s%s", scheme, r.Host, r.URL.Path)
			gdep.ReverseProxyFactoryPrerender.NewReverseProxy(usecase.NewRoundTripperImpl(&lentry)).ServeHTTP(w, r)
			lentry.ResponsedAt = time.Now()
			return
		}
		if r.URL.Path == "/" || r.URL.Path == "" {
			r.URL.Path = "/index.html"
		}
		gdep.ReverseProxyFactoryFront.NewReverseProxy(usecase.NewRoundTripperImpl(&lentry)).ServeHTTP(w, r)
		lentry.ResponsedAt = time.Now()
	}
}
