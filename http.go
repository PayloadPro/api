package main

import (
	"encoding/json"
	"io"
	"net/http"
)

const (
	// InternalServerError insidicates we couldn't handle something gracefully for the request
	InternalServerError = `{"error": {"message": "Internal server error"}}`
)

// HTTPErrorWrapper wraps an error for JSON output
type HTTPErrorWrapper struct {
	Error HTTPError `json:"error"`
}

// HTTPError is the body of the error with a message inside from error
type HTTPError struct {
	Message error `json:"message"`
}

type callback func() (interface{}, int, error)

// JSONEndpointHandler handles API endpoints in JSON
func JSONEndpointHandler(w http.ResponseWriter, r *http.Request, cb callback) error {

	var err error
	if w.Header().Get("Content-Type") == "" {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	}

	var resp interface{}
	var nativeErr []byte
	var status int

	// the callback failed, wrap the error and return
	if resp, status, err = cb(); err != nil {
		e := HTTPErrorWrapper{
			Error: HTTPError{err},
		}

		if nativeErr, err = json.Marshal(e); err != nil {
			// super failure here - we couldn't marshal the error we were sending
			// so send a plain internal server error
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, InternalServerError)
			return err
		}

		w.WriteHeader(status)
		io.WriteString(w, string(nativeErr))
		return err
	}

	// we couldn't encode the response
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, InternalServerError)
		return err
	}

	w.WriteHeader(status)

	return err
}
