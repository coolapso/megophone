package cmd

import (
	"github.com/coolapso/megophone/internal/util"
	"os"
	"testing"
)

func TestWriteConfigFile(t *testing.T) {
	profile := "megophone-test"

	if err := writeConfigFile(profile); err != nil {
		t.Fatal("Failed to write config file: ", err)
	}

	cfgFilePath, err := util.GetConfigFilePath(profile)
	if err != nil {
		t.Fatal("Failed to get config file path: ", err)
	}

	if _, err := os.Stat(cfgFilePath); os.IsNotExist(err) {
		t.Fatal("Expected config file did not find one")
	}

	os.Remove(cfgFilePath)
}
