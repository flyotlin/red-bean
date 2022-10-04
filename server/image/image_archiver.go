package image

type ImageArchiver struct {
}

func NewImageArchiver() *ImageArchiver {
	return &ImageArchiver{}
}

// TODO: maybe we can save uploaded images to /tmp (tmpfs, gets deleted after rebooting)
func (a *ImageArchiver) FindImageByUID(uid string) string {
	return "/home/flyotlin/Documents/gitlab.PNG"
}
