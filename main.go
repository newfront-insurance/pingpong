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
		log.Println("Received /ping request")
		fmt.Fprintf(w, "pong")
	})

	mux.HandleFunc("/ping/wait/longest", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received /ping/wait/longest request")
		time.Sleep(70 * time.Second)
		fmt.Fprintf(w, "Waited for 70 seconds")
	})

	mux.HandleFunc("/ping/wait/longer", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received /ping/wait/longer request")
		time.Sleep(60 * time.Second)
		fmt.Fprintf(w, "Waited for 60 seconds")
	})

	mux.HandleFunc("/ping/wait/long", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received /ping/wait/long request")
		time.Sleep(45 * time.Second)
		fmt.Fprintf(w, "Waited for 45 seconds")
	})

	mux.HandleFunc("/ping/wait/short", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received /ping/wait/short request")
		time.Sleep(10 * time.Second)
		fmt.Fprintf(w, "Waited for 10 seconds")
	})

	// New endpoint that waits for 10 seconds and then returns a 504
	mux.HandleFunc("/ping/wait/timeout", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received /ping/wait/timeout request")
		time.Sleep(10 * time.Second)
		http.Error(w, "Gateway Timeout", http.StatusGatewayTimeout)
	})

	// Endpoint that waits for 10 seconds and then returns a 502
	mux.HandleFunc("/ping/wait/badgateway", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received /ping/wait/badgateway request")
		time.Sleep(10 * time.Second)
		http.Error(w, "Bad Gateway", http.StatusBadGateway)
	})

	// Endpoint to simulate connection refused
	mux.HandleFunc("/ping/wait/refuse", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received /ping/wait/refuse request")
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
