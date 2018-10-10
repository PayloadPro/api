package deps

import (
	"os"
	"testing"

	"github.com/PayloadPro/api/configs"
)

func TestCockroachSetup(t *testing.T) {

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

	if c.App.Name != name {
		t.Errorf("Incorrect Name. Got: `%s`, expected: `%s`", c.App.Name, name)
	}
	if c.App.APILink != APILink {
		t.Errorf("Incorrect API Link. Got: `%s`, expected: `%s`", c.App.APILink, APILink)
	}

	if c.App.DocsLink != DocsLink {
		t.Errorf("Incorrect Docs Link. Got: `%s`, expected: `%s`", c.App.DocsLink, DocsLink)
	}

	if c.App.SiteLink != SiteLink {
		t.Errorf("Incorrect Site Link. Got: `%s`, expected: `%s`", c.App.SiteLink, SiteLink)
	}

	if c.DB.DSN != DSN {
		t.Errorf("Incorrect DSN. Got: `%s`, expected: `%s`", c.DB.DSN, DSN)
	}
}
