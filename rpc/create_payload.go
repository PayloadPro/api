package rpc

import (
	"net/http"

	"github.com/gorilla/mux"

	"golang.org/x/net/context"

	"github.com/andrew-waters/payload.pro/deps"
	"github.com/andrew-waters/payload.pro/models"
)

const (
	SUCCESS = "success"
)

type CreatePayloadResponse struct {
	Message string `json:"status"`
	ID      string `json:"id"`
	Bin     string `json:"bin"`
}

type CreatePayload func(context.Context, *http.Request) (*CreatePayloadResponse, int, error)

func NewCreatePayload(services *deps.Services) CreatePayload {
	return func(ctx context.Context, r *http.Request) (*CreatePayloadResponse, int, error) {

		// create the payload
		var payload *models.Payload
		var err error

		if payload, err = models.NewPayload(r); err != nil {
			return nil, http.StatusInternalServerError, err
		}

		// todo - get the bin from the DB based on ID
		vars := mux.Vars(r)
		payload.Bin = &models.Bin{ID: vars["id"]}

		// save the payload
		if err = services.Payload.Save(payload); err != nil {
			return nil, http.StatusInternalServerError, err
		}

		return &CreatePayloadResponse{
			Message: SUCCESS,
			ID:      payload.ID,
			Bin:     payload.Bin.ID,
		}, http.StatusCreated, nil
	}
}
