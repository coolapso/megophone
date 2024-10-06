package util

import (
	"strings"
	"mime"
	"path/filepath"
	"os"
	"io"
)

func GetMediaType(mediaPath string) string {
	mimeType := mime.TypeByExtension(filepath.Ext(mediaPath))
	
	return strings.Split(mimeType, "/")[0]
}

func OpenMediaFile(mediaPath string) ([]byte, error) {
	file, err := os.Open(mediaPath)
	if err != nil { 
		return nil, err
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return fileBytes, nil

}
