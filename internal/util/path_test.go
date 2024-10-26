package util

import (
	"os"
	"path/filepath"
	"testing"
	// "fmt"
)

func TestGetConfigDir(t *testing.T) {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		t.Fatalf("Failed to get user config dir: %v", err)

	}

	want := filepath.Join(cfgDir, "megophone")
	got, err := GetConfigDir()
	if err != nil {
		t.Fatalf("Got error didn't expect one: %v", err)
	}

	if want != got {
		t.Fatalf("Wrong config dir: want %v, got %v", want, got)
	}
}

func TestGetconfigFileName(t *testing.T) {
	want := "default.env"
	got := GetConfigFileName("default")

	if want != got {
		t.Fatalf("Wrong filename, want %s, got %s", want, got)
	}
}

func TestGetConfigFilePath(t *testing.T) {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		t.Fatalf("Failed to get user config dir: %v", err)

	}

	t.Run("Test default fileppath", func(t *testing.T) {
		want := filepath.Join(cfgDir, "megophone", "megophone-test.env")
		got, err := GetConfigFilePath("megophone-test")
		if err != nil {
			t.Fatalf("Got error didn't expect one: %v", err)
		}

		if want != got {
			t.Fatalf("Wrong file path: want %v, got %v", want, got)
		}
	})
}
