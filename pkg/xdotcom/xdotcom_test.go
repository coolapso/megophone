package xdotcom

import (
	// "os"
	"testing"
	// "fmt"
	// "bufio"
	// "strings"
	// "github.com/coolapso/megophone/internal/util"
)

// func TesSetXOauthToken(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		input    string
// 		expected string
// 	}{
// 		{"Empty input", "\n", ""},
// 		{"Valid input", "foo\n", "foo"},
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			var s Secrets
// 			// r := bufio.NewReader(strings.NewReader(tt.input))
// 			s.SetOauthToken(tt.input)
// 			if s.oauthToken != tt.expected {
// 				t.Errorf("input and saved value do not match: expected %v, got %v", tt.expected, s.oauthToken)
// 			}
// 		})
// 	}
// }
//
// func TestSetXOauthTokenSecret(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		input    string
// 		expected string
// 	}{
// 		{"Empty input", "\n", ""},
// 		{"Valid input", "oauthTokenSecret\n", "oauthTokenSecret"},
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			var s Secrets
// 			r := bufio.NewReader(strings.NewReader(tt.input))
// 			s.SetOauthTokenSecret(r)
// 			if s.oauthTokenSecret != tt.expected {
// 				t.Errorf("input and saved value do not match: expected %v, got %v", tt.expected, s.oauthTokenSecret)
// 			}
// 		})
// 	}
// }
//
// func TestSetXAPIKey(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		input    string
// 		expected string
// 	}{
// 		{"Empty input", "\n", ""},
// 		{"Valid input", "foo\n", "foo"},
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			var s Secrets
// 			r := bufio.NewReader(strings.NewReader(tt.input))
// 			s.SetApiKey(r)
// 			if s.apiKey != tt.expected {
// 				t.Errorf("input and saved value do not match: expected %v, got %v", tt.expected, s.apiKey)
// 			}
// 		})
// 	}
// }
//
// func TestSetXAPIKeySecret(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		input    string
// 		expected string
// 	}{
// 		{"Empty input", "\n", ""},
// 		{"Valid input", "someApiKeySecret\n", "someApiKeySecret"},
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			var s Secrets
// 			r := bufio.NewReader(strings.NewReader(tt.input))
// 			s.SetApiKeySecret(r)
// 			if s.apiKeySecret != tt.expected {
// 				t.Errorf("input and saved value do not match: expected %v, got %v", tt.expected, s.apiKeySecret)
// 			}
// 		})
// 	}
// }

func TestSetOauthToken(t *testing.T) {
	want := "oauthTokenValue"
	var s Secrets
	s.SetOauthToken(want)
	got := s.oauthToken

	if got != want {
		t.Fatalf("expected %v, got %v", want, got)
	}
}

func TestSetOauthTokenSecret(t *testing.T) {
	want := "oauthTokenSecretValue"
	var s Secrets
	s.SetOauthTokenSecret(want)
	got := s.oauthTokenSecret

	if got != want {
		t.Fatalf("expected %v, got %v", want, got)
	}
}

func TestSetApiKey(t *testing.T) {
	want := "apiKeyValue"
	var s Secrets
	s.SetApiKey(want)
	got := s.apiKey

	if got != want {
		t.Fatalf("expected %v, got %v", want, got)
	}
}

func TestSetApiKeySecret(t *testing.T) {
	want := "apiKeySecretValue"
	var s Secrets
	s.SetApiKeySecret(want)
	got := s.apiKeySecret

	if got != want {
		t.Fatalf("expected %v, got %v", want, got)
	}
}

func TestGetOauthToken(t *testing.T) {
	want := "oauthTokenValue"
	s := Secrets{
		oauthToken: want,
	}
	got := s.GetOauthToken()

	if got != want {
		t.Fatalf("expected %v, got %v", want, got)
	}

}

func TestGetOauthTokenSecret(t *testing.T) {
	want := "oauthTokenSecretValue"
	s := Secrets{
		oauthTokenSecret: want,
	}
	got := s.GetOauthTokenSecret()

	if got != want {
		t.Fatalf("expected %v, got %v", want, got)
	}
}

func TestGetApiKey(t *testing.T) {
	want := "apiKeyValue"
	s := Secrets{
		apiKey: want,
	}
	got := s.GetApiKey()

	if got != want {
		t.Fatalf("expected %v, got %v", want, got)
	}
}

func TestGetApiKeySecret(t *testing.T) {
	want := "apiKeySecretValue"
	s := Secrets{
		apiKeySecret: want,
	}
	got := s.GetApiKeySecret()

	if got != want {
		t.Fatalf("expected %v, got %v", want, got)
	}
}
