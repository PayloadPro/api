package configs

import (
	"os"
	"testing"
)

func TestAppSetup(t *testing.T) {

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

	if ac.Name != name {
		t.Errorf("Incorrect App Name. Got: `%s`, expected: `%s`", ac.Name, name)
	}

	if ac.APILink != APILink {
		t.Errorf("Incorrect API Link. Got: `%s`, expected: `%s`", ac.APILink, APILink)
	}

	if ac.DocsLink != DocsLink {
		t.Errorf("Incorrect Docs Link. Got: `%s`, expected: `%s`", ac.DocsLink, DocsLink)
	}

	if ac.SiteLink != SiteLink {
		t.Errorf("Incorrect Site Link. Got: `%s`, expected: `%s`", ac.SiteLink, SiteLink)
	}
}
