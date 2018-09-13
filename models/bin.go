package models

import (
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
	ID          string `jsonapi:"primary,bin"`
	Name        string `jsonapi:"attr,name"`
	Description string `jsonapi:"attr,description"`
	RemoteAddr  string
	Created     time.Time
	Stats       *Stats
	Config      *configs.AppConfig
}

// JSONAPILinks return links for the JSONAPI marshal
func (b Bin) JSONAPILinks() *jsonapi.Links {
	return &jsonapi.Links{
		"self":     fmt.Sprintf("%s/bins/%s", b.Config.APILink, b.ID),
		"request":  fmt.Sprintf("%s/bins/%s/request", b.Config.APILink, b.ID),
		"requests": fmt.Sprintf("%s/bins/%s/requests", b.Config.APILink, b.ID),
	}
}

// JSONAPIMeta return meta for the JSONAPI marshal
func (b Bin) JSONAPIMeta() *jsonapi.Meta {
	return &jsonapi.Meta{
		"stats": &jsonapi.Meta{
			"requests": &jsonapi.Meta{
				"total":      b.Stats.Total,
				"break_down": b.Stats.Breakdown(),
			},
		},
	}
}

// NewBin generates a Bin struct to use
func NewBin(r *http.Request, config *configs.AppConfig) (*Bin, error) {

	bin := &Bin{}

	if err := jsonapi.UnmarshalPayload(r.Body, bin); err != nil {
		return nil, err
	}

	bin.Config = config

	return bin, nil
}
