package cmd

import (
	"fmt"
	"os"
	"strings"
	"bufio"

	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
)

var configure = &cobra.Command{
	Use: "config",
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


func maskString(s string) string {
	l := len(s)
	if l <= 3 {
		return s
	}

	return strings.Repeat("*", l-3) + s[l-3:]

}

func loadXVars(c *config) {
	if user, isSet := os.LookupEnv("X_USER"); isSet { 
		c.xUser = user
	}

	if key, isSet := os.LookupEnv("X_API_KEY"); isSet { 
		c.xApiKey = key
	}
}

func loadMastodonVars(c *config) {
	if user, isSet := os.LookupEnv("MASTODON_USER"); isSet { 
		c.mUser = user
	}

	if key, isSet := os.LookupEnv("MASTODON_API_KEY"); isSet { 
		c.mApiKey = key
	}
}

func configxm() error {
	var c config
	// Before setting config file, load any env variables 
	loadXVars(&c)
	loadMastodonVars(&c)
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("X Username(%v): ", c.xUser)
	c.configXUser(reader)

	fmt.Printf("X API Key(%v): ", maskString(c.xApiKey))
	c.configXApiKey(reader)

	fmt.Printf("Mastodon Username(%v): ", c.xUser)
	c.configMastodonUser(reader)

	fmt.Printf("Mastodon API Key(%v): ", maskString(c.xApiKey))
	c.configMastodonApiKey(reader)

	return nil
}

// func createConfigFile() error { 
//
// }
