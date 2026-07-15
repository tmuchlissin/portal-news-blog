package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgfile string

var rootCmd = &cobra.Command{
	Use:   "core-api",
	Short: "this api for news portal",
	Run: func(cmd *cobra.Command, args []string) {
		startCmd.Run(startCmd, args)
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgfile, "config", "", "config file (default is .env)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgfile != "" {
		viper.SetConfigFile(cfgfile)
	} else {
		viper.SetConfigFile(".env")
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, "failed to read config file:", err)
		return
	}

	fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
}
