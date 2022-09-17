package server

import (
	"os"
	"path/filepath"
	"strings"
)

var VALID_IMG_EXTENSIONS = []string{".jpg", ".png"}

func isImgExtensionValid(path string) bool {
	// Check for image format (for now, only .jpg, .png supported)
	ext := filepath.Ext(path)
	ext = strings.ToLower(ext)

	for _, validExt := range VALID_IMG_EXTENSIONS {
		if ext == validExt {
			return true
		}
	}
	return false
}
