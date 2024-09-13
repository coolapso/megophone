package util

import (
	"os"
	"path/filepath"
)

func GetConfigDir() (string, error) {
	userCfgDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(userCfgDir, "xm-cli"), nil
}

func GetConfigFilePath() (string, error) {
	cfgDir, err := GetConfigDir()
	if err != nil {
		return "", err
	}

	if os.Getenv("GOLANG_TESTING") == "true" {
		return filepath.Join(cfgDir, ".twitch-cli-test.env"), nil
	}
	
	return filepath.Join(cfgDir, "xm-cli.env"), nil
}
