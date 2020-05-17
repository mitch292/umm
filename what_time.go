package main

import (
	"flag"
	"fmt"
	"time"
)

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

func convertTime(newLoc string, timeToCheck time.Time) time.Time {
	tz, err := time.LoadLocation(myTimeZones[newLoc])

	if err != nil {
		panic(err)
	}

	return timeToCheck.In(tz)
}

func determineFromTime(oldLoc string, tm string) time.Time {
	tz, err := time.LoadLocation(myTimeZones[oldLoc])

	layout := "15:04:05"
	newTm, err := time.ParseInLocation(layout, tm, tz)

	if err != nil {
		return time.Now()
	}

	return newTm
}

func main() {

	fromTime := flag.String("time", time.Now().String(), "The time that you want to convert from")
	fromTz := flag.String("from-tz", "UTC", "The time zone you want to convert from, defaults to UTC")
	toTz := flag.String("to-tz", "office", "The time zone you want to conver to, defaults to the America/New_York")

	flag.Parse()

	ogTime := determineFromTime(*fromTz, *fromTime)
	fmt.Println(ogTime.String())
	newTime := convertTime(*toTz, ogTime)
	fmt.Println(newTime.Format(time.Kitchen))
}
