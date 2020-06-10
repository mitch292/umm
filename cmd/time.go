package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "Some help with the time",
	Long: `
		Get the current time. Or get the current time in another place.
		How about convert the time from UTC to your local time. Alias time zones
		in your .umm.yaml file. Make time fun and easy.
	`,
	Run: func(cmd *cobra.Command, args []string) {

		timeToConvert, _ := cmd.Flags().GetString("convert")
		timeFormat, _ := cmd.Flags().GetString("format")
		originalTz, _ := cmd.Flags().GetString("from")
		newTz, _ := cmd.Flags().GetString("to")

		wt := whatTime{
			originalTimeAsString: timeToConvert,
			originalFormat:       timeFormat,
			originalTz:           originalTz,
			newTz:                newTz,
		}

		// Return it in the user's declared time format, defaults to time.Kitchen
		convertedTime, err := wt.convertTime()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(convertedTime.Format(viper.GetString("timeFormat")))
		}
	},
}

func init() {
	rootCmd.AddCommand(timeCmd)
	timeCmd.Flags().StringP("convert", "c", "", "The time you want to convert")
	timeCmd.Flags().StringP("format", "f", "15:04:05", "The format of the time to convert, it must be a subset of this string, 'Mon Jan 2 15:04:05 MST 2006'")
	timeCmd.Flags().StringP("from", "F", "utc", "The time zone you want to convert from")
	timeCmd.Flags().StringP("to", "T", "home", "The time zone you want to convert to")

	viper.SetDefault("timeFormat", time.Kitchen)
}

type whatTime struct {
	originalTimeAsString string
	originalFormat       string
	originalTz           string
	newTz                string
}

func (wt whatTime) convertTime() (time.Time, error) {
	tz, err := time.LoadLocation(wt.getTimeZone(wt.newTz))

	if err != nil {
		return time.Now(), err
	}

	convertedTime, err := wt.originalTime()

	if err != nil {
		return time.Now(), err
	}

	return convertedTime.In(tz), nil
}

func (wt whatTime) originalTime() (time.Time, error) {
	tz, err := time.LoadLocation(wt.getTimeZone(wt.originalTz))

	if err != nil {
		return time.Now(), err
	}

	// If they don't pass us a time, lets use now
	if len(wt.originalTimeAsString) == 0 {
		return time.Now().In(tz), nil
	}

	origTime, err := time.ParseInLocation(wt.originalFormat, wt.originalTimeAsString, tz)

	if err != nil {
		return time.Now(), err
	}

	return origTime, nil
}

func (wt whatTime) getTimeZone(tz string) string {
	timeZones := viper.GetStringMap("timezones")

	if timeZone, ok := timeZones[tz].(string); ok {
		return timeZone
	}

	fmt.Println("This timezone was not set in your .umm.yaml file, returning utc")
	return "UTC"
}
