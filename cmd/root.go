package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "umm",
		Short: "umm helps you when you dont feel like thinking",
		Long:  `umm can help you convert a time from utc to your local time`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("try typing -h if you need any help")
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.umm.yaml)")
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		er(err)
	}
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(".umm.yaml")
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			er(err)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".umm")
	}

	viper.AutomaticEnv()

	viper.ReadInConfig()

}
