/*
Copyright © 2020 Andrew Mitchell <andrewpmitchell7@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "Convert a time zone from one to another",
	Long: `
	Change a time zone from one to another. Default to UTC being passed to the 
	home time zone represented in your config file. But convert times between some of your
	aliased time zones as well`,
	Run: func(cmd *cobra.Command, args []string) {

		timeToConvert, _ := cmd.Flags().GetString("convert")
		originalTz, _ := cmd.Flags().GetString("from-tz")
		newTz, _ := cmd.Flags().GetString("to-tz")
		fmt.Println(timeToConvert, originalTz, newTz)

		wt := whatTime{
			originalTimeAsString: timeToConvert,
			originalTz:           originalTz,
			newTz:                newTz,
		}

		fmt.Println(wt.ConvertTime().Format(time.Kitchen))
	},
}

func init() {
	rootCmd.AddCommand(timeCmd)

	timeCmd.Flags().StringP("convert", "c", "", "The time you want to convert")
	timeCmd.Flags().StringP("from-tz", "F", "home", "The time zone you want to convert from")
	timeCmd.Flags().StringP("to-tz", "T", "UTC", "The time zone you want to convert to")
	// TODO:Add format here

}

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

func (wt whatTime) ConvertTime() time.Time {
	tz, err := time.LoadLocation(myTimeZones[wt.newTz])

	if err != nil {
		fmt.Println("parse error:", err.Error())
	}

	return wt.originalTime().In(tz)
}

func (wt whatTime) originalTime() time.Time {
	tz, err := time.LoadLocation(myTimeZones[wt.originalTz])

	layout := "15:04:05"
	fmt.Println(wt.originalTimeAsString)
	fmt.Println(wt.originalTimeAsString)

	newTm, err := time.ParseInLocation(layout, wt.originalTimeAsString, tz)

	if err != nil {
		fmt.Println("parse error:", err.Error())
	}

	return newTm
}
