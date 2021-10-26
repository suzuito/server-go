package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", server.Handler())
	http.ListenAndServe(":8080", nil)
}
