package cmd

import (
	"fmt"
	"os"
	"strings"
	"context"

	"github.com/coolapso/xm-cli/internal/util"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/managetweet"
	"github.com/michimani/gotwi/tweet/managetweet/types"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)


var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "xm-cli",
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
			fmt.Printf("Posting %v only to x\n", text)
			postX(text)
			os.Exit(0)
		} 

		if cmd.Flags().Changed("m-only") {
			fmt.Printf("Posting: %v only to mastodon\n", text)
			os.Exit(0)
		}

		fmt.Printf("Posting %v to twitter and Mastodon\n", text)
	},
}

func postX(text string) { 

	if !util.IsXLenght(text) {
		fmt.Println("Text is too long for a tweet")
		os.Exit(1)
	}

	//TODO: Twitter uses API keys and access tokens
	//Set both of them
	clientInput := &gotwi.NewClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
		OAuthToken:           viper.GetString("x_oauth_token"),
		OAuthTokenSecret:     viper.GetString("x_oauth_token_secret"),
	}

	x, err := gotwi.NewClient(clientInput)
	if err != nil { 
		fmt.Println("Failed to create X Client:", err)
		os.Exit(1)
	}

	fmt.Println(clientInput)

	post := &types.CreateInput {
		Text: gotwi.String(text),
	}

	resp, err := managetweet.Create(context.Background(), x, post)
	if err != nil { 
		fmt.Println("failed to post tweet: ", err)
		os.Exit(1)
	}
	fmt.Printf("[%s] %s\n", gotwi.StringValue(resp.Data.ID), gotwi.StringValue(resp.Data.Text))
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
		// Find root config directory.
		cfgDir, err := util.GetConfigDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(cfgDir)
		viper.SetConfigName("xm-cli.env")
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
