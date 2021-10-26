package usecase

import "net/http"

type ReverseProxy interface {
	ServeHTTP(rw http.ResponseWriter, req *http.Request)
}
