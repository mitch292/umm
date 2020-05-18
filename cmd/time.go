package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "Convert a time zone from one to another",
	Long: `
		Change a time zone from one to another. Default to UTC being passed to the 
		home time zone represented in your config file. But convert times between some of your
		aliased time zones as well
	`,
	Run: func(cmd *cobra.Command, args []string) {

		timeToConvert, _ := cmd.Flags().GetString("convert")
		originalTz, _ := cmd.Flags().GetString("from-tz")
		newTz, _ := cmd.Flags().GetString("to-tz")

		wt := whatTime{
			originalTimeAsString: timeToConvert,
			originalTz:           originalTz,
			newTz:                newTz,
		}

		fmt.Println(wt.convertTime().Format(time.Kitchen))
	},
}

func init() {
	rootCmd.AddCommand(timeCmd)
	timeCmd.Flags().StringP("convert", "c", "", "The time you want to convert")
	timeCmd.Flags().StringP("from-tz", "F", "utc", "The time zone you want to convert from")
	timeCmd.Flags().StringP("to-tz", "T", "home", "The time zone you want to convert to")
	// TODO:Add format here

}

type whatTime struct {
	originalTz           string
	newTz                string
	originalTimeAsString string
}

func (wt whatTime) convertTime() time.Time {
	tz, err := time.LoadLocation(wt.getTimeZone(wt.newTz))

	if err != nil {
		fmt.Println("parse error:", err.Error())
	}

	return wt.originalTime().In(tz)
}

func (wt whatTime) originalTime() time.Time {
	tz, err := time.LoadLocation(wt.getTimeZone(wt.originalTz))

	// If they don't pass us a time, lets use now
	if len(wt.originalTimeAsString) == 0 {
		return time.Now().In(tz)
	}

	layout := "15:04:05"
	origTime, err := time.ParseInLocation(layout, wt.originalTimeAsString, tz)

	if err != nil {
		fmt.Println("parse error:", err.Error())
	}

	return origTime
}

func (wt whatTime) getTimeZone(tz string) string {
	timeZones := viper.GetStringMap("timezones")

	if timeZone, ok := timeZones[tz].(string); ok {
		return timeZone
	}

	fmt.Println("This timezone was not set in your .umm.yaml file")
	return ""
}
