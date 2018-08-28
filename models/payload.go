package models

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
)

// Payload is the internal reqpresentation of a request
type Payload struct {
	ID            string
	Bin           *Bin
	Method        string
	Proto         string
	ContentLength int64
	UserAgent     string
	RemoteAddr    string
	Body          []byte
	Created       time.Time
}

// ErrBodyRead is returned when an body cannot be read
var ErrBodyRead = errors.New("could not read body")

// NewPayload generates a Payload struct to use
func NewPayload(r *http.Request) (*Payload, error) {

	payload := &Payload{}

	payload.Method = r.Method
	payload.Proto = r.Proto
	payload.ContentLength = r.ContentLength
	payload.UserAgent = r.UserAgent()
	payload.RemoteAddr = r.RemoteAddr

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return payload, ErrBodyRead
	}

	payload.Body = b
	payload.Created = time.Now()

	return payload, nil
}

// BSON transforms a Payload to BSON for storage in MongoDB
func (p *Payload) BSON() *bson.Document {
	return bson.NewDocument(
		bson.EC.String("id", p.ID),
		bson.EC.String("bin", p.Bin.ID),
		bson.EC.String("method", p.Method),
		bson.EC.String("proto", p.Proto),
		bson.EC.Int64("length", p.ContentLength),
		bson.EC.String("ua", p.UserAgent),
		bson.EC.String("remote", p.RemoteAddr),
		bson.EC.String("body", string(p.Body)),
		bson.EC.Time("created", p.Created),
	)
}
