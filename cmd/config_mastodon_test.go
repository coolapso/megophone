package cmd

import (
	"bufio"
	"fmt"
	gomasto "github.com/mattn/go-mastodon"
	"github.com/spf13/viper"
	"os"
	"strings"
	"testing"
)

func TestConfigMastodonServer(t *testing.T) {
	null, _ := os.Open(os.DevNull)
	defer null.Close()
	old := os.Stdout
	os.Stdout = null
	defer func() { os.Stdout = old }()

	wantMastodonServer := "https://mastodon.social"

	var got config
	input := fmt.Sprintf("%s\n", wantMastodonServer)
	reader := bufio.NewReader(strings.NewReader(input))
	configMastodonServer(reader, &got)

	if wantMastodonServer != got.m.GetServer() {
		t.Fatalf("Access Token does not match expected value. Want %v, got %v", wantMastodonServer, got.m.GetServer())
	}
}

func TestConfigMastodonAccessToken(t *testing.T) {
	null, _ := os.Open(os.DevNull)
	defer null.Close()
	old := os.Stdout
	os.Stdout = null
	defer func() { os.Stdout = old }()

	wantAuthCode := "MastodonUserAuthorizationCode"

	var app gomasto.Application
	input := fmt.Sprintf("%s\n", wantAuthCode)
	reader := bufio.NewReader(strings.NewReader(input))
	got, err := getMastodonUserAuthorizationCode(reader, &app)
	if err != nil {
		t.Fatal("Failed to configure mastodon, got error, did not expect one")
	}

	if wantAuthCode != got {
		t.Fatalf("Authorization code does not match expected value. Want %v, got %v", wantAuthCode, got)
	}
}

func TestMastodonClientConfig(t *testing.T) {
	configNotMatchMessage := "Configuration does not match."
	server := "https://mastodon.social"
	clientId := "MastodonClientID"
	clientSecret := "MastodonClientSecret"
	accessToken := "MastodonAccessToken"

	want := &gomasto.Config{
		Server:       server,
		ClientID:     clientId,
		ClientSecret: clientSecret,
		AccessToken:  accessToken,
	}

	viper.Set("mastodon_server", server)
	viper.Set("mastodon_client_id", clientId)
	viper.Set("mastodon_client_secret", clientSecret)
	viper.Set("mastodon_access_token", accessToken)

	got := mastodonClientConfig()

	if got.Server != want.Server {
		t.Fatalf("%v Wrong server, want %v, got %v", configNotMatchMessage, want.Server, got.Server)
	}

	if got.ClientID != want.ClientID {
		t.Fatalf("%v Wrong ClientID, want %v, got %v", configNotMatchMessage, want.ClientID, got.ClientID)
	}

	if got.ClientSecret != want.ClientSecret {
		t.Fatalf("%v Wrong ClientSecret, want %v, got %v", configNotMatchMessage, want.ClientSecret, got.ClientSecret)
	}

	if got.AccessToken != want.AccessToken {
		t.Fatalf("%v Wrong AccessToken, want %v, got %v", configNotMatchMessage, want.AccessToken, got.AccessToken)
	}
}
