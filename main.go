package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/tests", testsHandler)

	http.ListenAndServe(":6969", mux)
}

func testsHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("New %s request from %s\n", r.Method, r.RemoteAddr)
	for name, headers := range r.Header {
		for _, h := range headers {
			log.Printf("Header -> %v: %v\n", name, h)
		}
	}
}
