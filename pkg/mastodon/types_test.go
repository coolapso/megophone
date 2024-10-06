package mastodon

import (
	"testing"
)

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
