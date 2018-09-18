package services

import (
	"database/sql"
	"time"

	"github.com/PayloadPro/api/models"
)

// BinService deals with incoming requests
type BinService struct {
	DB *sql.DB
}

// Save an incoming request
func (s *BinService) Save(bin *models.Bin) error {

	var err error

	rows, err := s.DB.Query(
		"INSERT INTO bins (name, description, remote_addr) "+
			"VALUES ($1, $2, $3)"+
			"RETURNING id, created",
		bin.Name, bin.Description, bin.RemoteAddr,
	)
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var id string
		var created time.Time
		if err = rows.Scan(&id, &created); err != nil {
			return err
		}
		bin.ID = id
		bin.Created = created
	}

	return nil
}

// GetByID gets a bin by ID
func (s *BinService) GetByID(id string) (*models.Bin, error) {

	var err error

	bin := &models.Bin{}

	var created time.Time
	var name string
	var description string
	var remote string
	err = s.DB.QueryRow(
		"SELECT created, name, description, remote_addr FROM bins WHERE id = $1",
		id,
	).Scan(&created, &name, &description, &remote)

	switch {
	case err == sql.ErrNoRows:
		return nil, models.ErrBinNotFound
	case err != nil:
		return nil, err
	default:
		bin.ID = id
		bin.Name = name
		bin.Description = description
		bin.RemoteAddr = remote
		bin.Created = created
	}

	return bin, nil
}

// GetBins gets bins sorted by created date
func (s *BinService) GetBins() ([]*models.Bin, error) {

	var err error
	var bins []*models.Bin

	rows, err := s.DB.Query("SELECT id, created, name, description, remote_addr FROM bins ORDER BY created DESC LIMIT 100")

	switch {
	case err == sql.ErrNoRows:
		return bins, nil
	case err != nil:
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			id          string
			created     time.Time
			name        string
			description string
			remote      string
		)
		if err := rows.Scan(&id, &created, &name, &description, &remote); err != nil {
			return nil, err
		}
		bins = append(bins, &models.Bin{
			ID:          id,
			Name:        name,
			Description: description,
			Created:     created,
			RemoteAddr:  remote,
		})
	}

	return bins, nil
}
