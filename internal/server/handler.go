package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/suzuito/server-go/internal/entity"
	"github.com/suzuito/server-go/internal/inject"
	"github.com/suzuito/server-go/internal/setting"
	"github.com/suzuito/server-go/internal/usecase"
)

func HandlerBlog(
	env *setting.Environment,
	gdeps *inject.GlobalDepends,
) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		le := entity.LogEntry{}
		le.Method = r.Method
		le.RemoteAddr = r.RemoteAddr
		le.StartedAt = time.Now()
		le.URI = r.URL.String()
		r = usecase.SetContextLogEntry(r, &le)
		defer func() {
			b, _ := json.Marshal(le)
			fmt.Println(string(b))
		}()

		if r.URL.Path == "/sitemap.xml" {
			q := r.URL.Query()
			q.Set("origin", env.SitemapOrigin)
			r.URL.RawQuery = q.Encode()
			gdeps.ReverseProxySitemap.ServeHTTP(w, r)
			return
		}
		if gdeps.HealthCheckBotMatcher.IsMatched(r.Header.Get("user-agent")) {
			le.TargetStatusCode = http.StatusOK
			fmt.Fprintf(w, "ok\n")
			le.TargetResponsedAt = time.Now()
			return
		}
		if gdeps.ExternalAppBotMatcher.IsMatched(r.Header.Get("user-agent")) {
			r.URL.Path = fmt.Sprintf("/render/https://%s%s", r.Host, r.URL.Path)
			gdeps.ReverseProxyPrerender.ServeHTTP(w, r)
			return
		}
		result, _ := regexp.MatchString(`\.txt$|\.png$|\.html$|\.js$|\.xml$|\.css$`, r.URL.Path)
		if result {
			gdeps.ReverseProxyFront.ServeHTTP(w, r)
			return
		}
		r.URL.Path = "/index.html"
		gdeps.ReverseProxyFront.ServeHTTP(w, r)
	}
}
