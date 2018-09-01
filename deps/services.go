package deps

import (
	"github.com/PayloadPro/pro.payload.api/services"
)

// Services wrapped in a container
type Services struct {
	Request *services.RequestService
	Bin     *services.BinService
}
