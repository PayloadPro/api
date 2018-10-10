package configs

import (
	"os"
	"testing"
)

func TestCockroachSetup(t *testing.T) {

	DSN := "postgresql://user@host:port/table"

	os.Setenv("DB_DSN", DSN)

	db := &CockroachConfig{}
	db.Setup()

	if db.DSN != DSN {
		t.Errorf("Incorrect DSN. Got: `%s`, expected: `%s`", db.DSN, DSN)
	}
}
