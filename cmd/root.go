package cmd

import (
	"fmt"
	"os"

	"github.com/coolapso/xm-cli/internal/util"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "xm-cli",
	Short: "Post to twitter and mastodon from your CLI",
	Long: `xm is a cli tool that allows you to post to both x (twitter)
	and mastodon at the same time, with a single command, from you CLI`,

	Run: func(cmd *cobra.Command, args []string) { 
		if cmd.Flags().Changed("x-only") {
			fmt.Println("Posting only to x")
			os.Exit(0)
		} 

		if cmd.Flags().Changed("m-only") {
			fmt.Println("Posting only to mastodon")
			os.Exit(0)
		}

		fmt.Println("Posting to twitter and Mastodon")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $XDG_HOME_CONFIG/xm-cli/config.yaml)")
	rootCmd.Flags().BoolP("x-only", "x", false, "Post to X only")
	rootCmd.Flags().BoolP("m-only", "m", false, "Post to Mastodon Only")
	rootCmd.AddCommand(configure)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find roo config directory.
		cfgDir, err := util.GetConfigDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(cfgDir)
		viper.SetConfigName("xm-cli.env")
		viper.SetConfigType("env")
	}

	viper.SetEnvPrefix("xm")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
