package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// Envelope allowing flexible JSON serialization by enabling different types of data to be passed
type Envelope map[string]interface {}

// Standardise JSON response writing
func WriteJSON(w http.ResponseWriter, status int, data Envelope) error {
	js, err := json.MarshalIndent(data, "", "	") // convert data to JSON with specified indentation
	if err != nil {
		return err
	}

	js = append(js, '\n')
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

// Extract and convert ID parameters from HTTP requests
func ReadIDParam(r *http.Request) (int64, error) {
	idParam := chi.URLParam(r, "id") // retrieve id parameter
	if idParam == "" {
		return 0, errors.New("Invalid id parameter")
	}
	id, err := strconv.ParseInt(idParam, 10, 64) // convert from a string to int64
	if err != nil {
		return 0, errors.New("Invalid id parameter type")
	}
	return id, nil
}