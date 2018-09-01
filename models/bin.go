package models

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
)

// Bin is a designated space to partition requests
type Bin struct {
	ID      string    `bson:"_id" json:"id" jsonapi:"primary,bin"`
	Name    string    `bson:"name" json:"name" jsonapi:"attr,name"`
	Created time.Time `bson:"created" json:"created" jsonapi:"attr,created"`
}

// NewBin generates a Bin struct to use
func NewBin(r *http.Request) (*Bin, error) {

	bin := &Bin{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&bin)
	if err != nil {
		return nil, err
	}

	bin.Created = time.Now()

	return bin, nil
}

// ErrBinNotFound is returned when an bin cannot be found
var ErrBinNotFound = errors.New("Bin could not be found")

// BSON transforms a Payload to BSON for storage in MongoDB
func (b *Bin) BSON() *bson.Document {
	return bson.NewDocument(
		bson.EC.String("_id", b.ID),
		bson.EC.String("name", b.Name),
		bson.EC.Time("created", b.Created),
	)
}
