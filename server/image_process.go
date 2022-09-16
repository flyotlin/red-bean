package server

import (
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	"gocv.io/x/gocv"
)

func OpenImageInCVMat(path string) *gocv.Mat {
	if !isValidExtension(path) {
		log.Infof("invalid extension with image [%v]", path)
		return nil
	}

	img := gocv.IMRead(path, gocv.IMReadUnchanged)
	if img.Empty() {
		log.Errorf("failed to read image [%v] by gocv", path)
		return nil
	}
	return &img
}

func OpenImageInFile(path string) *os.File {
	if !isValidExtension(path) {
		log.Infof("invalid extension with image [%v]", path)
		return nil
	}

	img, err := os.Open(path)
	if err != nil {
		log.Errorf("failed to open image [%v] by package os: [%v]", path, err.Error())
		return nil
	}
	return img
}

var VALID_EXTENSIONS = []string{".jpg", ".png"}

func isValidExtension(path string) bool {
	// Check for image format (for now, only .jpg, .png supported)
	ext := filepath.Ext(path)
	ext = strings.ToLower(ext)

	for _, validExt := range VALID_EXTENSIONS {
		if ext == validExt {
			return true
		}
	}
	return false
}
