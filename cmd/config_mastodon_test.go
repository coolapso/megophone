package cmd

import (
	"fmt"
	"os"
	"testing"
	"bufio"
	"strings"
	gomasto "github.com/mattn/go-mastodon"
)

func TestConfigMastodonServer(t *testing.T) {
	null, _ := os.Open(os.DevNull)
	defer null.Close()
	old := os.Stdout
	os.Stdout = null
	defer func() { os.Stdout = old }()

	wantMastodonServer := "https://mastodon.social"

	var got config
	input := fmt.Sprintf("%s\n",wantMastodonServer)
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

	wantMastodonAccessToken := "mastodonAccessToken"
	
	var got config
	var app gomasto.Application
	input := fmt.Sprintf("%s\n", wantMastodonAccessToken)
	reader := bufio.NewReader(strings.NewReader(input))
	if err := configMastodonAccessToken(reader, &app, &got); err != nil {
		t.Fatalf("Failed to configure mastodon, got error, did not expect one")
	}

	if wantMastodonAccessToken != got.m.GetAccessToken() {
		t.Fatalf("Access Token does not match expected value. Want %v, got %v", wantMastodonAccessToken, got.m.GetAccessToken())
	}
}
