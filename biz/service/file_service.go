package service

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/check_in_sys/biz/bo"
	"github.com/lutasam/check_in_sys/biz/utils"
	"strings"
	"sync"
)

type FileService struct{}

var (
	fileService     *FileService
	fileServiceOnce sync.Once
)

func GetFileService() *FileService {
	fileServiceOnce.Do(func() {
		fileService = &FileService{}
	})
	return fileService
}

func (ins *FileService) UploadImage(c *gin.Context, req *bo.UploadImageRequest) (*bo.UploadImageResponse, error) {
	isCorrect, err := utils.IsCorrectImg(req.FileHeader)
	if !isCorrect {
		return nil, err
	}
	newFilename := utils.Uint64ToString(utils.GenerateFileID()) + "." + strings.Split(req.FileHeader.Filename, ".")[1]
	url := utils.GetConfigString("file.imgs.url") + newFilename
	err = c.SaveUploadedFile(req.FileHeader, utils.GetConfigString("file.imgs.address")+newFilename)
	if err != nil {
		return nil, err
	}
	return &bo.UploadImageResponse{ImageURL: url}, nil
}
