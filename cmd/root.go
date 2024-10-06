package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/coolapso/megophone/internal/util"
	"github.com/coolapso/megophone/pkg/xdotcom"

	"github.com/michimani/gotwi"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)


var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "megophone",
	Short: "Post to twitter and mastodon from your CLI",
	Long: `xm is a cli tool that allows you to post to both x (twitter)
and mastodon at the same time, with a single command, from you CLI`,
	Args: cobra.MinimumNArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if !util.IsToothLenght(args[0]) {
			return fmt.Errorf("Message to long for Mastodon")
		}

		if !util.IsXLenght(args[0]) {
			return fmt.Errorf("Message too long for X")
		}

		return nil
	},

	Run: func(cmd *cobra.Command, args []string) { 
		text := strings.ReplaceAll(args[0], "\\n", "\n")

		if cmd.Flags().Changed("x-only") {
			fmt.Println("Posting...")
			if err := postX(text); err != nil {
				fmt.Println("Failed posting to X,", err)
				os.Exit(1)
			}
			fmt.Println("Done!")
		} 

		if cmd.Flags().Changed("m-only") {
			fmt.Printf("Posting: %v only to mastodon\n", text)
			os.Exit(0)
		}

		// fmt.Printf("Posting %v to twitter and Mastodon\n", text)
	},
}

func postX(text string) (err error) { 
	clientInput := &gotwi.NewClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
		APIKey:				  viper.GetString("x_api_key"),
		APIKeySecret:		  viper.GetString("x_api_key_secret"),
		OAuthToken:           viper.GetString("x_oauth_token"),
		OAuthTokenSecret:     viper.GetString("x_oauth_token_secret"),
	}

	x, err := gotwi.NewClient(clientInput)
	if err != nil { 
		return fmt.Errorf("Failed to create X Client: %v\n", err)
	}

	id, err :=  xdotcom.CreatePost(context.Background(), x, text)
	if err != nil {
		return err
	}

	fmt.Println("X Post created with ID:", id)
	return nil
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $XDG_HOME_CONFIG/megophone/config.yaml)")
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
		// Find root config directory.
		cfgDir, err := util.GetConfigDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(cfgDir)
		viper.SetConfigName("megophone.env")
		viper.SetConfigType("env")
	}

	viper.SetEnvPrefix("xm")
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		return
		// fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
