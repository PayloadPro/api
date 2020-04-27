package main

import (
	"net/http"

	"github.com/gofiber/fiber"
)

type response struct {
	Data   data        `json:"data,omitempty"`
	Errors []respError `json:"errors,omitempty"`
}

type data struct {
	Type       string                 `json:"type,omitempty"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Links      map[string]string      `json:"links,omitempty"`
}

type respError struct {
	Message string `json:"message"`
}

func errHandler(c *fiber.Ctx, err error) {
	s := http.StatusInternalServerError
	c.JSON(response{
		Errors: []respError{
			{
				Message: err.Error(),
			},
		},
	})
	c.SendStatus(s)
}
