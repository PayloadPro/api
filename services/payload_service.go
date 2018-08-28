package services

import (
	"github.com/andrew-waters/payload.pro/models"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/satori/go.uuid"
)

// PayloadService deals with incoming requests
type PayloadService struct {
	Collection *mongo.Collection
}

// Save an incoming request
func (s *PayloadService) Save(payload *models.Payload) error {

	id, err := uuid.NewV4()
	if err != nil {
		return err
	}
	payload.ID = id.String()

	_, err = s.Collection.InsertOne(nil, payload.BSON())

	if err != nil {
		return err
	}

	return nil
}
