package rpc

import (
	"net/http"

	"golang.org/x/net/context"
)

// LandingPayloadResponse is the response from the LandingPayload endpoint
type LandingPayloadResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// LandingPayload is a func which takes the incoming request, saves it persistently
// and returns the CreatePayloadResponse for the consumer
type LandingPayload func(context.Context, *http.Request) (*LandingPayloadResponse, int, error)

// NewLandingPayload is the concrete func for LandingPayload
func NewLandingPayload() LandingPayload {
	return func(ctx context.Context, r *http.Request) (*LandingPayloadResponse, int, error) {

		return &LandingPayloadResponse{
			Status:  SUCCESS,
			Message: "Welcome to the payload.pro API",
		}, http.StatusOK, nil
	}
}
