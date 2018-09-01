package configs

import "os"

// AppConfig is the config for the application
type AppConfig struct {
	APILink  string
	DocsLink string
	Name     string
	SiteLink string
}

// Setup the AppConfig
func (ac *AppConfig) Setup() {
	ac.APILink = os.Getenv("APP_API_LINK")
	ac.DocsLink = os.Getenv("APP_DOCS_LINK")
	ac.Name = os.Getenv("APP_NAME")
	ac.SiteLink = os.Getenv("APP_SITE_LINK")
}
