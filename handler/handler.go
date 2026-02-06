package handler

import (
	"encoding/json"
	"log"
	"net/http"

	appErrors "go-maps-app/errors"
	"go-maps-app/service"
)

// HandleMap handles /map requests and centralizes error -> response mapping.
func HandleMap(w http.ResponseWriter, r *http.Request) {
	log.Printf("[REQUEST] %s %s?%s", r.Method, r.URL.Path, r.URL.RawQuery)

	location := r.URL.Query().Get("location")

	mapsURL, err := service.BuildMapURL(location)
	if err != nil {
		handleError(w, err)
		return
	}

	log.Printf("[SUCCESS] redirecting to %s", mapsURL)
	http.Redirect(w, r, mapsURL, http.StatusFound)
}

// Health is a simple root health-check endpoint.
func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok", "service": "map-redirect"})
}

// handleError maps service errors to HTTP responses.
func handleError(w http.ResponseWriter, err error) {
	var apiErr *appErrors.APIError
	if e, ok := err.(*appErrors.APIError); ok {
		apiErr = e
	} else {
		apiErr = appErrors.New(http.StatusInternalServerError, 2000, "internal server error")
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(apiErr.HTTPStatus)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{"code": apiErr.Code, "message": apiErr.Message})
}
