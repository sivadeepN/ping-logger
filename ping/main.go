package main

import (
	"encoding/json"
	"io/ioutil"
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
	log.Printf("Decoding the body")

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	// Log the request body
	log.Printf("Request body: %s\n", body)

	// You can perform additional processing or logging here

	// Send a JSON response
	response := map[string]string{"message": "Ping received!", "requestBody": string(body)}
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
