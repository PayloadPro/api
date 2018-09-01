package services

import (
	"context"
	"log"

	"github.com/PayloadPro/pro.payload.api/models"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/findopt"
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

	_, err = s.Collection.InsertOne(nil, bin)

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

// GetBins gets bins sorted by created date
func (s *BinService) GetBins() ([]models.Bin, error) {

	sort := findopt.Sort(bson.NewDocument(bson.EC.Int32("created", -1)))
	limit := findopt.Limit(100)

	cur, err := s.Collection.Find(nil, nil, sort, limit)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var bins []models.Bin

	for cur.Next(nil) {
		bin := models.Bin{}
		err := cur.Decode(&bin)
		if err != nil {
			log.Fatal("Decode error ", err)
		}
		bins = append(bins, bin)
	}

	if err := cur.Err(); err != nil {
		log.Fatal("Cursor error ", err)
	}

	return bins, nil
}
