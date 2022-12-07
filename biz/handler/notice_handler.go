package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/check_in_sys/biz/bo"
	"github.com/lutasam/check_in_sys/biz/common"
	"github.com/lutasam/check_in_sys/biz/service"
	"github.com/lutasam/check_in_sys/biz/utils"
)

type NoticeController struct{}

func RegisterNoticeRouter(r *gin.RouterGroup) {
	noticeController := &NoticeController{}
	{
		r.POST("/find_all_notices", noticeController.FindAllNotices)
		r.POST("/upload_notice", noticeController.UploadNotice)
		r.POST("/delete_notice", noticeController.DeleteNotice)
	}
}

func (ins *NoticeController) FindAllNotices(c *gin.Context) {
	req := &bo.FindAllNoticesRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetNoticeService().FindAllNotices(c, req)
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

func (ins *NoticeController) UploadNotice(c *gin.Context) {
	req := &bo.UploadNoticeRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetNoticeService().UploadNotice(c, req)
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

func (ins *NoticeController) DeleteNotice(c *gin.Context) {
	req := &bo.DeleteNoticeRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetNoticeService().DeleteNotice(c, req)
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
