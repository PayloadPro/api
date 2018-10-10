package configs

import (
	"os"
	"testing"

	. "github.com/franela/goblin"
)

func TestCockroachSetup(t *testing.T) {

	g := Goblin(t)

	DSN := "postgresql://user@host:port/table"

	os.Setenv("DB_DSN", DSN)

	db := &CockroachConfig{}
	db.Setup()

	g.Describe("Returns OS env vars as values", func() {

		g.It("Contains the correct DSN", func() {
			g.Assert(db.DSN).Equal(DSN)
		})

	})

}
