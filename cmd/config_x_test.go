package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
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

func TestConfigX(t *testing.T) {
	wantXOauthToken := "xoauthToken"
	wantXOauthTokenSecret := "wantXoauthTokenSecret"
	wantXApiKey := "xapikey"
	wantXApiKeySecret := "xapiKeysecretstring"

	null, _ := os.Open(os.DevNull)
	defer null.Close()
	old := os.Stdout
	os.Stdout = null
	defer func() { os.Stdout = old }()

	var got config
	input := fmt.Sprintf("%s\n%s\n%s\n%s\n", wantXOauthToken, wantXOauthTokenSecret, wantXApiKey, wantXApiKeySecret)
	reader := bufio.NewReader(strings.NewReader(input))
	configX(reader, &got)

	if wantXOauthToken != got.x.GetOauthToken() {
		t.Fatalf("Oauth token does not match expected value. Want %v, got %v", wantXOauthToken, got.x.GetOauthToken())
	}

	if wantXOauthTokenSecret != got.x.GetOauthTokenSecret() {
		t.Fatalf("Oauth token secret does not match expected value. Want %v, got %v", wantXOauthToken, got.x.GetOauthTokenSecret())
	}

	if wantXApiKey != got.x.GetApiKey() {
		t.Fatalf("X Api Key does not match expected value. Want %v, got %v", wantXApiKey, got.x.GetApiKey())
	}

	if wantXApiKeySecret != got.x.GetApiKeySecret() {
		t.Fatalf("X Api Key Secret does not match expected value. Want %v, got %v", wantXApiKey, got.x.GetApiKeySecret())
	}
}
