package mastodon

import (
	"testing"
)

func TestServer(t *testing.T) {
	want := "https://mastodon.social"

	t.Run("Test set server", func(t *testing.T) {
		var s Secrets
		s.SetServer(want)
		got := s.server

		if got != want {
			t.Fatalf("expected %v, got %v", want, got)
		}
	})

	t.Run("Test Get server", func(t *testing.T) {
		s := Secrets{
			server: want,
		}
		got := s.GetServer()

		if got != want {
			t.Fatalf("expected %v, got %v", want, got)
		}
	})
}

func TestClientID(t *testing.T) {
	want := "clientIDValue"

	t.Run("Test SetClientID", func(t *testing.T) {
		var s Secrets
		s.SetClientID(want)
		got := s.clientID

		if got != want {
			t.Fatalf("expected %v, got %v", want, got)
		}
	})

	t.Run("Test GetClientID", func(t *testing.T) {
		s := Secrets{
			clientID: want,
		}
		got := s.GetClientID()

		if got != want {
			t.Fatalf("expected %v, got %v", want, got)
		}
	})
}

func TestClientSecret(t *testing.T) {
	want := "clientSecretValue"

	t.Run("Test SetClientSecret", func(t *testing.T) {
		var s Secrets
		s.SetClientSecret(want)
		got := s.clientSecret

		if got != want {
			t.Fatalf("expected %v, got %v", want, got)
		}
	})

	t.Run("Test GetClientSecret", func(t *testing.T) {
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

	t.Run("Test SetAccessToken", func(t *testing.T) {
		var s Secrets
		s.SetAccessToken(want)
		got := s.accessToken

		if got != want {
			t.Fatalf("expected %v, got %v", want, got)
		}
	})

	t.Run("test GetAccessToken", func(t *testing.T) {
		s := Secrets{
			accessToken: want,
		}

		got := s.GetAccessToken()

		if got != want {
			t.Fatalf("expected %v, got %v", want, got)
		}
	})
}
