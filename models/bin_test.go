package models

import (
	"net/http"
	"strings"
	"testing"

	"github.com/PayloadPro/api/configs"
	. "github.com/franela/goblin"
	"github.com/google/jsonapi"
)

func TestBinModel(t *testing.T) {

	g := Goblin(t)

	g.Describe("Bins are handled correctly", func() {

		g.It("NewBin() returns a correct bin", func() {

			e := Bin{
				Name: "I'm a bin",
				Config: &configs.AppConfig{
					APILink: "https://api.payload.pro",
				},
			}

			data := []byte(`{"data":{"attributes":{"name":"I'm a bin"}}}`)

			r, _ := http.NewRequest("GET", "", strings.NewReader(string(data)))
			c := &configs.AppConfig{
				APILink: "https://api.payload.pro",
			}

			bin, _ := NewBin(r, c)
			g.Assert(bin.Name).Equal(e.Name)
			g.Assert(bin.Config).Equal(e.Config)

		})

		g.It("Bin not found returns correct error", func() {
			nf := ErrBinNotFound
			g.Assert(nf.Error()).Equal("Bin could not be found")
		})

	})

	g.Describe("JSONAPI utils return correct values", func() {

		bin := Bin{
			ID: "3d760a4b-cc85-46d6-bc23-fc77b4c68f30",
			Config: &configs.AppConfig{
				APILink: "https://api.payload.pro",
			},
			Stats: &Stats{
				Total:   100,
				GET:     5,
				POST:    8,
				PUT:     11,
				PATCH:   14,
				OPTIONS: 17,
				HEAD:    20,
				DELETE:  25,
			},
		}
		binLinks := bin.JSONAPILinks()

		expectedLinks := &jsonapi.Links{
			"self":     "https://api.payload.pro/bins/3d760a4b-cc85-46d6-bc23-fc77b4c68f30",
			"request":  "https://api.payload.pro/bins/3d760a4b-cc85-46d6-bc23-fc77b4c68f30/request",
			"requests": "https://api.payload.pro/bins/3d760a4b-cc85-46d6-bc23-fc77b4c68f30/requests",
		}

		g.It("JSONAPILinks() returns correct links", func() {
			g.Assert(binLinks).Equal(expectedLinks)
		})

	})

}
