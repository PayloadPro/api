package models

// Stats is a designated struct for headline bin statistics
type Stats struct {
	Bin     *Bin
	Total   int64
	GET     int64
	POST    int64
	PUT     int64
	PATCH   int64
	OPTIONS int64
	HEAD    int64
	DELETE  int64
}

// Breakdown provides a user friendly breakdown of request count by request type
func (s Stats) Breakdown() map[string]interface{} {
	b := make(map[string]interface{}, 0)
	b["GET"] = s.GET
	b["POST"] = s.POST
	b["PUT"] = s.PUT
	b["PATCH"] = s.PATCH
	b["OPTIONS"] = s.OPTIONS
	b["HEAD"] = s.HEAD
	b["DELETE"] = s.DELETE
	return b
}
