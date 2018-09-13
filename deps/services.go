package deps

import (
	"github.com/PayloadPro/pro.payload.api/services"
)

// Services wrapped in a container
type Services struct {
	Bin     *services.BinService
	Request *services.RequestService
	Root    *services.RootService
	Stats   *services.StatsService
}
