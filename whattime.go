package whattime

import (
	"flag"
	"fmt"
	"time"
)

func main() {

	fromTime := *flag.String("time", time.Now().String(), "The time that you want to convert from")
	fromTz := *flag.String("from-tz", "UTC", "The time zone you want to convert from, defaults to UTC")
	toTz := *flag.String("to-tz", "office", "The time zone you want to conver to, defaults to the America/New_York")

	flag.Parse()

	wt := whatTime{
		originalTimeAsString: fromTime,
		originalTz:           fromTz,
		newTz:                toTz,
	}

	fmt.Println(wt.convertTime().Format(time.Kitchen))
}
