package services

import (
	"github.com/andrew-waters/pro.payload.api/models"
	"github.com/mongodb/mongo-go-driver/bson"
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

// GetByID gets a bin by ID
func (s *BinService) GetByID(id string) (*models.Bin, error) {

	bin := &models.Bin{}
	result := s.Collection.FindOne(nil, bson.NewDocument(bson.EC.String("_id", id)))
	result.Decode(bin)

	if bin.ID == "" {
		return nil, models.ErrBinNotFound
	}

	return bin, nil
}
