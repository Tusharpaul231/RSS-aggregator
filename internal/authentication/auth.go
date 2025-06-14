package auth

import (
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	// GetAPIKey retrieves the API key from the request headers.
	// Exmaple implementation:
	// Authentication : APIKey {your-api-key-here}
	val := headers.Get("Authentication")
	if val == "" {
		return "", http.ErrNoCookie
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 || vals[0] != "APIKey" {
		return "", http.ErrNoCookie
	}
	return vals[1], nil
}