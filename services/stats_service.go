package services

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/PayloadPro/api/models"
)

// StatsService deals with incoming requests
type StatsService struct {
	DB *sql.DB
}

// GetStatsForBin gets the stats for a single bin
func (s *StatsService) GetStatsForBin(bin *models.Bin) error {
	var err error
	stats := &models.Stats{}

	var (
		total   int64
		get     int64
		post    int64
		put     int64
		patch   int64
		options int64
		head    int64
		delete  int64
	)
	err = s.DB.QueryRow(
		"SELECT total, get, post, put, patch, options, head, delete FROM stats WHERE bin = $1",
		bin.ID,
	).Scan(&total, &get, &post, &put, &patch, &options, &head, &delete)

	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}
	}
	stats.Total = total
	stats.GET = get
	stats.POST = post
	stats.PUT = put
	stats.PATCH = patch
	stats.OPTIONS = options
	stats.HEAD = head
	stats.DELETE = delete

	bin.Stats = stats

	return nil
}

// GetStatsForBins gets the stats for muliple bins
func (s *StatsService) GetStatsForBins(bins []*models.Bin) error {

	var err error

	var ids []string
	for _, bin := range bins {
		ids = append(ids, bin.ID)
	}

	stmt := fmt.Sprintf(`SELECT bin, total, get, post, put, patch, options, head, delete FROM stats WHERE bin IN (%s)`, "'"+strings.Join(ids, "', '")+"'")
	rows, err := s.DB.Query(stmt)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			bin     string
			total   int64
			get     int64
			post    int64
			put     int64
			patch   int64
			options int64
			head    int64
			delete  int64
		)
		if err = rows.Scan(&bin, &total, &get, &post, &put, &patch, &options, &head, &delete); err != nil {
			return err
		}
		for _, mb := range bins {
			if mb.ID == bin {
				mb.Stats = &models.Stats{
					Total:   total,
					GET:     get,
					POST:    post,
					PUT:     put,
					PATCH:   patch,
					OPTIONS: options,
					HEAD:    head,
					DELETE:  delete,
				}
			}
		}
	}

	for _, bin := range bins {
		if bin.Stats == nil {
			bin.Stats = &models.Stats{}
		}
	}

	return err
}

// AddRequest updates the stats for a bin
func (s *StatsService) AddRequest(request *models.Request) error {

	var err error
	stats := &models.Stats{}

	var (
		total   int64
		get     int64
		post    int64
		put     int64
		patch   int64
		options int64
		head    int64
		delete  int64
	)
	err = s.DB.QueryRow(
		"SELECT total, get, post, put, patch, options, head, delete FROM stats WHERE bin = $1",
		request.Bin.ID,
	).Scan(&total, &get, &post, &put, &patch, &options, &head, &delete)

	insert := false

	switch {
	case err == sql.ErrNoRows:
		insert = true
	case err != nil:
		return err
	default:
		stats.Total = total
		stats.GET = get
		stats.POST = post
		stats.PUT = put
		stats.PATCH = patch
		stats.OPTIONS = options
		stats.HEAD = head
		stats.DELETE = delete
	}

	stats.Bin = request.Bin

	// add one more request
	stats.Total++

	// add the request type
	switch request.Method {
	case "GET":
		stats.GET++
	case "POST":
		stats.POST++
	case "PUT":
		stats.PUT++
	case "PATCH":
		stats.PATCH++
	case "OPTIONS":
		stats.OPTIONS++
	case "HEAD":
		stats.HEAD++
	case "DELETE":
		stats.DELETE++
	}

	if insert {
		_, err := s.DB.Query(
			"INSERT INTO stats (bin, total, get, post, put, patch, options, head, delete) "+
				"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
			stats.Bin.ID, stats.Total, stats.GET, stats.POST, stats.PUT, stats.PATCH, stats.OPTIONS, stats.HEAD, stats.DELETE,
		)
		if err != nil {
			return err
		}
	} else {
		_, err := s.DB.Exec(
			"UPDATE stats SET total = $2, get = $3, post = $4, put = $5, patch = $6, options = $7, head = $8, delete = $9 WHERE bin = $1",
			stats.Bin.ID, stats.Total, stats.GET, stats.POST, stats.PUT, stats.PATCH, stats.OPTIONS, stats.HEAD, stats.DELETE,
		)
		if err != nil {
			return err
		}
	}

	return nil
}
