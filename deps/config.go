package deps

import "github.com/andrew-waters/payload.pro/configs"

// Config wrapped in a container
type Config struct {
	App *configs.AppConfig
	DB  *configs.DatabaseConfig
}

func (c Config) init() {
	c.App.Setup()
	c.DB.Setup()
}
