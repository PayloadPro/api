package deps

import (
	"os"
	"testing"

	"github.com/PayloadPro/api/configs"
	. "github.com/franela/goblin"
)

func TestCockroachSetup(t *testing.T) {

	g := Goblin(t)

	name := "Payload Pro"
	APILink := "https://api.payload.pro"
	DocsLink := "https://docs.payload.pro"
	SiteLink := "https://payload.pro"
	DSN := "postgresql://user@host:port/table"

	os.Setenv("APP_NAME", name)
	os.Setenv("APP_API_LINK", APILink)
	os.Setenv("APP_DOCS_LINK", DocsLink)
	os.Setenv("APP_SITE_LINK", SiteLink)
	os.Setenv("DB_DSN", DSN)

	c := &Config{
		App: &configs.AppConfig{},
		DB:  &configs.CockroachConfig{},
	}
	c.Setup()

	g.Describe("Sets up App and DB configs correctly", func() {

		g.It("Contains the correct DB DSN", func() {
			g.Assert(c.DB.DSN).Equal(DSN)
		})

		g.It("Contains the correct app name", func() {
			g.Assert(c.App.Name).Equal(name)
		})

		g.It("Contains the correct app API link", func() {
			g.Assert(c.App.APILink).Equal(APILink)
		})

		g.It("Contains the correct app docs link", func() {
			g.Assert(c.App.DocsLink).Equal(DocsLink)
		})

		g.It("Contains the correct app site link", func() {
			g.Assert(c.App.SiteLink).Equal(SiteLink)
		})

	})
}
