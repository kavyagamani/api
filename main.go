package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"go-maps-app/handler"
)

func main() {
	// Read port from environment with fallback to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := fmt.Sprintf(":%s", port)

	mux := http.NewServeMux()
	// Serve static files from ./static (index.html will be served at "/")
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.Handle("/", fs)

	// API route(s)
	mux.HandleFunc("/map", handler.HandleMap)
	mux.HandleFunc("/map/", handler.HandleMap)

	log.Printf("GO MAP REDIRECT API - starting on http://localhost:%s", port)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server failed to start on %s: %v", addr, err)
	}
}
