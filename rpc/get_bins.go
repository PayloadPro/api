package rpc

import (
	"net/http"

	"golang.org/x/net/context"

	"github.com/andrew-waters/pro.payload.api/deps"
	"github.com/andrew-waters/pro.payload.api/models"
)

// GetBinsResponse is the response from the GetBins endpoint
type GetBinsResponse struct {
	Bins []models.Bin `json:"bins"`
}

// GetBins is a func which takes the incoming request for the bins and returns the GetBinsResponse
type GetBins func(context.Context, *http.Request) (*GetBinsResponse, int, error)

// NewGetBins is the concrete func for GetBins
func NewGetBins(services *deps.Services) GetBins {
	return func(ctx context.Context, r *http.Request) (*GetBinsResponse, int, error) {

		var bins = make([]models.Bin, 0)
		var err error

		if bins, err = services.Bin.GetBins(); err != nil {
			return nil, http.StatusInternalServerError, err
		}

		return &GetBinsResponse{
			Bins: bins,
		}, http.StatusOK, nil
	}
}
