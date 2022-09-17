package server

import (
	"testing"
)

func TestIsValidExtensionTrue(t *testing.T) {
	paths := []string{
		"/home/nick/abc.png",
		"/home/nick/abc.JPG",
	}

	for _, path := range paths {
		isValid := isImgExtensionValid(path)
		if !isValid {
			t.Error("This should be valid")
		}
	}
}

func TestIsValidExtensionFalse(t *testing.T) {
	paths := []string{
		"/home/nick/def.ARW",
	}

	for _, path := range paths {
		isValid := isImgExtensionValid(path)
		if isValid {
			t.Error("This should be invalid")
		}
	}
}
