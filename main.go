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

	mux.HandleFunc("/ping/wait/longest", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(70 * time.Second)
		fmt.Fprintf(w, "Waited for 70 seconds")
	})

	mux.HandleFunc("/ping/wait/longer", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(60 * time.Second)
		fmt.Fprintf(w, "Waited for 60 seconds")
	})

	mux.HandleFunc("/ping/wait/long", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(45 * time.Second)
		fmt.Fprintf(w, "Waited for 45 seconds")
	})

	mux.HandleFunc("/ping/wait/short", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Second)
		fmt.Fprintf(w, "Waited for 10 seconds")
	})

	// New endpoint that waits for 10 seconds and then returns a 504
	mux.HandleFunc("/ping/wait/timeout", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Second)
		http.Error(w, "Gateway Timeout", http.StatusGatewayTimeout)
	})

	// Endpoint that waits for 10 seconds and then returns a 502
	mux.HandleFunc("/ping/wait/badgateway", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Second)
		http.Error(w, "Bad Gateway", http.StatusBadGateway)
	})

	// Endpoint to simulate connection refused
	mux.HandleFunc("/ping/wait/refuse", func(w http.ResponseWriter, r *http.Request) {
		conn, _, _ := w.(http.Hijacker).Hijack()
		conn.Close()
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
