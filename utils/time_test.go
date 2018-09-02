package utils

import (
	"testing"
	"time"

	. "github.com/franela/goblin"
)

func Test(t *testing.T) {
	g := Goblin(t)
	g.Describe("FormatTimeMeta", func() {
		g.It("Returns the correct time", func() {
			t := time.Date(2018, 9, 1, 12, 34, 45, 123456789, time.UTC)
			g.Assert(FormatTimeMeta(t)).Equal(map[string]interface{}{
				"utc": "2018-09-01 12:34:45.123456789 +0000 UTC",
				"unix": map[string]interface{}{
					"epoch": 1535805285,
					"nano":  1535805285123456789,
				},
			})
		})
	})
}
