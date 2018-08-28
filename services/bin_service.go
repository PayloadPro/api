package services

import (
	"github.com/andrew-waters/pro.payload.api/models"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/satori/go.uuid"
)

// BinService deals with incoming requests
type BinService struct {
	Collection *mongo.Collection
}

// Save an incoming request
func (s *BinService) Save(bin *models.Bin) error {

	id, err := uuid.NewV4()
	if err != nil {
		return err
	}
	bin.ID = id.String()

	_, err = s.Collection.InsertOne(nil, bin.BSON())

	if err != nil {
		return err
	}

	return nil
}
