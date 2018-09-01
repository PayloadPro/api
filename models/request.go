package models

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
)

// Request is the internal representation of a request to a bin
type Request struct {
	ID            string    `bson:"_id" json:"id" jsonapi:"primary,request"`
	Bin           *Bin      `json:"bin" jsonapi:"relation,bin"`
	Method        string    `json:"method" jsonapi:"attr,method"`
	Proto         string    `json:"protocol" jsonapi:"attr,protocol"`
	ContentLength int64     `json:"content_length" jsonapi:"attr,content_length"`
	UserAgent     string    `json:"user_agent" jsonapi:"attr,user_agent"`
	RemoteAddr    string    `json:"remote_addr" jsonapi:"attr,remote_addr"`
	Body          string    `json:"body" jsonapi:"attr,body"`
	Created       time.Time `json:"created" jsonapi:"attr,created"`
}

// ErrBodyRead is returned when an body cannot be read
var ErrBodyRead = errors.New("could not read body")

// NewRequest generates a Request struct to use
func NewRequest(r *http.Request) (*Request, error) {

	request := &Request{}

	request.Method = r.Method
	request.Proto = r.Proto
	request.ContentLength = r.ContentLength
	request.UserAgent = r.UserAgent()
	request.RemoteAddr = r.RemoteAddr

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return request, ErrBodyRead
	}

	request.Body = string(b)
	request.Created = time.Now()

	return request, nil
}

// BSON transforms a Payload to BSON for storage in MongoDB
func (r *Request) BSON() *bson.Document {
	return bson.NewDocument(
		bson.EC.String("_id", r.ID),
		bson.EC.String("bin", r.Bin.ID),
		bson.EC.String("method", r.Method),
		bson.EC.String("proto", r.Proto),
		bson.EC.Int64("length", r.ContentLength),
		bson.EC.String("ua", r.UserAgent),
		bson.EC.String("remote", r.RemoteAddr),
		bson.EC.String("body", r.Body),
		bson.EC.Time("created", r.Created),
	)
}
