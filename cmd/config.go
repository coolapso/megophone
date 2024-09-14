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
		return configxm()
	},
}

type config struct {
	xUser	string
	xApiKey string
	mUser	string
	mApiKey string
}

func(c *config) configXUser(r *bufio.Reader) { 
	i, _ := r.ReadString('\n')
	if s := strings.TrimSpace(i); s != "" { 
		c.xUser = s
	}
}

func(c *config) configXApiKey (r *bufio.Reader) {
	i, _ := r.ReadString('\n')
	if s := strings.TrimSpace(i); s != "" { 
		c.xApiKey = s
	}
}

func(c *config) configMastodonUser(r *bufio.Reader) { 
	i, _ := r.ReadString('\n')
	if s := strings.TrimSpace(i); s != "" { 
		c.mUser = s
	}
}

func(c *config) configMastodonApiKey (r *bufio.Reader) {
	i, _ := r.ReadString('\n')
	if s := strings.TrimSpace(i); s != "" { 
		c.mApiKey = s
	}
}

func configxm() error {
	// Before setting config file, load any env variables into config struct
	loadXVars(&c)
	loadMastodonVars(&c)
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("X Username(%v): ", c.xUser)
	c.configXUser(reader)
	viper.Set("x_user", c.xUser)

	fmt.Printf("X API Key(%v): ", util.MaskString(c.xApiKey))
	c.configXApiKey(reader)
	viper.Set("x_api_key", c.xApiKey)

	fmt.Printf("Mastodon Username(%v): ", c.mUser)
	c.configMastodonUser(reader)
	viper.Set("mastodon_user", c.mUser)

	fmt.Printf("Mastodon API Key(%v): ", util.MaskString(c.mApiKey))
	c.configMastodonApiKey(reader)
	viper.Set("mastodon_api_key", c.mApiKey)

	if err := writeConfigFile(); err != nil {
		return err
	}

	return nil
}

func loadXVars(c *config) {
	c.xUser = viper.GetString("x_user")
	if user, isSet := os.LookupEnv("XM_X_USER"); isSet { 
		c.xUser = user
	}

	c.xApiKey = viper.GetString("x_api_key")
	if key, isSet := os.LookupEnv("XM_X_API_KEY"); isSet { 
		c.xApiKey = key
	}
}

func loadMastodonVars(c *config) {
	c.mUser = viper.GetString("mastodon_user")
	if user, isSet := os.LookupEnv("XM_MASTODON_USER"); isSet { 
		c.mUser = user
	}

	c.mApiKey = viper.GetString("mastodon_api_key")
	if key, isSet := os.LookupEnv("XM_MASTODON_API_KEY"); isSet { 
		c.mApiKey = key
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
