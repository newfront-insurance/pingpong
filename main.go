package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	mux.HandleFunc("/wait", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(60 * time.Second) // Wait for 60 seconds
		fmt.Fprintf(w, "Waited for 60 seconds")
	})

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		IdleTimeout:  10 * time.Second, // The maximum amount of time to wait for the next request when keep-alives are enabled
		ReadTimeout:  10 * time.Second, // The maximum duration for reading the entire request, including the body
		WriteTimeout: 10 * time.Second, // The maximum duration before timing out writes of the response
	}

	log.Fatal(server.ListenAndServe())
}
