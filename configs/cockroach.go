package configs

import (
	"fmt"
	"os"
	"strconv"
)

// CockroachConfig is the config for the cockroach database
type CockroachConfig struct {
	User     string
	Host     string
	Port     int
	Database string
}

// Setup the DatabaseConfig
func (cc *CockroachConfig) Setup() {
	cc.User = os.Getenv("DB_USER")
	cc.Host = os.Getenv("DB_HOST")
	cc.Port, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	cc.Database = os.Getenv("DB_DATABASE")
}

// ConnectionString generates the connection string fo the postgres/cockroach driver
func (cc *CockroachConfig) ConnectionString() string {
	return fmt.Sprintf("postgresql://%s@%s:%d/%s?sslmode=disable", cc.User, cc.Host, cc.Port, cc.Database)
}
