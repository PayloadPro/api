package handlers

import (
	"net/http"

	"github.com/PayloadPro/api/responses"
	"github.com/gofiber/fiber"
)

func Root(c *fiber.Ctx) {
	a := make(map[string]interface{}, 0)
	a["message"] = "Welcome to Payload Pro"

	c.JSON(responses.SingleResponse{
		Data: responses.Data{
			Type:       "root",
			Attributes: a,
		},
		Meta: responses.Meta{
			Links: responses.Links(),
		},
	})
	c.SendStatus(http.StatusOK)
}
