package util

import (
	"fmt"
	"os"
	"path/filepath"
)

// GetConfigDir returns the default directory to hold configuration profiles
func GetConfigDir() (string, error) {
	userCfgDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(userCfgDir, "megophone"), nil
}

// GetConfigFileName returns the file name for a given profile
func GetConfigFileName(profile string) string {
	return fmt.Sprintf("%s.env", profile)
}

// GetConfigFilePath eturns absolute path to the profile configuration file
func GetConfigFilePath(profile string) (string, error) {
	cfgDir, err := GetConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(cfgDir, GetConfigFileName(profile)), nil
}
