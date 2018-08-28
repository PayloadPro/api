package configs

import "os"

// AppConfig is the config for the application
type AppConfig struct {
	Name       string
	APIDomain  string
	DocsDomain string
	SiteDomain string
}

// Setup the AppConfig
func (ac *AppConfig) Setup() {
	ac.Name = os.Getenv("APP_NAME")
	ac.APIDomain = os.Getenv("APP_API_DOMAIN")
	ac.DocsDomain = os.Getenv("APP_DOCS_DOMAIN")
	ac.SiteDomain = os.Getenv("APP_SITE_DOMAIN")
}
