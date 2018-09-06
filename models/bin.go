package models

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/PayloadPro/pro.payload.api/configs"
	"github.com/PayloadPro/pro.payload.api/utils"
	"github.com/google/jsonapi"
)

// ErrBinNotFound is returned when an bin cannot be found
var ErrBinNotFound = errors.New("Bin could not be found")

// Bin is a designated space to partition requests
type Bin struct {
	ID      string             `bson:"_id" jsonapi:"primary,bin"`
	Name    string             `bson:"name" jsonapi:"attr,name"`
	Created time.Time          `bson:"created"`
	Stats   Stats              `bson:"stats"`
	Config  *configs.AppConfig `bson:"-"`
}

// Stats contains cached stats for the bin
type Stats struct {
	Requests RequestStats `bson:"requests"`
}

// RequestStats contains stats for the requests made to the bin
type RequestStats struct {
	Total int64 `bson:"total"`
	// Breakdown contains Request Method breakdown figures eg Breakdown["GET"]=123
	Breakdown map[string]int64 `bson:"breakdown"`
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
		"created": utils.FormatTimeMeta(b.Created),
		"stats": &jsonapi.Meta{
			"requests": &jsonapi.Meta{
				"total":      b.Stats.Requests.Total,
				"break_down": b.Stats.Requests.Breakdown,
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
	bin.Created = time.Now()
	bin.Config = config

	return bin, nil
}

// AddRequest is executed when a request is successfully and we update the headline stats for the bin
func (b *Bin) AddRequest(r *Request) {
	b.Stats.Requests.Total++
	if b.Stats.Requests.Breakdown == nil {
		b.Stats.Requests.Breakdown = make(map[string]int64, 0)
	}
	if _, ok := b.Stats.Requests.Breakdown[r.Method]; !ok {
		b.Stats.Requests.Breakdown[r.Method] = 0
	}
	b.Stats.Requests.Breakdown[r.Method]++
}
