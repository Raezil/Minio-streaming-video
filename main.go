package main

import (
	"fmt"
	"log"
	"net/http"
	. "pkg"
)

const (
	streamEndpoint = "GET /video" // HTTP endpoint for streaming

)

func main() {
	streaming := NewStreaming()
	http.HandleFunc(streamEndpoint, func(w http.ResponseWriter, r *http.Request) {
		streaming.Stream(w, r)
	})

	port := "8080"
	fmt.Printf("Server is listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
