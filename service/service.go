package service

import (
	"net/url"

	appErrors "go-maps-app/errors"
)

// BuildMapURL constructs a Google Maps search URL for the given location.
// It validates input and returns an error (possibly an *errors.APIError) on failure.
func BuildMapURL(location string) (string, error) {
	if location == "" {
		return "", appErrors.ErrMissingLocation
	}

	base := "https://www.google.com/maps/search/"
	u := base + "?api=1&query=" + url.QueryEscape(location)
	return u, nil
}
