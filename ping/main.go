package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", handlePing)
	log.Println("Ping Logger started. Listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	// Get the URL or IP address from where the ping is coming from
	url := r.Header.Get("X-Forwarded-For")
	if url == "" {
		url = r.RemoteAddr
	}

	// Log the ping information
	log.Printf("Ping request from %s\n", url)

	// You can perform additional processing or logging here

	// Send a response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Ping received!"))
}
