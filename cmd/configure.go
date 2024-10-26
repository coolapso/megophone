package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/coolapso/megophone/internal/util"
	"github.com/coolapso/megophone/pkg/mastodon"
	"github.com/coolapso/megophone/pkg/xdotcom"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var c config

var configure = &cobra.Command{
	Use:   "configure",
	Short: "Configures megophone",
	Long: `Creates megophone configuration file in $XDG_HOME_CONFIG and populates it
	with settings and secrets necessary to talk to apis`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return configMegophone(bufio.NewReader(os.Stdin))
	},
}

type config struct {
	x xdotcom.Secrets
	m mastodon.Secrets
}

func configMegophone(reader *bufio.Reader) error {
	loadXVars(&c)
	configX(reader, &c)
	if err := configMastodon(context.Background(), reader, &c, profile); err != nil {
		return err
	}

	if err := writeConfigFile(profile); err != nil {
		return err
	}

	return nil
}

func writeConfigFile(p string) error {
	cfgDir, err := util.GetConfigDir()
	if err != nil {
		return fmt.Errorf("Failed to get config directory: %v", err.Error())
	}

	cfgFilePath, err := util.GetConfigFilePath(p)
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
