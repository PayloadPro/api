package services

import (
	"database/sql"
	"time"

	"github.com/PayloadPro/pro.payload.api/models"
)

// RequestService deals with incoming requests
type RequestService struct {
	DB *sql.DB
}

// Save an incoming request
func (s *RequestService) Save(request *models.Request) error {

	var err error

	rows, err := s.DB.Query(
		"INSERT INTO requests (bin, method, body, content_length, content_type, remote_addr, protocol, user_agent) "+
			"VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"+
			"RETURNING id, created",
		request.Bin.ID, request.Method, request.Body, request.ContentLength, request.ContentType, request.RemoteAddr, request.Proto, request.UserAgent,
	)
	if err != nil {
		return err
	}

	for rows.Next() {
		var id string
		var created time.Time
		if err = rows.Scan(&id, &created); err != nil {
			return err
		}
		request.ID = id
		request.Created = created
	}

	return nil
}

// GetRequest gets a request
func (s *RequestService) GetRequest(id string, bin *models.Bin) (*models.Request, error) {

	var err error
	request := &models.Request{}

	var (
		created     time.Time
		method      string
		body        []byte
		remote      string
		length      int64
		contentType string
		protocol    string
		ua          string
	)
	err = s.DB.QueryRow(
		"SELECT created, method, body, content_length, content_type, remote_addr, protocol, user_agent FROM requests WHERE id = $1",
		id,
	).Scan(&created, &method, &body, &length, &contentType, &remote, &protocol, &ua)

	switch {
	case err == sql.ErrNoRows:
		return nil, models.ErrRequestNotFound
	case err != nil:
		return nil, err
	default:
		request.ID = id
		request.Created = created
		request.Method = method
		request.Body = body
		request.Bin = bin
		request.ContentLength = length
		request.ContentType = contentType
		request.RemoteAddr = remote
		request.Proto = protocol
		request.UserAgent = ua
	}

	return request, nil
}

// GetRequestsForBin gets requests for a bin
func (s *RequestService) GetRequestsForBin(bin *models.Bin) ([]*models.Request, error) {

	var err error
	var requests []*models.Request

	rows, err := s.DB.Query(
		"SELECT id, created, method, body, content_length, content_type, remote_addr, protocol, user_agent "+
			"FROM requests WHERE bin = $1 ORDER BY created DESC LIMIT 100",
		bin.ID,
	)

	switch {
	case err == sql.ErrNoRows:
		return requests, nil
	case err != nil:
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			id          string
			created     time.Time
			method      string
			body        []byte
			remote      string
			length      int64
			contentType string
			protocol    string
			ua          string
		)
		if err := rows.Scan(&id, &created, &method, &body, &length, &contentType, &remote, &protocol, &ua); err != nil {
			return nil, err
		}
		requests = append(requests, &models.Request{
			ID:            id,
			Created:       created,
			Method:        method,
			Body:          body,
			ContentLength: length,
			ContentType:   contentType,
			RemoteAddr:    remote,
			Proto:         protocol,
			UserAgent:     ua,
			Bin:           bin,
		})
	}

	return requests, nil
}
