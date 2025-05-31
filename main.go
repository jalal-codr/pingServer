package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// MakeGetRequest makes an HTTP GET request to the given URL and returns the response body or an error.
func MakeGetRequest(url string) (string, error) {
	// Make the GET request
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error making GET request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	// Return the response as a string
	return string(body), nil
}

// StartServer will create an HTTP server that listens on the specified port.
func StartServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Simple response to HTTP requests
		fmt.Fprintf(w, "Server is running. Making periodic GET requests in the background.")
	})

	// Start the server on port 8080
	log.Println("Server starting on :7080...")
	if err := http.ListenAndServe(":7080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func main() {

	// Example URL to send the GET request to
	url := "https://tradingbot-2w81.onrender.com/"

	// Create a ticker that ticks every 9 minutes
	ticker := time.NewTicker(9 * time.Minute)
	defer ticker.Stop()

	// Start the server in a separate goroutine
	go StartServer()

	// Initialize lastRequestTime with the current time
	lastRequestTime := time.Now()
	log.Println("Starting periodic requests. First request will be made in 9 minutes.")

	// Run the GET request every 9 minutes in the background
	for range ticker.C {
		// Calculate and log the time since the last request
		currentTime := time.Now()
		elapsedTime := currentTime.Sub(lastRequestTime)
		log.Printf("Time since last request: %v", elapsedTime)

		// Update lastRequestTime to the current time
		lastRequestTime = currentTime

		// Log that we're making a request
		log.Printf("Making GET request to %s", url)

		// Make the request
		body, err := MakeGetRequest(url)
		if err != nil {
			log.Println("Error:", err)
			continue
		}

		// Print the response body
		log.Println("Response received at:", currentTime.Format(time.RFC3339))
		log.Println("Response Body:", body)
	}
}
