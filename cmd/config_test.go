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
		x: xdotcom{
			oauthToken: "apikey",
			oauthTokenSecret: "apikeysecretstring",
		},
	}

	os.Setenv("XM_X_OAUTH_TOKEN", want.x.oauthToken)
	os.Setenv("XM_X_OAUTH_TOKEN_SECRET", want.x.oauthTokenSecret)
	defer os.Unsetenv("XM_X_OAUTH_TOKEN")
	defer os.Unsetenv("XM_X_OAUTH_TOKEN_SECRET")

	var got config
	loadXVars(&got)
	if got.x.oauthToken != want.x.oauthToken { 
		t.Fatalf("Api key does not match expected value: want %v, got %v", want.x.oauthToken, got.x.oauthToken)
	}

	if got.x.oauthTokenSecret != want.x.oauthTokenSecret { 
		t.Fatalf("Api key Secret does not math expected value: want %v, got %v", want.x.oauthTokenSecret, got.x.oauthTokenSecret)
	}
}

func TestLoadMastodonVars(t *testing.T) {
	want := config{
		m: mastodon{
			apiKey: "apikey",
			apiKeySecret: "apikeysecretstring",
		},
	}

	os.Setenv("XM_MASTODON_API_KEY", want.m.apiKey)
	os.Setenv("XM_MASTODON_API_KEY_SECRET", want.m.apiKeySecret)
	defer os.Unsetenv("XM_MASTODON_API_KEY")
	defer os.Unsetenv("XM_MASTODON_API_KEY_SECRET")

	var got config
	loadMastodonVars(&got)
	if got.m.apiKey != want.m.apiKey { 
		t.Fatalf("Api key does not match expected value: want %v, got %v", want.m.apiKey, got.m.apiKey)
	}

	if got.m.apiKeySecret != want.m.apiKeySecret { 
		t.Fatalf("Api Key does not matc expected value: want %v, got %v", want.m.apiKeySecret, got.m.apiKeySecret)
	}
}

func TestConfigXOauthToken(t *testing.T) {
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
			c.x.configOauthToken(r)
			if c.x.oauthToken != tt.expected {
				t.Errorf("input and saved value do not match: expected %v, got %v", tt.expected, c.x.oauthToken)
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
		{"Valid input", "foo\n", "foo"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var c config
			r := bufio.NewReader(strings.NewReader(tt.input))
			c.m.configApiKey(r)
			if c.m.apiKey != tt.expected {
				t.Fatalf("input and saved value do not match: expected %v, got %v", tt.expected, c.m.apiKey)
			}
		})
	}
}


func TestConfigXOauthTokenSecret(t *testing.T) {
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
			c.x.configOauthTokenSecret(r)
			if c.x.oauthTokenSecret != tt.expected {
				t.Errorf("input and saved value do not match: expected %v, got %v", tt.expected, c.x.oauthTokenSecret)
			}
		})
	}
}

func TestConfigMastodonApiKeySecret(t *testing.T) {
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
			c.m.configApiKeySecret(r)
			if c.m.apiKeySecret != tt.expected {
				t.Errorf("input and saved value do not match: expected %v, got %v", tt.expected, c.m.apiKeySecret)
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
	want, err := os.ReadFile("../fixtures/xm-cli.env")
	if err != nil {
		t.Fatal("Failed to open example env file: ", err)
	}

	// Redirect stdout to null device to suppress output
	null, _ := os.Open(os.DevNull)
	defer null.Close()
	old := os.Stdout
	os.Stdout = null
	defer func() { os.Stdout = old }()
	os.Setenv("GOLANG_TESTING", "true")
	defer os.Unsetenv("GOLANG_TESTING")

	cfgFilePath, err := util.GetConfigFilePath()
	if err != nil { 
		t.Fatal("Failed to get configuration file path: ", err)
	}

	t.Run("test user intput", func(t *testing.T) {
		os.Remove(cfgFilePath)

		input := "xoauthToken\nxoauthTokenSecret\nmapikey\nmapikeysecretstring\n"
		reader := bufio.NewReader(strings.NewReader(input))

		if err := configxm(reader); err != nil {
			t.Fatal("got error didn't expect one: ", err)
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
	})

	t.Run("test env vars", func(t *testing.T) {
		input := "\n\n\n\n"
		reader := bufio.NewReader(strings.NewReader(input))
		os.Setenv("XM_X_OAUTH_TOKEN", "xoauthToken")
		os.Setenv("XM_MASTODON_API_KEY", "mapikey")
		os.Setenv("XM_X_OAUTH_TOKEN_SECRET", "xoauthTokenSecret")
		os.Setenv("XM_MASTODON_API_KEY_SECRET", "mapikeysecretstring")

		defer os.Unsetenv("XM_X_OAUTH_TOKEN")
		defer os.Unsetenv("XM_MASTODON_API_KEY")
		defer os.Unsetenv("XM_X_OAUTH_TOKEN_SECRET")
		defer os.Unsetenv("XM_MASTODON_API_KEY_SECRET")

		if err := configxm(reader); err != nil { 
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
	})
}
