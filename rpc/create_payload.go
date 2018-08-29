package rpc

import (
	"net/http"

	"github.com/gorilla/mux"

	"golang.org/x/net/context"

	"github.com/andrew-waters/pro.payload.api/deps"
	"github.com/andrew-waters/pro.payload.api/models"
)

// CreatePayloadResponse is the response from the CreatePayload endpoint
type CreatePayloadResponse struct {
	Status string `json:"status"`
	ID     string `json:"id"`
	Bin    string `json:"bin"`
}

// CreatePayload is a func which takes the incoming request, saves it persistently
// and returns the CreatePayloadResponse for the consumer
type CreatePayload func(context.Context, *http.Request) (*CreatePayloadResponse, int, error)

// NewCreatePayload is the concrete func for CreatePayload
func NewCreatePayload(services *deps.Services) CreatePayload {
	return func(ctx context.Context, r *http.Request) (*CreatePayloadResponse, int, error) {

		// create the payload
		var payload *models.Payload
		var err error

		if payload, err = models.NewPayload(r); err != nil {
			return nil, http.StatusInternalServerError, err
		}

		// get the bin from the DB based on ID in the URL
		vars := mux.Vars(r)
		bin, err := services.Bin.GetByID(vars["id"])

		if err != nil {
			return nil, http.StatusNotFound, err
		}

		payload.Bin = bin

		// save the payload
		if err = services.Payload.Save(payload); err != nil {
			return nil, http.StatusInternalServerError, err
		}

		return &CreatePayloadResponse{
			Status: SUCCESS,
			ID:     payload.ID,
			Bin:    payload.Bin.ID,
		}, http.StatusCreated, nil
	}
}
