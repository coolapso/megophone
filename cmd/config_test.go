package cmd

import (
	"os"
	"testing"
	"bufio"
	"strings"
	"github.com/coolapso/megophone/internal/util"
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
	var want config
	want.m.SetClientKey("mastodonClientKey")
	want.m.SetClientSecret("mastodonClientSecret")
	want.m.SetAccessToken("mastodonAccessToken")

	os.Setenv("MEGOPHONE_MASTODON_CLIENT_KEY", want.m.GetClientKey())
	os.Setenv("MEGOPHONE_MASTODON_CLIENT_SECRET", want.m.GetClientSecret())
	os.Setenv("MEGOPHONE_MASTODON_ACCESS_TOKEN", want.m.GetAccessToken())
	defer os.Unsetenv("MEGOPHONE_MASTODON_CLIENT_KEY")
	defer os.Unsetenv("MEGOPHONE_MASTODON_CLIENT_SECRET")
	defer os.Unsetenv("MEGOPHONE_MASTODON_ACCESS_TOKEN")

	var got config
	loadMastodonVars(&got)
	if got.m.GetClientKey() != want.m.GetClientKey() { 
		t.Fatalf("Api key does not match expected value: want %v, got %v", want.m.GetClientKey(), got.m.GetClientKey())
	}

	if got.m.GetClientSecret() != want.m.GetClientSecret() { 
		t.Fatalf("Api Key does not match expected value: want %v, got %v", want.m.GetClientSecret(), got.m.GetClientSecret())
	}

	if got.m.GetAccessToken() != want.m.GetAccessToken() { 
		t.Fatalf("Access Token expected value: want %v, got %v", want.m.GetAccessToken(), got.m.GetAccessToken())
	}
}

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

func TestConfigMegophone(t *testing.T) {
	want, err := os.ReadFile("../fixtures/megophone.env")
	if err != nil {
		t.Fatal("Failed to open example env file: ", err)
	}

	os.Setenv("GOLANG_TESTING", "true")
	defer os.Unsetenv("GOLANG_TESTING")

	cfgFilePath, err := util.GetConfigFilePath()
	if err != nil { 
		t.Fatal("Failed to get configuration file path: ", err)
	}

	// Redirect stdout to null device to suppress output
	null, _ := os.Open(os.DevNull)
	defer null.Close()
	old := os.Stdout
	os.Stdout = null
	defer func() { os.Stdout = old }()

	t.Run("test user intput", func(t *testing.T) {
		os.Remove(cfgFilePath)

		input := "xoauthToken\nxoauthTokenSecret\nxapikey\nxapikeysecretstring\nmapikey\nmapikeysecretstring\nmastodonaccesstoken\n"
		reader := bufio.NewReader(strings.NewReader(input))

		if err := configMegophone(reader); err != nil {
			t.Fatal("got error didn't expect one: ", err)
		}
		defer os.Remove(cfgFilePath)

		got, err := os.ReadFile(cfgFilePath)
		if err != nil {
			t.Fatal("Failed to read test configuration file")
		}

		if string(want) != string(got) { 
			t.Fatalf("Configuration file does not match, want:\n%v\ngot\n%v", string(want), string(got))
		}
	})

	t.Run("test env vars", func(t *testing.T) {
		input := "\n\n\n\n"
		reader := bufio.NewReader(strings.NewReader(input))
		os.Setenv("MEGOPHONE_X_OAUTH_TOKEN", "xoauthToken")
		os.Setenv("MEGOPHONE_X_OAUTH_TOKEN_SECRET", "xoauthTokenSecret")
		os.Setenv("MEGOPHONE_MASTODON_CLIENT_KEY", "mapikey")
		os.Setenv("MEGOPHONE_MASTODON_CLIENT_SECRET", "mapikeysecretstring")
		os.Setenv("MEGOPHONE_MASTODON_ACCESS_TOKEN", "mastodonaccesstoken")

		defer os.Unsetenv("MEGOPHONE_X_OAUTH_TOKEN")
		defer os.Unsetenv("MEGOPHONE_MASTODON_CLIENT_KEY")
		defer os.Unsetenv("MEGOPHONE_X_OAUTH_TOKEN_SECRET")
		defer os.Unsetenv("MEGOPHONE_MASTODON_CLIENT_SECRET")
		defer os.Unsetenv("MEGOPHONE_MASTODON_ACCESS_TOKEN")

		if err := configMegophone(reader); err != nil { 
			t.Fatal("Got error didn't expect one: ", err)
		}
		defer os.Remove(cfgFilePath)

		got, err := os.ReadFile(cfgFilePath)
		if err != nil {
			t.Fatal("Failed to read test configuration file")
		}

		if string(want) != string(got) { 
			t.Fatalf("Configuration file does not match, want:\n%v\ngot\n%v", string(want), string(got))
		}
	})
}
