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
	ioMgr := NewImageIOMgr() // TODO: 加到 jack 裡面
	mat := ioMgr.OpenInCvMat(path)
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

func (j *ImageJack) ResizeByFactor(fw float64, fh float64) *gocv.Mat {
	resizedImg := gocv.NewMat()
	gocv.Resize(*j.cvMat, &resizedImg, image.Point{}, fw, fh, gocv.InterpolationDefault)
	return &resizedImg
}

func (j *ImageJack) Crop(head image.Point, tail image.Point) *gocv.Mat {
	if !isValidRect(head, tail) {
		log.Errorf("invalid crop rectangle: [%v, %v]", head, tail)
		return nil
	}

	imgHead := image.Point{0, 0}
	imgTail := image.Point{j.cvMat.Size()[1], j.cvMat.Size()[0]}
	head = clipAbove(head, imgHead)
	tail = clipLower(tail, imgTail)

	cropRect := image.Rect(head.X, head.Y, tail.X, tail.Y)
	croppedImg := j.cvMat.Region(cropRect)
	return &croppedImg
}

func (j *ImageJack) ConvertColorSpace(code gocv.ColorConversionCode) *gocv.Mat {
	// TODO: Check if this image is convertable (at lease 3 channels or so)
	convertedImg := gocv.NewMat()
	gocv.CvtColor(*j.cvMat, &convertedImg, code)
	return &convertedImg
}
