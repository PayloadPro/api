package configs

import (
	"fmt"
	"os"
	"strconv"
)

// DatabaseConfig is the config for the mongo database
type DatabaseConfig struct {
	Host                 string
	Port                 int
	BinDatabase          string
	BinCollection        string
	BinRequestCollection string
}

// Setup the DatabaseConfig
func (dc *DatabaseConfig) Setup() {
	dc.Host = os.Getenv("DB_HOST")
	dc.Port, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	dc.BinDatabase = os.Getenv("DB_BIN_DATABASE")
	dc.BinCollection = os.Getenv("DB_BIN_COLLECTION")
	dc.BinRequestCollection = os.Getenv("DB_BIN_REQUEST_COLLECTION")
}

// ConnectionString generates the connection string fo rth emonog driver
func (dc *DatabaseConfig) ConnectionString() string {
	return fmt.Sprintf("mongodb://%s:%d", dc.Host, dc.Port)
}
