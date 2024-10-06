package cmd

import (
	"os"
	"testing"

	"fmt"
	"bufio"
	"strings"


	"github.com/coolapso/megophone/internal/util"
	// "github.com/coolapso/megophone/pkg/xdotcom"
)

func TestLoadXVars(t *testing.T) {
	var want config
	want.x.SetOauthToken("oauthToken")
	want.x.SetOauthTokenSecret("oauthTokenSecret")
	want.x.SetApiKey("xapikey")
	want.x.SetApiKeySecret("xapiKeysecretstring")

	os.Setenv("MEGOPHONE_X_OAUTH_TOKEN", want.x.GetOauthToken())
	os.Setenv("MEGOPHONE_X_OAUTH_TOKEN_SECRET", want.x.GetOauthTokenSecret())
	os.Setenv("MEGOPHONE_X_API_KEY", want.x.GetApiKey())
	os.Setenv("MEGOPHONE_X_API_KEY_SECRET", want.x.GetApiKeySecret())
	defer os.Unsetenv("MEGOPHONE_X_OAUTH_TOKEN")
	defer os.Unsetenv("MEGOPHONE_X_OAUTH_TOKEN_SECRET")
	defer os.Unsetenv("MEGOPHONE_X_API_KEY")
	defer os.Unsetenv("MEGOPHONE_X_API_KEY_SECRET")

	var got config
	loadXVars(&got)
	if got.x.GetOauthToken() != want.x.GetOauthToken() { 
		t.Fatalf("Oauth token does not match expected value: want %v, got %v", want.x.GetOauthToken(), got.x.GetOauthToken())
	}

	if got.x.GetOauthTokenSecret() != want.x.GetOauthTokenSecret() { 
		t.Fatalf("Oauth token Secret does not math expected value: want %v, got %v", want.x.GetOauthTokenSecret(), got.x.GetOauthTokenSecret())
	}

	if got.x.GetApiKey() != want.x.GetApiKey() { 
		t.Fatalf("Api key does not match expected value: want %v, got %v", want.x.GetApiKey(), got.x.GetApiKey())
	}

	if got.x.GetApiKeySecret() != want.x.GetApiKeySecret() {
		t.Fatalf("Api key Secret does not math expected value: want %v, got %v", want.x.GetApiKeySecret(), got.x.GetApiKeySecret())
	}
}

func TestLoadMastodonVars(t *testing.T) {
	want := config{
		m: mastodon{
			apiKey: "apikey",
			apiKeySecret: "apikeysecretstring",
		},
	}

	os.Setenv("MEGOPHONE_MASTODON_API_KEY", want.m.apiKey)
	os.Setenv("MEGOPHONE_MASTODON_API_KEY_SECRET", want.m.apiKeySecret)
	defer os.Unsetenv("MEGOPHONE_MASTODON_API_KEY")
	defer os.Unsetenv("MEGOPHONE_MASTODON_API_KEY_SECRET")

	var got config
	loadMastodonVars(&got)
	if got.m.apiKey != want.m.apiKey { 
		t.Fatalf("Api key does not match expected value: want %v, got %v", want.m.apiKey, got.m.apiKey)
	}

	if got.m.apiKeySecret != want.m.apiKeySecret { 
		t.Fatalf("Api Key does not matc expected value: want %v, got %v", want.m.apiKeySecret, got.m.apiKeySecret)
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


func TestConfigMegophone(t *testing.T) {
	want, err := os.ReadFile("../fixtures/megophone.env")
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

		input := "xoauthToken\nxoauthTokenSecret\nxapikey\nxapikeysecretstring\nmapikey\nmapikeysecretstring\n"
		reader := bufio.NewReader(strings.NewReader(input))

		if err := configMegophone(reader); err != nil {
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
		os.Setenv("MEGOPHONE_X_OAUTH_TOKEN", "xoauthToken")
		os.Setenv("MEGOPHONE_X_OAUTH_TOKEN_SECRET", "xoauthTokenSecret")
		os.Setenv("MEGOPHONE_MASTODON_API_KEY", "mapikey")
		os.Setenv("MEGOPHONE_MASTODON_API_KEY_SECRET", "mapikeysecretstring")

		defer os.Unsetenv("MEGOPHONE_X_OAUTH_TOKEN")
		defer os.Unsetenv("MEGOPHONE_MASTODON_API_KEY")
		defer os.Unsetenv("MEGOPHONE_X_OAUTH_TOKEN_SECRET")
		defer os.Unsetenv("MEGOPHONE_MASTODON_API_KEY_SECRET")

		if err := configMegophone(reader); err != nil { 
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
