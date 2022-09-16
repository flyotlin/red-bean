package server

import (
	"image"

	log "github.com/sirupsen/logrus"
	"gocv.io/x/gocv"
)

type ImageJack struct {
	Path  string
	cvMat *gocv.Mat // Originally read from Path, won't be changed
}

func NewImageJack(path string) *ImageJack {
	mat := OpenImageInCVMat(path)
	if mat == nil {
		log.Error("failed to create ImageJack: [empty cv Mat]")
	}

	return &ImageJack{
		Path:  path,
		cvMat: mat,
	}
}

func (j *ImageJack) Resize(width int, height int) *gocv.Mat {
	resizedImg := gocv.NewMat()
	gocv.Resize(*j.cvMat, &resizedImg, image.Point{width, height}, 0, 0, gocv.InterpolationDefault)
	return &resizedImg
}
