package deps

import (
	"github.com/andrew-waters/payload.pro/services"
)

// Services wrapped in a container
type Services struct {
	Payload *services.PayloadService
}
