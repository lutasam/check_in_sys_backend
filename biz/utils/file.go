package utils

import (
	"github.com/lutasam/check_in_sys/biz/common"
	"mime/multipart"
	"strings"
)

func IsCorrectImg(header *multipart.FileHeader) (bool, error) {
	fileType := strings.Split(header.Filename, ".")[1]
	if fileType != "png" && fileType != "jpeg" && fileType != "jpg" {
		return false, common.IMGFORMATERROR
	}
	if header.Size > common.MAXIMGSPACE {
		return false, common.IMGTOOLARGEERROR
	}
	return true, nil
}
