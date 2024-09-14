package cmd

import (
	"os"
	"testing"

	"fmt"
	"bufio"
	"strings"

	"github.com/coolapso/xm-cli/internal/util"
)

func TestLoadXVars(t *testing.T) {
	want := config{
		xUser: "foo",
		xApiKey: "apikeystring",
	}

	os.Setenv("XM_X_USER", want.xUser)
	os.Setenv("XM_X_API_KEY", want.xApiKey)
	defer os.Unsetenv("X_USER")
	defer os.Unsetenv("X_API_KEY")

	var got config
	loadXVars(&got)
	if got.xUser != want.xUser { 
		t.Fatalf("X Username does not match expected value: want %v, got %v", want.xUser, got.xUser)
	}

	if got.xApiKey != want.xApiKey { 
		t.Fatalf("X Api Key does not math expected value: want %v, got %v", want.xApiKey, got.xApiKey)
	}
}

func TestLoadMastodonVars(t *testing.T) {
	want := config{
		mUser: "foo",
		mApiKey: "apikeystring",
	}

	os.Setenv("XM_MASTODON_USER", want.mUser)
	os.Setenv("XM_MASTODON_API_KEY", want.mApiKey)

	var got config
	loadMastodonVars(&got)
	if got.mUser != want.mUser { 
		t.Fatalf("Mastodon Username does not match expected value: want %v, got %v", want.mUser, got.mUser)
	}

	if got.mApiKey != want.mApiKey { 
		t.Fatalf("X Api Key does not matc expected value: want %v, got %v", want.mApiKey, got.mApiKey)
	}
}

func TestConfigXUser(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Empty input", "\n", ""},
		{"Valid input", "foo\n", "foo"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var c config
			r := bufio.NewReader(strings.NewReader(tt.input))
			c.configXUser(r)
			if c.xUser != tt.expected {
				t.Errorf("input and saved value do not match: expected %v, got %v", tt.expected, c.xUser)
			}
		})
	}
}

func TestConfigMastodonUser(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Empty input", "\n", ""},
		{"Valid input", "foo\n", "foo"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var c config
			r := bufio.NewReader(strings.NewReader(tt.input))
			c.configMastodonUser(r)
			if c.mUser != tt.expected {
				t.Errorf("input and saved value do not match: expected %v, got %v", tt.expected, c.xUser)
			}
		})
	}
}


func TestConfigXApiKey(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Empty input", "\n", ""},
		{"Valid input", "someApiKey\n", "someApiKey"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var c config
			r := bufio.NewReader(strings.NewReader(tt.input))
			c.configXApiKey(r)
			if c.xApiKey != tt.expected {
				t.Errorf("input and saved value do not match: expected %v, got %v", tt.expected, c.xApiKey)
			}
		})
	}
}

func TestConfigMastodonApiKey(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Empty input", "\n", ""},
		{"Valid input", "someApiKey\n", "someApiKey"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var c config
			r := bufio.NewReader(strings.NewReader(tt.input))
			c.configMastodonApiKey(r)
			if c.mApiKey != tt.expected {
				t.Errorf("input and saved value do not match: expected %v, got %v", tt.expected, c.xApiKey)
			}
		})
	}
}

func TestWriteConfigFile(t *testing.T) {
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


func TestConfigxm(t *testing.T) {
	os.Setenv("GOLANG_TESTING", "true")
	os.Setenv("XM_X_USER", "foo")
	os.Setenv("XM_MASTODON_USER", "bar")
	os.Setenv("XM_X_API_KEY", "somexapikey")
	os.Setenv("XM_MASTODON_API_KEY", "somemastodonapikey")
	
	want, err := os.ReadFile("../fixtures/xm-cli.env")
	if err != nil {
		t.Fatal("Failed to open example env file: ", err)
	}

	cfgFilePath, err := util.GetConfigFilePath()
	if err != nil { 
		t.Fatal("Failed to get configuration file path: ", err)
	}

	if err := configxm(); err != nil { 
		t.Fatal("Got error didn't expect one: ", err)
	}
	defer os.Remove(cfgFilePath)

	got, err := os.ReadFile(cfgFilePath)
	if err != nil {
		t.Fatal("Failed to read test configuration file")
	}
	fmt.Println(string(cfgFilePath))


	if string(want) != string(got) { 
		t.Fatalf("Configuration file does not match, want:\n%v\ngot\n%v", string(want), string(got))
	}
}
