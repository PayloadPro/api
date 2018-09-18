package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/PayloadPro/api/configs"
	"github.com/PayloadPro/api/utils"
	"github.com/google/jsonapi"
)

// ErrRequestNotFound is returned when an request cannot be found
var ErrRequestNotFound = errors.New("Request could not be found")

// Request is the internal representation of a request to a bin
type Request struct {
	ID            string `jsonapi:"primary,request"`
	Bin           *Bin
	Method        string      `jsonapi:"attr,method"`
	Proto         string      `jsonapi:"attr,protocol"`
	ContentLength int64       `jsonapi:"attr,content_length"`
	ContentType   string      `jsonapi:"attr,content_type"`
	UserAgent     string      `jsonapi:"attr,user_agent"`
	RemoteAddr    string      `jsonapi:"attr,remote_addr"`
	Body          []byte      `jsonapi:"attr,body,omitempty"`
	BodyI         interface{} `jsonapi:"attr,body,omitempty"`
	Created       time.Time
	Config        *configs.AppConfig
}

// JSONAPILinks return links for the JSONAPI marshal
func (r Request) JSONAPILinks() *jsonapi.Links {
	return &jsonapi.Links{
		"self":     fmt.Sprintf("%s/bins/%s/requests/%s", r.Config.APILink, r.Bin.ID, r.ID),
		"bin":      fmt.Sprintf("%s/bins/%s", r.Config.APILink, r.Bin.ID),
		"request":  fmt.Sprintf("%s/bins/%s/request", r.Config.APILink, r.Bin.ID),
		"requests": fmt.Sprintf("%s/bins/%s/requests", r.Config.APILink, r.Bin.ID),
	}
}

// JSONAPIMeta return meta for the JSONAPI marshal
func (r Request) JSONAPIMeta() *jsonapi.Meta {
	return &jsonapi.Meta{
		"created": utils.FormatTimeMeta(r.Created),
	}
}

// ErrBodyRead is returned when an body cannot be read
var ErrBodyRead = errors.New("Could not read the request body - please check it's valid JSON")

// NewRequest generates a Request struct to use
func NewRequest(r *http.Request, bin *Bin) (*Request, error) {

	request := &Request{}

	request.Method = r.Method
	request.Proto = r.Proto
	request.ContentLength = r.ContentLength
	request.ContentType = r.Header.Get("Content-Type")
	request.UserAgent = r.UserAgent()
	request.RemoteAddr = r.RemoteAddr
	request.Bin = bin

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return request, ErrBodyRead
	}

	request.Body = b

	return request, nil
}

// PrepareBody for presentation - will take a string and unmarshal it
func (r *Request) PrepareBody() {
	var nb interface{}
	if err := json.Unmarshal([]byte(r.Body), &nb); err != nil {
		r.BodyI = nil
		return
	}
	r.BodyI = nb
}
