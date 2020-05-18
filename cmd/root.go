package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "umm",
	Short: "umm helps you when you dont feel like thinking",
	Long:  `umm can help you convert a time from utc to your local time`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("try typing -h if you need any help")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
