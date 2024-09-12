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
	// Before setting config file, load any env variables 
	loadXVars(&c)
	loadMastodonVars(&c)
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("X Username(%v): ", c.xUser)
	c.configXUser(reader)
	viper.Set("x_user", c.xUser)

	fmt.Printf("X API Key(%v): ", maskString(c.xApiKey))
	c.configXApiKey(reader)
	viper.Set("x_api_key", c.xApiKey)

	fmt.Printf("Mastodon Username(%v): ", c.mUser)
	c.configMastodonUser(reader)
	viper.Set("mastodon_user", c.mUser)

	fmt.Printf("Mastodon API Key(%v): ", maskString(c.mApiKey))
	c.configMastodonApiKey(reader)
	viper.Set("mastodon_api_key", c.mApiKey)

	if err := writeConfigFile(); err != nil {
		return err
	}

	return nil
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

	// if the full path doesn't exist, make all the folders to get there
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

func maskString(s string) string {
	l := len(s)
	if l <= 3 {
		return s
	}

	return strings.Repeat("*", l-3) + s[l-3:]

}

func loadXVars(c *config) {
	if user, isSet := os.LookupEnv("XM_X_USER"); isSet { 
		c.xUser = user
	}

	if key, isSet := os.LookupEnv("XM_X_API_KEY"); isSet { 
		c.xApiKey = key
	}
}

func loadMastodonVars(c *config) {
	if user, isSet := os.LookupEnv("XM_MASTODON_USER"); isSet { 
		c.mUser = user
	}

	if key, isSet := os.LookupEnv("XM_MASTODON_API_KEY"); isSet { 
		c.mApiKey = key
	}
}

