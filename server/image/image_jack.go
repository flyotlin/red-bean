package image

import (
	"image"

	"github.com/rwcarlsen/goexif/exif"
	log "github.com/sirupsen/logrus"
	"gocv.io/x/gocv"
)

// TODO: 想一下每個物件實際的職責，Jack 應該專心照顧一個 Image 還是能照顧多個 Image?
// TODO: ImageJack 提供 Read/Write，用 interface 提供不同種類的行為 (cv.Mat, file)
type ImageJack struct {
	Path  string
	cvMat *gocv.Mat // Originally read from Path, won't be changed
	ioMgr *ImageIOMgr
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
		ioMgr: ioMgr,
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

// Now we convert BGR to HSV (V: 0~180), and set V to change brightness.
// **WARNING**: If V is over 180, the result would be weird.
func (j *ImageJack) SetBrightness(value float32) *gocv.Mat {
	// TODO: Check brightness value
	hsvImg := gocv.NewMat()
	defer hsvImg.Close()
	gocv.CvtColor(*j.cvMat, &hsvImg, gocv.ColorBGRToHSV)

	channels := gocv.Split(hsvImg)
	channels[2].AddFloat(value)

	gocv.Merge(channels, &hsvImg)
	setImg := gocv.NewMat()
	gocv.CvtColor(hsvImg, &setImg, gocv.ColorHSVToBGR)

	return &setImg
}

func (j *ImageJack) GetImageFileInfo() map[string]interface{} {
	infos := make(map[string]interface{})

	infos["Resolution"] = j.cvMat.Size()

	file := j.ioMgr.OpenInFile(j.Path)
	if file == nil {
		return nil
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		log.Errorf("failed to get stat of file <%v>: [%v]", j.Path, err.Error())
		return infos
	}

	infos["FileName"] = stat.Name()
	infos["FileSize"] = stat.Size()
	infos["FileModTime"] = stat.ModTime()

	return infos
}

func (j *ImageJack) GetExif() map[string]string {
	file := j.ioMgr.OpenInFile(j.Path)
	if file == nil {
		return nil
	}
	defer file.Close()

	x, err := exif.Decode(file)
	if err != nil {
		log.Errorf("failed to decode exif info from file [%v]: [%v]", j.Path, err.Error())
		return nil
	}

	tags := make(map[string]string)
	tagNames := []string{"Make", "Model", "LensModel", "DateTimeDigitized", "DateTime", "DateTimeOriginal"}
	for _, name := range tagNames {
		tags[name] = getExifTag(x, exif.FieldName(name))
	}

	return tags
}
