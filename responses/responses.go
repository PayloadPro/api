package responses

import (
	"net/http"

	"github.com/gofiber/fiber"
)

type SingleResponse struct {
	Data   Data        `json:"data,omitempty"`
	Errors []RespError `json:"errors,omitempty"`
	Meta   Meta        `json:"meta,omitempty"`
}

type MultipleResponse struct {
	Data   []Data      `json:"data,omitempty"`
	Errors []RespError `json:"errors,omitempty"`
	Meta   Meta        `json:"meta,omitempty"`
}

type Data struct {
	Type       string                 `json:"type,omitempty"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Links      map[string]string      `json:"links,omitempty"`
}

type Meta struct {
	Links map[string]string `json:"links,omitempty"`
}

type RespError struct {
	Message string `json:"message"`
}

func NotFound(c *fiber.Ctx) {
	s := http.StatusNotFound
	c.JSON(SingleResponse{
		Errors: []RespError{
			{
				Message: "Not Found",
			},
		},
		Meta: Meta{
			Links: Links(),
		},
	})
	c.SendStatus(s)
}

func ErrHandler(c *fiber.Ctx, err error) {
	s := http.StatusInternalServerError
	c.JSON(SingleResponse{
		Errors: []RespError{
			{
				Message: err.Error(),
			},
		},
		Meta: Meta{
			Links: Links(),
		},
	})
	c.SendStatus(s)
}
