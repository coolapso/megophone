package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/coolapso/megophone/internal/util"
	"github.com/coolapso/megophone/pkg/mastodon"
	"github.com/coolapso/megophone/pkg/xdotcom"

	"github.com/michimani/gotwi"

	//Twitter API V2 does not yet support uploading media therefore these "legacy" packages are needed
	"github.com/dghubble/oauth1"
	twitterv1 "github.com/drswork/go-twitter/twitter"
	gomasto "github.com/mattn/go-mastodon"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	profile string
)

const (
	redirectUri = "urn:ietf:wg:oauth:2.0:oob"
)

var rootCmd = &cobra.Command{
	Use:   "megophone",
	Short: "Post to multiple social networks from your CLI",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		text := strings.ReplaceAll(args[0], "\\n", "\n")
		mediaPath, _ := cmd.Flags().GetString("media-path")

		if cmd.Flags().Changed("x-only") {
			if err := postX(text, mediaPath); err != nil {
				fmt.Println("Failed posting to X,", err)
				os.Exit(1)
			}
			os.Exit(0)
		}

		if cmd.Flags().Changed("m-only") {
			if err := postMastodon(text, mediaPath); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			os.Exit(0)
		}

		if errors := postAll(text, mediaPath); len(errors) > 0 {
			for _, err := range errors {
				fmt.Println(err)
			}
			os.Exit(1)
		}
		os.Exit(0)
	},
}

func postX(text, mediaPath string) (err error) {
	clientInput := &gotwi.NewClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
		APIKey:               viper.GetString("x_api_key"),
		APIKeySecret:         viper.GetString("x_api_key_secret"),
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

	id, err := xdotcom.CreatePost(context.Background(), x, text)
	if err != nil {
		return err
	}

	fmt.Println("X Post created with ID:", id)
	return nil
}

func postMastodon(text, mediaPath string) (err error) {

	config := mastodonClientConfig()
	client := gomasto.NewClient(config)

	if mediaPath != "" {
		media, err := client.UploadMedia(context.Background(), mediaPath)
		if err != nil {
			return fmt.Errorf("Failed to upload media file, %v\n", err)
		}

		id, err := mastodon.CreatePostWithMedia(context.Background(), client, string(media.ID), text, "public")
		if err != nil {
			return err
		}
		toothCreated(id)

		return nil
	}

	id, err := mastodon.CreatePost(context.Background(), client, text, "public")
	if err != nil {
		return fmt.Errorf("Failed to post to mastodon, %v\n", err)
	}

	toothCreated(id)
	return nil
}

func toothCreated(id string) {
	fmt.Println("Toot created with ID:", id)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func postAll(text, mediaPath string) (errors []error) {
	var wg sync.WaitGroup
	errChan := make(chan error, 2)
	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := postX(text, mediaPath); err != nil {
			errChan <- err
		}
		errChan <- nil

	}()

	go func() {
		defer wg.Done()
		if err := postMastodon(text, mediaPath); err != nil {
			errChan <- err
		}
		errChan <- nil

	}()

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			errors = append(errors, err)
		}
	}

	return errors
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $XDG_HOME_CONFIG/megophone/config.yaml)")
	rootCmd.PersistentFlags().StringVar(&profile, "profile", "default", "The configuration profile to use")
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
		viper.SetConfigName(fmt.Sprintf("%s.env", profile))
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
