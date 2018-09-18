package models

import (
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/PayloadPro/api/configs"
	"github.com/PayloadPro/api/utils"
	. "github.com/franela/goblin"
	"github.com/google/jsonapi"
)

func TestBinModel(t *testing.T) {

	g := Goblin(t)

	g.Describe("NewBin creates a bin based on the HTTP Request", func() {

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

	})

	g.Describe("JSONAPI funcs return correct values", func() {

		g.It("JSONAPILinks() returns correct self link", func() {
			b := Bin{
				ID: "3d760a4b-cc85-46d6-bc23-fc77b4c68f30",
				Config: &configs.AppConfig{
					APILink: "https://api.payload.pro",
				},
			}
			s := "https://api.payload.pro/bins/3d760a4b-cc85-46d6-bc23-fc77b4c68f30"
			g.Assert(b.JSONAPILinks()).Equal(&jsonapi.Links{"self": s})
		})

		g.It("JSONAPIMeta() returns correct created time", func() {
			now := time.Now()
			b := Bin{
				Created: now,
			}
			e := utils.FormatTimeMeta(now)
			g.Assert(b.JSONAPIMeta()).Equal(&jsonapi.Meta{"created": e})
		})

	})

	g.Describe("Errors are expected", func() {
		g.It("Bin not found returns correct error", func() {
			nf := ErrBinNotFound
			g.Assert(nf.Error()).Equal("Bin could not be found")
		})
	})
}
