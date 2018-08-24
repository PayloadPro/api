package models

import (
	"errors"
	"io/ioutil"
	"net/http"
)

// Payload is the internal reqpresentation of a request
type Payload struct {
	ID            string
	Bin           *Bin
	Method        string
	Proto         string
	ContentLength int64
	Host          string
	UserAgent     string
	RemoteAddr    string
	Body          []byte
}

// ErrBodyRead is returned when an body cannot be read
var ErrBodyRead = errors.New("could not read body")

// NewPayload generates a Payload struct to use
func NewPayload(r *http.Request) (*Payload, error) {

	payload := &Payload{}

	payload.Method = r.Method
	payload.Proto = r.Proto
	payload.ContentLength = r.ContentLength
	payload.Host = r.Host
	payload.UserAgent = r.UserAgent()
	payload.RemoteAddr = r.RemoteAddr

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return payload, ErrBodyRead
	}

	payload.Body = b

	return payload, nil
}
