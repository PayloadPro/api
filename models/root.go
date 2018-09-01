package models

import (
	"github.com/PayloadPro/pro.payload.api/configs"
	"github.com/google/jsonapi"
)

// Root is a designated root endpoint
type Root struct {
	ID      string `jsonapi:"primary,root"`
	Message string `jsonapi:"attr,message"`
	Config  *configs.AppConfig
}

// NewRoot generates a Root struct to use
func NewRoot(config *configs.AppConfig) (*Root, error) {
	return &Root{
		Message: "Welcome to the Payload Pro API",
		Config:  config,
	}, nil
}

// JSONAPILinks return links for the JSONAPI marshal
func (r Root) JSONAPILinks() *jsonapi.Links {
	return &jsonapi.Links{
		"api":  r.Config.APILink,
		"docs": r.Config.DocsLink,
		"site": r.Config.SiteLink,
	}
}
