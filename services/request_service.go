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

// RequestService deals with incoming requests
type RequestService struct {
	Collection *mongo.Collection
}

// Save an incoming request
func (s *RequestService) Save(request *models.Request) (*models.Request, error) {

	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	request.ID = id.String()

	_, err = s.Collection.InsertOne(nil, request)

	if err != nil {
		return nil, err
	}

	return request, nil
}

// GetRequestsForBin gets requests for a bin
func (s *RequestService) GetRequestsForBin(id string) ([]*models.Request, error) {

	sort := findopt.Sort(bson.NewDocument(bson.EC.Int32("created", -1)))
	limit := findopt.Limit(100)

	cur, err := s.Collection.Find(nil, bson.NewDocument(
		bson.EC.String("bin", id),
	), sort, limit)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var requests []*models.Request

	for cur.Next(nil) {
		request := models.Request{}
		err := cur.Decode(&request)
		if err != nil {
			log.Fatal("Decode error ", err)
		}
		requests = append(requests, &request)
	}

	if err := cur.Err(); err != nil {
		log.Fatal("Cursor error ", err)
	}

	return requests, nil
}
