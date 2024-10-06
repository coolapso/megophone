package mastodon

import (
	"testing"
)

func TestClientKey(t *testing.T) {
	want := "clientKeyValue"

	t.Run("Test SetClientKey", func (t *testing.T) {
		var s Secrets
		s.SetClientKey(want)
		got := s.clientKey

		if got != want {
			t.Fatalf("expected %v, got %v", want, got)
		}
	})

	t.Run("Test GetClientKey", func (t *testing.T) {
		s := Secrets{
			clientKey: want,
		}
		got := s.GetClientKey()

		if got != want {
			t.Fatalf("expected %v, got %v", want, got)
		}
	})
}

func TestClientSecret(t *testing.T) {
	want := "clientSecretValue"

	t.Run("Test SetClientSecret", func (t *testing.T) {
		var s Secrets
		s.SetClientSecret(want)
		got := s.clientSecret

		if got != want {
			t.Fatalf("expected %v, got %v", want, got)
		}
	})

	t.Run("Test GetClientSecret", func (t *testing.T) {
		s := Secrets{
			clientSecret: want,
		}
		got := s.GetClientSecret()

		if got != want {
			t.Fatalf("expected %v, got %v", want, got)
		}
	})
}

func TestAccessToken(t *testing.T) {
	want := "accessTokenValue"

	t.Run("Test SetAccessToken", func (t *testing.T) {
		var s Secrets
		s.SetAccessToken(want)
		got := s.accessToken

		if got != want {
			t.Fatalf("expected %v, got %v", want, got)
		}
	})

	t.Run("test GetAccessToken", func (t *testing.T) {
		s := Secrets{
			accessToken: want,
		}

		got := s.GetAccessToken()

		if got != want {
			t.Fatalf("expected %v, got %v", want, got)
		}
	})
}
