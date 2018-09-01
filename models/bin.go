package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/PayloadPro/pro.payload.api/configs"
	"github.com/google/jsonapi"
)

// ErrBinNotFound is returned when an bin cannot be found
var ErrBinNotFound = errors.New("Bin could not be found")

// Bin is a designated space to partition requests
type Bin struct {
	ID      string             `bson:"_id" jsonapi:"primary,bin"`
	Name    string             `bson:"name" jsonapi:"attr,name"`
	Created time.Time          `bson:"created"`
	Config  *configs.AppConfig `bson:"-"`
}

// JSONAPILinks return links for the JSONAPI marshal
func (b Bin) JSONAPILinks() *jsonapi.Links {
	return &jsonapi.Links{
		"self": fmt.Sprintf("%s/bins/%s", b.Config.APILink, b.ID),
	}
}

// JSONAPIMeta return meta for the JSONAPI marshal
func (b Bin) JSONAPIMeta() *jsonapi.Meta {
	return &jsonapi.Meta{
		"created": b.Created,
	}
}

// NewBin generates a Bin struct to use
func NewBin(r *http.Request, config *configs.AppConfig) (*Bin, error) {

	bin := &Bin{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&bin)
	if err != nil {
		return nil, err
	}

	bin.Created = time.Now()
	bin.Config = config

	return bin, nil
}
