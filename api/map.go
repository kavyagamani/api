package main

import (
	"net/http"

	"go-maps-app/handler"
)

// Handler is the Vercel serverless entrypoint.
// Vercel will invoke this for the route configured in vercel.json.
func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/map", "/map/":
		handler.HandleMap(w, r)
	case "/", "":
		handler.Health(w, r)
	default:
		http.NotFound(w, r)
	}
}
