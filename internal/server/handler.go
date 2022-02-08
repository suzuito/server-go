package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/suzuito/server-go/internal/entity"
	"github.com/suzuito/server-go/internal/setting"
	"github.com/suzuito/server-go/internal/usecase"
)

func Handler(env *setting.Environment, pxy *usecase.ReverseProxyRoutes) func(http.ResponseWriter, *http.Request) {
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
		for _, entry := range pxy.Routes {
			if entry.Check(r) {
				entry.ServeHTTP(w, r)
				return
			}
		}
		if pxy.DefaultRoute != nil {
			pxy.DefaultRoute.ServeHTTP(w, r)
			return
		}
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "not found\n")
	}
}
