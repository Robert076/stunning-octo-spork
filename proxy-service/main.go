package main

import (
	"net/http"

	"github.com/robert076/stunning-octo-spork/proxy-service/handler"
)

func main() {
	http.HandleFunc("/proxy", handler.ProxyHandler)

	http.ListenAndServe(":8080", nil)
}
