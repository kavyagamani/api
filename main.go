package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// handleMap handles the /map endpoint
func handleMap(w http.ResponseWriter, r *http.Request) {
	// Debug logging to confirm handler is being called
	log.Printf("[DEBUG] Handler called! Method: %s, Path: %s, Query: %s\n",
		r.Method, r.URL.Path, r.URL.RawQuery)

	// Get the location query parameter
	location := r.URL.Query().Get("location")

	// Validate that location is not empty
	if location == "" {
		log.Printf("[DEBUG] Missing location parameter\n")
		http.Error(w, "location parameter is required", http.StatusBadRequest)
		return
	}

	log.Printf("[DEBUG] Location parameter: %s\n", location)

	// Build the Google Maps URL with URL-encoded location
	mapsURL := "https://www.google.com/maps/search/?api=1&query=" + url.QueryEscape(location)

	log.Printf("[DEBUG] Redirecting to: %s\n", mapsURL)

	// Redirect to Google Maps
	http.Redirect(w, r, mapsURL, http.StatusFound)
}

func main() {
	// Register the handler for BOTH /map and /map/ to avoid trailing slash issues
	http.HandleFunc("/map", handleMap)
	http.HandleFunc("/map/", handleMap)

	// Start the server on port 8080
	port := ":8080"
	fmt.Printf("Server starting on http://localhost:8080\n")
	fmt.Printf("Test with: http://localhost:8080/map?location=Bangalore\n")
	fmt.Printf("or http://localhost:8080/map/?location=Bangalore\n")
	log.Fatal(http.ListenAndServe(port, nil))
}
