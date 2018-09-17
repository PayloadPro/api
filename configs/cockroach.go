package configs

import (
	"os"
)

// CockroachConfig is the config for the cockroach database
type CockroachConfig struct {
	DSN string
}

// Setup the DatabaseConfig
func (cc *CockroachConfig) Setup() {
	cc.DSN = os.Getenv("DB_DSN")
}
