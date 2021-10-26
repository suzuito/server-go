package handler

import "net/http"

func Handler() func(http.ResponseWriter, *http.Request) {
	return func(http.ResponseWriter, *http.Request) {}
}
