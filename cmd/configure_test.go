package cmd

import (
	"github.com/coolapso/megophone/internal/util"
	"os"
	"testing"
)

func TestWriteConfigFile(t *testing.T) {
	os.Setenv("GOLANG_TESTING", "true")
	defer os.Unsetenv("GOLANG_TESTING")

	if err := writeConfigFile(); err != nil {
		t.Fatal("Failed to write config file: ", err)
	}

	cfgFilePath, err := util.GetConfigFilePath()
	if err != nil {
		t.Fatal("Failed to get config file path: ", err)
	}

	if _, err := os.Stat(cfgFilePath); os.IsNotExist(err) {
		t.Fatal("Expected config file did not find one")
	}

	os.Remove(cfgFilePath)
}
