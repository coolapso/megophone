package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/coolapso/megophone/internal/util"
	"github.com/spf13/viper"
)

func loadXVars(c *config) {
	c.x.SetOauthToken(viper.GetString("x_oauth_token"))
	if key, isSet := os.LookupEnv("MEGOPHONE_X_OAUTH_TOKEN"); isSet {
		c.x.SetOauthToken(key)
	}

	c.x.SetOauthTokenSecret(viper.GetString("x_oauth_token_secret"))
	if secret, isSet := os.LookupEnv("MEGOPHONE_X_OAUTH_TOKEN_SECRET"); isSet {
		c.x.SetOauthTokenSecret(secret)
	}

	c.x.SetApiKey(viper.GetString("x_api_key"))
	if key, isSet := os.LookupEnv("MEGOPHONE_X_API_KEY"); isSet {
		c.x.SetApiKey(key)
	}

	c.x.SetApiKeySecret(viper.GetString("x_api_key_secret"))
	if secret, isSet := os.LookupEnv("MEGOPHONE_X_API_KEY_SECRET"); isSet {
		c.x.SetApiKeySecret(secret)
	}
}

func configX(reader *bufio.Reader, c *config) {
	fmt.Printf("X Oauth Token(%v): ", util.MaskString(c.x.GetOauthToken()))
	oauthTokenInput, _ := reader.ReadString('\n')
	if cleanInput := strings.TrimSpace(oauthTokenInput); cleanInput != "" {
		c.x.SetOauthToken(cleanInput)
	}
	viper.Set("x_oauth_token", c.x.GetOauthToken())

	fmt.Printf("X Oauth Token Secret(%v): ", util.MaskString(c.x.GetOauthTokenSecret()))
	oauthTokenSecretInput, _ := reader.ReadString('\n')
	if cleanInput := strings.TrimSpace(oauthTokenSecretInput); cleanInput != "" {
		c.x.SetOauthTokenSecret(cleanInput)
	}
	viper.Set("x_oauth_token_secret", c.x.GetOauthTokenSecret())

	fmt.Printf("X API Key(%v): ", util.MaskString(c.x.GetApiKey()))
	GetApiKeyInput, _ := reader.ReadString('\n')
	if cleanInput := strings.TrimSpace(GetApiKeyInput); cleanInput != "" {
		c.x.SetApiKey(cleanInput)
	}
	viper.Set("x_api_key", c.x.GetApiKey())

	fmt.Printf("X API Key Secret(%v): ", util.MaskString(c.x.GetApiKeySecret()))
	GetApiKeySecretInput, _ := reader.ReadString('\n')
	if cleanInput := strings.TrimSpace(GetApiKeySecretInput); cleanInput != "" {
		c.x.SetApiKeySecret(cleanInput)
	}
	viper.Set("x_api_key_secret", c.x.GetApiKeySecret())
}
