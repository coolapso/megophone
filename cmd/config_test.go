package cmd

import (
	"os"
	"testing"
	// "fmt"
	"strings"
	"bufio"
	// "github.com/spf13/cobra"
	// "github.com/spf13/viper"
)

func TestMaskString(t *testing.T) {
	t.Run("test 3 char string", func(t *testing.T) {
		want := "bar"
		got := maskString("bar")
		if got != want { 
			t.Fatalf("Strings do not match, want %v, got %v", want, got)
		}
	})

	t.Run("test masked string", func(t *testing.T) {
		want := "******baz"
		got := maskString("foobarbaz")
		if got != want { 
			t.Fatalf("Strings do not match, want %v, got %v", want, got)
		}

	})
}

func TestLoadXVars(t *testing.T) {
	want := config{
		xUser: "foo",
		xApiKey: "apikeystring",
	}

	os.Setenv("X_USER", want.xUser)
	os.Setenv("X_API_KEY", want.xApiKey)
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

	os.Setenv("MASTODON_USER", want.mUser)
	os.Setenv("MASTODON_API_KEY", want.mApiKey)

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
