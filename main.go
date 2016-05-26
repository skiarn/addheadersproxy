package main

import (
	"log"
	"net/http"
)

func main() {
	NewTestServer(8000)
	proxy := ReverseProxy("http://localhost:8000")
	mux := http.NewServeMux()
	mux.Handle("/", proxy)
	log.Fatal(http.ListenAndServe(":8090", mux))
}
