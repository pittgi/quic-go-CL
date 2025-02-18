package main

import (
	"io"
	"log"
	"net/http"

	"github.com/quic-go/quic-go/http3"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// Read the entire body from the POST request.
			body, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Failed to read request body", http.StatusInternalServerError)
				return
			}
			// Log the received request body.
			log.Printf("Received POST request body: %s", string(body))
		}
		// Send the response.
		w.Write([]byte("hello world"))
	})

	// Start the HTTP/3 server on port 443 with TLS.
	err := http3.ListenAndServeQUIC(":443", "cert.pem", "key.pem", mux)
	if err != nil {
		log.Fatal(err)
	}
}
