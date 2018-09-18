package services

import (
	"github.com/PayloadPro/api/configs"
	"github.com/PayloadPro/api/models"
)

// RootService deals with root requests (/)
type RootService struct {
}

// Get the root
func (s *RootService) Get(config *configs.AppConfig) (*models.Root, error) {
	return models.NewRoot(config)
}
