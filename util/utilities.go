package util

import (
	"encoding/json"
	"net/http"
)

func RenderJson(w http.ResponseWriter, val any, statusCode int) {
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(val)
}
