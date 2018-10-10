package configs

import (
	"os"
	"testing"

	. "github.com/franela/goblin"
)

func TestAppSetup(t *testing.T) {

	g := Goblin(t)

	name := "Payload Pro"
	APILink := "https://api.payload.pro"
	DocsLink := "https://docs.payload.pro"
	SiteLink := "https://payload.pro"

	os.Setenv("APP_NAME", name)
	os.Setenv("APP_API_LINK", APILink)
	os.Setenv("APP_DOCS_LINK", DocsLink)
	os.Setenv("APP_SITE_LINK", SiteLink)

	ac := &AppConfig{}
	ac.Setup()

	g.Describe("Returns OS env vars as values", func() {

		g.It("Contains the correct App Name", func() {
			g.Assert(ac.Name).Equal(name)
		})

		g.It("Contains the correct API link", func() {
			g.Assert(ac.APILink).Equal(APILink)
		})

		g.It("Contains the correct docs link", func() {
			g.Assert(ac.DocsLink).Equal(DocsLink)
		})

		g.It("Contains the correct site link", func() {
			g.Assert(ac.SiteLink).Equal(SiteLink)
		})

	})

}
