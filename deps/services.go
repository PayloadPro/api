package deps

import (
	"github.com/andrew-waters/pro.payload.api/services"
)

// Services wrapped in a container
type Services struct {
	Payload *services.PayloadService
}
