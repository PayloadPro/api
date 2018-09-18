package deps

import "github.com/PayloadPro/api/configs"

// Config wrapped in a container
type Config struct {
	App *configs.AppConfig
	DB  *configs.CockroachConfig
}

// Setup the config
func (c Config) Setup() {
	c.App.Setup()
	c.DB.Setup()
}
