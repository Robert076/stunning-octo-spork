package main

import (
	"fmt"
	"net/http"

	"github.com/robert076/stunning-octo-spork/proxy-service/handler"
)

func main() {
	http.HandleFunc("/proxy", handler.ProxyHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error setting up http server")
	}
}
