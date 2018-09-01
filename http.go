package main

import (
	"net/http"
	"strconv"

	"github.com/google/jsonapi"
)

const (
	// InternalServerError insidicates we couldn't handle something gracefully for the request
	InternalServerError = `{"errors": [{"message": "Internal server error"}]}`
)

// HTTPErrorWrapper wraps an error for JSON output
type HTTPErrorWrapper struct {
	Errors []HTTPError `jsonapi:"errors"`
}

// HTTPError is the body of the error with a message inside from error
type HTTPError struct {
	Message string `json:"message"`
}

type callback func() (interface{}, int, error)

// JSONEndpointHandler handles API endpoints in JSON
func JSONEndpointHandler(w http.ResponseWriter, r *http.Request, cb callback) error {

	w.Header().Set("Content-Type", jsonapi.MediaType)

	var err error
	var status int

	var resp interface{}

	// the callback failed, write the error and return
	if resp, status, err = cb(); err != nil {

		w.WriteHeader(status)
		jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
			Title:  "Processing Error",
			Detail: err.Error(),
			Status: strconv.Itoa(status),
		}})

		return nil
	}

	w.WriteHeader(status)

	if err := jsonapi.MarshalPayload(w, resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return nil
}
