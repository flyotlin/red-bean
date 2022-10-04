package image

import (
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"gocv.io/x/gocv"
)

type ImageIOMgr struct {
}

func NewImageIOMgr() *ImageIOMgr {
	return &ImageIOMgr{}
}

func (mgr *ImageIOMgr) OpenInCvMat(path string) *gocv.Mat {
	if !isImgExtensionValid(path) {
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

func (mgr *ImageIOMgr) OpenInFile(path string) *os.File {
	if !isImgExtensionValid(path) {
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

func (mgr *ImageIOMgr) WriteToFile(path string, mat *gocv.Mat) error {
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		msg := fmt.Sprintf("directory [%v] to write gocv mat not exist: [%v]", dir, err.Error())
		log.Errorf(msg)
		return fmt.Errorf(msg)
	}

	ok := gocv.IMWrite(path, *mat)
	if !ok {
		msg := fmt.Sprintf("gocv failed to write to [%v]", path)
		log.Errorf(msg)
		return fmt.Errorf(msg)
	}
	return nil
}
