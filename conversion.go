package whattime

import "time"

var myTimeZones = map[string]string{
	"dc":       "America/New_York",
	"home":     "America/New_York",
	"office":   "America/New_York",
	"central":  "America/Chicago",
	"pacific":  "America/LosAngeles",
	"mountain": "America/Denver",
	"kevin":    "Asia/Seoul",
	"sk":       "Asia/Seoul",
	"korea":    "Asia/Seoul",
	"utc":      "UTC",
}

type whatTime struct {
	originalTz           string
	newTz                string
	originalTimeAsString string
}

func (wt whatTime) originalTime() time.Time {
	// TODO: make this a method
	tz, err := time.LoadLocation(wt.originalTz)

	layout := "15:04:05"
	newTm, err := time.ParseInLocation(layout, wt.originalTimeAsString, tz)

	if err != nil {
		return time.Now()
	}

	return newTm
}

func (wt whatTime) convertTime() time.Time {
	tz, err := time.LoadLocation(myTimeZones[wt.newTz])

	if err != nil {
		panic(err)
	}

	return wt.originalTime().In(tz)
}
