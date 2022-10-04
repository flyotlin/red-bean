package image

import (
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
	"gocv.io/x/gocv"
)

func newFakeImage(width int, height int) *gocv.Mat {
	mat := gocv.NewMatWithSize(height, width, gocv.MatTypeCV16S)
	return &mat
}

func TestResize(t *testing.T) {
	w, h := 1200, 800
	resizedW, resizedH := 500, 300

	mat := newFakeImage(w, h)
	jack := ImageJack{cvMat: mat}

	resizedMat := jack.Resize(resizedW, resizedH)

	assert.Equal(t, resizedW, resizedMat.Size()[1])
	assert.Equal(t, resizedH, resizedMat.Size()[0])
}

func TestResizeByFactor(t *testing.T) {
	w, h := 1200, 800
	fw, fh := 0.6, 0.4

	mat := newFakeImage(w, h)
	jack := ImageJack{cvMat: mat}

	resizedMat := jack.ResizeByFactor(fw, fh)

	assert.Equal(t, float64(w)*fw, float64(resizedMat.Size()[1]))
	assert.Equal(t, float64(h)*fh, float64(resizedMat.Size()[0]))
}

func TestCropReturnCorrectSize(t *testing.T) {
	w, h := 640, 480
	head := image.Point{50, 60}
	tail := image.Point{100, 120}

	mat := newFakeImage(w, h)
	jack := ImageJack{cvMat: mat}

	croppedImg := jack.Crop(head, tail)
	assert.Equal(t, tail.X-head.X, croppedImg.Size()[1])
	assert.Equal(t, tail.Y-head.Y, croppedImg.Size()[0])
}

// 搞錯 w,h -> row,col -> .Size() 0 1，透過 testing + debugger 快速排除
func TestCropReturnCorrectSizeWhenOutOfBound(t *testing.T) {
	w, h := 640, 480
	heads := []image.Point{
		{-100, 0},
		{4, -100},
		{-20, -30},
	}
	tails := []image.Point{
		{50, h + 20},
		{w + 10, 50},
		{w + 15, h + 30},
	}
	expected := []image.Point{
		{50, 480},  // 0,0->50,480
		{636, 50},  // 4,0->640,50
		{640, 480}, // 0,0->640,480
	}

	mat := newFakeImage(w, h)
	jack := ImageJack{cvMat: mat}

	if len(heads) != len(tails) {
		t.Log("the number of elements in heads and tails must be the same")
		t.FailNow()
	}
	lens := len(heads)

	for i := 0; i < lens; i++ {
		croppedImg := jack.Crop(heads[i], tails[i])
		assert.Equal(t, expected[i].X, croppedImg.Size()[1])
		assert.Equal(t, expected[i].Y, croppedImg.Size()[0])
	}
}

func TestConvertColorSpace(t *testing.T) {
	jack := NewImageJack("/home/flyotlin/Documents/Screen Shot 2022-09-01 at 9.48.55 PM.png")
	ioMgr := NewImageIOMgr()

	ioMgr.WriteToFile("/home/flyotlin/Documents/Program/redbean/server/orig.png", jack.cvMat)

	cvtImg := jack.ConvertColorSpace(gocv.ColorBGRToGray)

	ioMgr.WriteToFile("/home/flyotlin/Documents/Program/redbean/server/cvt.png", cvtImg)
}

func TestSetBrightness(t *testing.T) {
	jack := NewImageJack("/home/flyotlin/Documents/Screen Shot 2022-09-01 at 9.48.55 PM.png")
	ioMgr := NewImageIOMgr()

	ioMgr.WriteToFile("/home/flyotlin/Documents/Program/redbean/server/orig.png", jack.cvMat)

	cvtImg := jack.SetBrightness(20)

	ioMgr.WriteToFile("/home/flyotlin/Documents/Program/redbean/server/cvt.png", cvtImg)
}

func TestGetImageFileInfo(t *testing.T) {
	jack := NewImageJack("/mnt/wd1/sony-a6000/2022/2022-08-20/DSC01610.JPG")
	infos := jack.GetImageFileInfo()
	t.Log(infos)
	t.Fail()
}

func TestGetExif(t *testing.T) {
	jack := NewImageJack("/mnt/wd1/sony-a6000/2022/2022-08-20/DSC01610.JPG")
	tags := jack.GetExif()
	t.Log(tags)
	t.Fail()
}
