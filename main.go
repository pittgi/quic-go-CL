package main

import (
	"log"
	"net/http"

	"github.com/quic-go/quic-go/http3"
)

func main() {
	// Create a basic HTTP handler that always responds with "hello world"
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	// ListenAndServeQUIC starts an HTTP/3 server on the given address.
	// It requires a TLS certificate and key.
	err := http3.ListenAndServeQUIC(":443", "cert.pem", "key.pem", mux)
	if err != nil {
		log.Fatal(err)
	}
}
