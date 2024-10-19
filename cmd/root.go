package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"
	// "net/url"

	"github.com/coolapso/megophone/internal/util"
	"github.com/coolapso/megophone/pkg/xdotcom"
	"github.com/coolapso/megophone/pkg/mastodon"

	"github.com/michimani/gotwi"
	//Twitter API V2 does not yet support uploading media therefore these "legacy" packages are needed
	"github.com/dghubble/oauth1"
	twitterv1 "github.com/drswork/go-twitter/twitter"
	gomasto "github.com/mattn/go-mastodon"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)


var	(
	cfgFile string
)

const (
	redirectUri = "urn:ietf:wg:oauth:2.0:oob"
)


var rootCmd = &cobra.Command{
	Use:   "megophone",
	Short: "Post to twitter and mastodon from your CLI",
	Long: `xm is a cli tool that allows you to post to both x (twitter)
and mastodon at the same time, with a single command, from you CLI`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) { 
		text := strings.ReplaceAll(args[0], "\\n", "\n")
		mediaPath, _ := cmd.Flags().GetString("media-path")

		if cmd.Flags().Changed("x-only") {
			fmt.Println(posting())
			if err := postX(text, mediaPath); err != nil {
				fmt.Println("Failed posting to X,", err)
				os.Exit(1)
			}
			fmt.Println(done())
			os.Exit(0)
		} 

		if cmd.Flags().Changed("m-only") {
			fmt.Println(posting())
			if err := postMastodon(text, mediaPath); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Println(done())
			os.Exit(0)
		}

		if err := postAll(text, mediaPath); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(done())
		os.Exit(0)
	},
}

func done() string {
	return "Done!"
}

func posting() string {
	return "Posting..."
}

func postX(text, mediaPath string) (err error) { 
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

	if mediaPath != "" { 
		xv1Config := oauth1.NewConfig(clientInput.APIKey, clientInput.APIKeySecret)
		token := oauth1.NewToken(clientInput.OAuthToken, clientInput.OAuthTokenSecret)
		httpClient := xv1Config.Client(oauth1.NoContext, token)
		xV1 := twitterv1.NewClient(httpClient)
		mediaType := util.GetMediaType(mediaPath)

		if mediaType != "image" && mediaType != "video" {
			return fmt.Errorf("Media type %v not supported, please make sure you provide a image or video", mediaType)
		}

		media, err := util.OpenMediaFile(mediaPath)
		if err != nil { 
			return fmt.Errorf("Failed to open media file, %v\n", err)
		}

		id, err := xdotcom.CreatePostWithMedia(context.Background(), x, xV1, text, media, mediaType)
		if err != nil {
			return err
		}
		fmt.Println("X Post created with ID:", id)

		return nil
	}


	id, err :=  xdotcom.CreatePost(context.Background(), x, text)
	if err != nil {
		return err
	}

	fmt.Println("X Post created with ID:", id)
	return nil
}

func postMastodon(text, mediaPath string) (err error) {

	config := mastodonClientConfig()
	client := gomasto.NewClient(config)

	fmt.Println(client.Config)
	fmt.Println(client.UserAgent)

	if mediaPath != "" {
		return 
	}

	id, err := mastodon.CreatePost(context.Background(), client, text, "public")
	if err != nil { 
		return fmt.Errorf("Failed to post to mastodon, %v\n", err)
	}


	fmt.Println("Toot created with ID:", id)
	return nil
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func postAll(text, mediaPath string) error {

	fmt.Println(posting())
	fmt.Printf("Posting %v to twitter and Mastodon\n", text)

	return nil
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $XDG_HOME_CONFIG/megophone/config.yaml)")
	rootCmd.Flags().BoolP("x-only", "x", false, "Post to X only")
	rootCmd.Flags().BoolP("m-only", "m", false, "Post to Mastodon Only")
	rootCmd.Flags().StringP("media-path", "p", "", "Path of media to be uploaded")
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
