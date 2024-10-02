package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/coolapso/xm-cli/internal/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var c config

var configure = &cobra.Command{
	Use: "configure",
	Short: "Configures xm-cli",
	Long: `Creates xm-cli configuration file in $XDG_HOME_CONFIG and populates it
	with settings`,
	RunE: func(cmd *cobra.Command, args []string) error { 
		return configxm(bufio.NewReader(os.Stdin))
	},
}

type xdotcom struct {
	oauthToken string
	oauthTokenSecret string
}

func(x *xdotcom) configOauthToken(r *bufio.Reader) { 
	i, _ := r.ReadString('\n')
	if s := strings.TrimSpace(i); s != "" { 
		x.oauthToken = s
	}
}

func(x *xdotcom) configOauthTokenSecret (r *bufio.Reader) {
	i, _ := r.ReadString('\n')
	if s := strings.TrimSpace(i); s != "" { 
		x.oauthTokenSecret = s
	}
}


type mastodon struct {
	apiKey string
	apiKeySecret string
}

func(m *mastodon) configApiKey(r *bufio.Reader) { 
	i, _ := r.ReadString('\n')
	if s := strings.TrimSpace(i); s != "" { 
		m.apiKey = s
	}
}

func(m *mastodon) configApiKeySecret (r *bufio.Reader) {
	i, _ := r.ReadString('\n')
	if s := strings.TrimSpace(i); s != "" { 
		m.apiKeySecret = s
	}
}


type config struct {
	x xdotcom
	m mastodon
}

func configxm(reader *bufio.Reader) error {
	// Before setting config file, load any env variables into config struct
	loadXVars(&c)
	loadMastodonVars(&c)

	fmt.Printf("X Api Key(%v): ", c.x.oauthToken)
	c.x.configOauthToken(reader)
	viper.Set("x_oauth_token", c.x.oauthToken)

	fmt.Printf("X API Key Secret(%v): ", util.MaskString(c.x.oauthTokenSecret))
	c.x.configOauthTokenSecret(reader)
	viper.Set("x_oauth_token_secret", c.x.oauthTokenSecret)

	fmt.Printf("Mastodon Api Key(%v): ", c.m.apiKey)
	c.m.configApiKey(reader)
	viper.Set("mastodon_api_key", c.m.apiKey)

	fmt.Printf("Mastodon API Key Secret(%v): ", util.MaskString(c.m.apiKeySecret))
	c.m.configApiKeySecret(reader)
	viper.Set("mastodon_api_key_secret", c.m.apiKeySecret)

	if err := writeConfigFile(); err != nil {
		return err
	}

	return nil
}

func loadXVars(c *config) {
	c.x.oauthToken = viper.GetString("x_oauth_token")
	if key, isSet := os.LookupEnv("XM_X_OAUTH_TOKEN"); isSet { 
		c.x.oauthToken = key
	}

	c.x.oauthTokenSecret = viper.GetString("x_oauth_token_secret")
	if secret, isSet := os.LookupEnv("XM_X_OAUTH_TOKEN_SECRET"); isSet { 
		c.x.oauthTokenSecret = secret
	}
}

func loadMastodonVars(c *config) {
	c.m.apiKey = viper.GetString("mastodon_api_key")
	if key, isSet := os.LookupEnv("XM_MASTODON_API_KEY"); isSet { 
		c.m.apiKey = key
	}

	c.m.apiKeySecret = viper.GetString("mastodon_api_key_secret")
	if secret, isSet := os.LookupEnv("XM_MASTODON_API_KEY_SECRET"); isSet { 
		c.m.apiKeySecret = secret
	}
}

func writeConfigFile() error {
	cfgDir, err := util.GetConfigDir()
	if err != nil { 
		return fmt.Errorf("Failed to get config directory: %v", err.Error())
	}

	cfgFilePath, err := util.GetConfigFilePath()
	if err != nil {
		return fmt.Errorf("Failed to get config file path: %v", err.Error())
	}

	// Create necessary directories if they don't exist
	if _, err := os.Stat(cfgDir); os.IsNotExist(err) {
		err := os.MkdirAll(cfgDir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	if err := viper.WriteConfigAs(cfgFilePath); err != nil { 
		return fmt.Errorf("Failed to create config file: %v", err.Error())
	}

	return nil
}
