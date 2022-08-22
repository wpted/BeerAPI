package util

import (
	"encoding/json"
	"net/http"
)

// RenderJson encodes and writes the given jsonified value to the w
// , simultaneously setting the status code to the header
func RenderJson(w http.ResponseWriter, val any, statusCode int) {
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(val)
}
