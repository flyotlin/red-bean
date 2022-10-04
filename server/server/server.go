package server

import (
	"context"

	"github.com/flyotlin/red-bean/image"
	pb "github.com/flyotlin/red-bean/proto"
)

type Server struct {
	pb.UnimplementedImageServiceServer
}

func (s *Server) GetImageInfo(ctx context.Context, req *pb.ImageInfoRequest) (*pb.ImageInfoResponse, error) {
	// Find image on server
	uid := req.GetUid()
	archiver := image.NewImageArchiver()
	path := archiver.FindImageByUID(uid.String())

	// Open image by path using imageJack
	jack := image.NewImageJack(path)
	info := jack.GetImageFileInfo()

	// Prepare response value
	dim := GetDimension(info["Resolution"])
	filename := GetFileName(info["FileName"])

	return &pb.ImageInfoResponse{Dim: dim, Filename: filename}, nil
}

func GetDimension(res interface{}) *pb.Dimension {
	resolution := res.([]int)
	w, h := resolution[1], resolution[0]

	return &pb.Dimension{
		Width:  int32(w),
		Height: int32(h),
	}
}

func GetFileName(it interface{}) string {
	filename := it.(string)
	return filename
}
