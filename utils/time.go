package utils

import "time"

// FormatTimeMeta takes an incomign time and converts it to a standard meta format
func FormatTimeMeta(t time.Time) map[string]interface{} {
	return map[string]interface{}{
		"utc": t.UTC(),
		"unix": map[string]interface{}{
			"epoch": t.Unix(),
			"nano":  t.UnixNano(),
		},
	}
}
