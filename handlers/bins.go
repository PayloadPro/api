package handlers

import (
	"net/http"

	"github.com/PayloadPro/api/entities"
	"github.com/PayloadPro/api/responses"
	"github.com/gofiber/fiber"
)

func CreateBin(c *fiber.Ctx) {
	bin := entities.Bin{}
	c.JSON(responses.SingleResponse{
		Data: bin.Data(),
		Meta: responses.Meta{
			Links: responses.Links(),
		},
	})
	c.SendStatus(http.StatusOK)
}

func GetBins(c *fiber.Ctx) {
	var bins entities.Bins
	c.JSON(responses.MultipleResponse{
		Data: bins.Data(),
		Meta: responses.Meta{
			Links: responses.Links(),
		},
	})
	c.SendStatus(http.StatusOK)
}

func GetBin(c *fiber.Ctx) {
	// bin := entities.Bin{}
	c.JSON(responses.SingleResponse{
		Data: responses.Data{
			Type: "bins",
		},
		Meta: responses.Meta{
			Links: responses.Links(),
		},
	})
	c.SendStatus(http.StatusOK)
}
