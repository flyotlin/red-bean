package server

import (
	"image"
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

func isValidRect(head image.Point, tail image.Point) bool {
	if tail.X < head.X {
		return false
	}
	if tail.Y < head.Y {
		return false
	}
	return true
}

func clipAbove(p image.Point, clipP image.Point) image.Point {
	return clipPoint(p, clipP, true)
}

func clipLower(p image.Point, clipP image.Point) image.Point {
	return clipPoint(p, clipP, false)
}

func clipPoint(p image.Point, clipP image.Point, isAbove bool) image.Point {
	if isAbove {
		if p.X < clipP.X {
			p.X = clipP.X
		}
		if p.Y < clipP.Y {
			p.Y = clipP.Y
		}
	} else {
		if p.X > clipP.X {
			p.X = clipP.X
		}
		if p.Y > clipP.Y {
			p.Y = clipP.Y
		}
	}
	return p
}
