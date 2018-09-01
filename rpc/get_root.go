package rpc

import (
	"net/http"

	"github.com/PayloadPro/pro.payload.api/deps"

	"github.com/PayloadPro/pro.payload.api/models"
	"golang.org/x/net/context"
)

// GetRoot is a func which returns the Root model
type GetRoot func(context.Context, *http.Request) (*models.Root, int, error)

// NewGetRoot is the concrete func for Root
func NewGetRoot(services *deps.Services, config *deps.Config) GetRoot {
	return func(ctx context.Context, r *http.Request) (*models.Root, int, error) {

		root, err := services.Root.Get(config.App)
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}

		return root, http.StatusOK, nil
	}
}
