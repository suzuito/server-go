package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/suzuito/blog1-server-go/entity"
	"github.com/suzuito/blog1-server-go/inject"
	"github.com/suzuito/blog1-server-go/setting"
	"github.com/suzuito/blog1-server-go/usecase"
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
			gdep.ReverseProxyFactoryPrerender.NewReverseProxy(usecase.NewRoundTripperImpl(&lentry)).ServeHTTP(w, r)
			lentry.ResponsedAt = time.Now()
			return
		}
		gdep.ReverseProxyFactoryFront.NewReverseProxy(usecase.NewRoundTripperImpl(&lentry)).ServeHTTP(w, r)
		lentry.ResponsedAt = time.Now()
	}
}
