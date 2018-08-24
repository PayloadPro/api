package services

import (
	"github.com/andrew-waters/payload.pro/models"
	"github.com/satori/go.uuid"
)

// PayloadService deals with incoming requests
type PayloadService struct{}

// Save an incoming request
func (s *PayloadService) Save(payload *models.Payload) error {

	id, err := uuid.NewV4()
	if err != nil {
		return err
	}
	payload.ID = id.String()

	// todo - perform the save

	return nil
}
