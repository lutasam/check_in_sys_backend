package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/check_in_sys/biz/bo"
	"github.com/lutasam/check_in_sys/biz/common"
	"github.com/lutasam/check_in_sys/biz/service"
	"github.com/lutasam/check_in_sys/biz/utils"
)

type FileController struct{}

func RegisterFileRouter(r *gin.RouterGroup) {
	fileController := &FileController{}
	{
		r.POST("/upload_image", fileController.UploadImage)
		r.POST("/upload_file", fileController.UploadFile)
		r.POST("/download_file", fileController.DownloadFile)
	}
}

func (ins *FileController) UploadImage(c *gin.Context) {
	req := &bo.UploadImageRequest{}
	_, header, err := c.Request.FormFile("image")
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	req.FileHeader = header
	resp, err := service.GetFileService().UploadImage(c, req)
	if err != nil {
		if utils.IsClientError(err) {
			utils.ResponseClientError(c, err.(common.Error))
			return
		} else {
			utils.ResponseServerError(c, err.(common.Error))
			return
		}
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *FileController) UploadFile(c *gin.Context) {
	req := &bo.UploadFileRequest{}
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	req.FileHeader = header
	resp, err := service.GetFileService().UploadFile(c, req)
	if err != nil {
		if utils.IsClientError(err) {
			utils.ResponseClientError(c, err.(common.Error))
			return
		} else {
			utils.ResponseServerError(c, err.(common.Error))
			return
		}
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *FileController) DownloadFile(c *gin.Context) {
	req := &bo.DownloadFileRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	_, err = service.GetFileService().DownloadFile(c, req)
	if err != nil {
		if utils.IsClientError(err) {
			utils.ResponseClientError(c, err.(common.Error))
			return
		} else {
			utils.ResponseServerError(c, err.(common.Error))
			return
		}
	}
	// utils.ResponseSuccess(c, resp) // download does not need respond success
}
