package util

import (
	"io"
	"os"
	"testing"
)

func TestGetMediaType(t *testing.T) {
	filePath := "../../media/megophone.png"
	want := "image"
	got := GetMediaType(filePath)

	if got != want {
		t.Fatalf("Got wront file type, expected %v, got %v", want, got)
	}
}

func TestOpenMediaFile(t *testing.T) {
	filePath := "../../media/megophone.png"
	file, err := os.Open(filePath)
	if err != nil {
		t.Fatal("Failed to open file:", err)
	}

	defer file.Close()

	want, err := io.ReadAll(file)
	if err != nil {
		t.Fatal("Failed to read file:", err)
	}

	got, err := OpenMediaFile(filePath)
	if err != nil {
		t.Fatal("Got error, didn't expect one:", err)
	}

	if string(got) != string(want) {
		t.Fatal("files do not match")
	}

}
