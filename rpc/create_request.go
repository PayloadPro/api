package rpc

import (
	"net/http"

	"github.com/gorilla/mux"

	"golang.org/x/net/context"

	"github.com/PayloadPro/pro.payload.api/deps"
	"github.com/PayloadPro/pro.payload.api/models"
)

// CreateRequest is a func which takes the incoming request, saves it persistently
// and returns the CreateRequestResponse for the consumer
type CreateRequest func(context.Context, *http.Request) (*models.Request, int, error)

// NewCreateRequest is the concrete func for CreateRequest
func NewCreateRequest(services *deps.Services) CreateRequest {
	return func(ctx context.Context, r *http.Request) (*models.Request, int, error) {

		// create the request
		var request *models.Request
		var err error

		if request, err = models.NewRequest(r); err != nil {
			return nil, http.StatusInternalServerError, err
		}

		// get the bin from the DB based on ID in the URL
		vars := mux.Vars(r)
		bin, err := services.Bin.GetByID(vars["id"])

		if err != nil {
			return nil, http.StatusNotFound, err
		}

		request.Bin = bin

		// save the payload
		if err = services.Request.Save(request); err != nil {
			return nil, http.StatusInternalServerError, err
		}

		return request, http.StatusCreated, nil
	}
}
