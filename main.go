package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	http.HandleFunc("/wait", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(60 * time.Second) // Wait for 60 seconds
		fmt.Fprintf(w, "Waited for 60 seconds")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
