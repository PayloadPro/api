package rpc

import (
	"net/http"

	"golang.org/x/net/context"

	"github.com/andrew-waters/pro.payload.api/deps"
	"github.com/andrew-waters/pro.payload.api/models"
)

// CreateBinResponse is the response from the CreateBin endpoint
type CreateBinResponse struct {
	Status string `json:"status"`
	ID     string `json:"id"`
}

// CreateBin is a func which takes the incoming request, saves it persistently
// and returns the CreateBinResponse for the consumer
type CreateBin func(context.Context, *http.Request) (*CreateBinResponse, int, error)

// NewCreateBin is the concrete func for CreateBin
func NewCreateBin(services *deps.Services) CreateBin {
	return func(ctx context.Context, r *http.Request) (*CreateBinResponse, int, error) {

		// create the payload
		var bin *models.Bin
		var err error

		if bin, err = models.NewBin(r); err != nil {
			return nil, http.StatusInternalServerError, err
		}

		// save the bin
		if err = services.Bin.Save(bin); err != nil {
			return nil, http.StatusInternalServerError, err
		}

		return &CreateBinResponse{
			Status: SUCCESS,
			ID:     bin.ID,
		}, http.StatusCreated, nil
	}
}
