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
