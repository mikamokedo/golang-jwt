package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
)

var Validate = validator.New()


func ParseJsonBody(r *http.Request, v interface{}) error {
	if r.Body == nil {
		return fmt.Errorf("request body is empty")
	}
	return json.NewDecoder(r.Body).Decode(v)
}



func WriteJsonResponse(w http.ResponseWriter,status int, v interface {}) error{
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}


func WriteJsonError(w http.ResponseWriter, status int, error error) error {
	return WriteJsonResponse(w, status, map[string]string{"error": error.Error()})
}