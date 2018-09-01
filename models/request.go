package models

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/PayloadPro/pro.payload.api/configs"
	"github.com/google/jsonapi"
)

// ErrRequestNotFound is returned when an request cannot be found
var ErrRequestNotFound = errors.New("Request could not be found")

// Request is the internal representation of a request to a bin
type Request struct {
	ID            string             `bson:"_id" jsonapi:"primary,request"`
	Bin           string             `bson:"bin"`
	Method        string             `bson:"method" jsonapi:"attr,method"`
	Proto         string             `bson:"protocol" jsonapi:"attr,protocol"`
	ContentLength int64              `bson:"content_length" jsonapi:"attr,content_length"`
	UserAgent     string             `bson:"user_agent" jsonapi:"attr,user_agent"`
	RemoteAddr    string             `bson:"remote_addr" jsonapi:"attr,remote_addr"`
	Body          string             `bson:"body" jsonapi:"attr,body"`
	Created       time.Time          `bson:"created"`
	Config        *configs.AppConfig `bson:"-"`
}

// JSONAPILinks return links for the JSONAPI marshal
func (r Request) JSONAPILinks() *jsonapi.Links {
	return &jsonapi.Links{
		"self": fmt.Sprintf("%s/bins/%s/requests/%s", r.Config.APILink, r.Bin, r.ID),
		"bin":  fmt.Sprintf("%s/bins/%s", r.Config.APILink, r.Bin),
	}
}

// JSONAPIMeta return meta for the JSONAPI marshal
func (r Request) JSONAPIMeta() *jsonapi.Meta {
	return &jsonapi.Meta{
		"created": r.Created,
	}
}

// ErrBodyRead is returned when an body cannot be read
var ErrBodyRead = errors.New("could not read body")

// NewRequest generates a Request struct to use
func NewRequest(r *http.Request, bin *Bin) (*Request, error) {

	request := &Request{}

	request.Method = r.Method
	request.Proto = r.Proto
	request.ContentLength = r.ContentLength
	request.UserAgent = r.UserAgent()
	request.RemoteAddr = r.RemoteAddr
	request.Bin = bin.ID

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return request, ErrBodyRead
	}

	request.Body = string(b)
	request.Created = time.Now()

	return request, nil
}
