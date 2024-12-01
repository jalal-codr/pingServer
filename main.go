package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
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

	if os.Getenv("RENDER") == "" { // Render sets RENDER env variable in production
		err := godotenv.Load()
		if err != nil {
			log.Println("No .env file found, running without .env")
		}
	}

	// Example URL to send the GET request to
	url := os.Getenv("API_URL")

	// Create a ticker that ticks every 30 seconds
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	// Start the server in a separate goroutine
	go StartServer()

	// Run the GET request every 30 seconds in the background
	for {
		select {
		case <-ticker.C:
			body, err := MakeGetRequest(url)
			if err != nil {
				log.Println("Error:", err)
				continue
			}
			// Print the response body
			fmt.Println("Response Body:", body)
		}
	}
}
