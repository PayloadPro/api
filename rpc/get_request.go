package rpc

import (
	"net/http"

	"golang.org/x/net/context"

	"github.com/PayloadPro/pro.payload.api/deps"
	"github.com/PayloadPro/pro.payload.api/models"
	"github.com/gorilla/mux"
)

// GetRequest is a func which takes the incoming request for a bin request and returns the request
type GetRequest func(context.Context, *http.Request) (*models.Request, int, error)

// NewGetRequestForBin is the concrete func for GetRequest
func NewGetRequestForBin(services *deps.Services, config *deps.Config) GetRequest {
	return func(ctx context.Context, r *http.Request) (*models.Request, int, error) {

		// get the bin from the DB based on ID in the URL to check it exists
		vars := mux.Vars(r)
		bin, err := services.Bin.GetByID(vars["id"])
		bin.Config = config.App

		if err != nil {
			return nil, http.StatusNotFound, models.ErrBinNotFound
		}

		// bin exists, get the request
		request, err := services.Request.GetRequest(vars["request_id"])
		if request.Bin != bin.ID {
			return nil, http.StatusNotFound, models.ErrRequestNotFound
		}

		request.Config = config.App

		return request, http.StatusOK, nil
	}
}
