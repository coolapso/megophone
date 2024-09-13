package util

import ( 
	"testing"
	"os" 
	"path/filepath"
	// "fmt"
)

func TestGetConfigDir(t *testing.T) {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		t.Fatalf("Failed to get user config dir: %v", err)

	}

	want := filepath.Join(cfgDir, "xm-cli")

	got, err := GetConfigDir()
	if err != nil {
		t.Fatalf("Got error didn't expext one: %v", err)
	}

	if want != got {
		t.Fatalf("Wrong config dir: want %v, got %v", want, got)
	}
}


func TestGetConfigFilePath(t *testing.T) {

}
