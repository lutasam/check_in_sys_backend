package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/check_in_sys/biz/bo"
	"github.com/lutasam/check_in_sys/biz/common"
	"github.com/lutasam/check_in_sys/biz/service"
	"github.com/lutasam/check_in_sys/biz/utils"
)

type RecordController struct{}

func RegisterRecordRouter(r *gin.RouterGroup) {
	recordController := &RecordController{}
	{
		r.POST("/find_user_all_records", recordController.FindUserAllRecords)
		r.POST("/upload_user_record", recordController.UploadUserRecord)
		r.POST("/find_user_today_record", recordController.FindUserTodayRecord)
		r.POST("/notice_user_finish_record", recordController.NoticeUserFinishRecord)
		r.POST("/notice_all_user_not_finish_record", recordController.NoticeAllUserNotFinishRecord)
	}
}

func (ins *RecordController) UploadUserRecord(c *gin.Context) {
	req := &bo.UploadUserRecordRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetRecordService().UploadUserRecord(c, req)
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

func (ins *RecordController) FindUserAllRecords(c *gin.Context) {
	req := &bo.FindUserAllRecordsRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetRecordService().FindUserAllRecords(c, req)
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

func (ins *RecordController) FindUserTodayRecord(c *gin.Context) {
	req := &bo.FindUserTodayRecordRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetRecordService().FindUserTodayRecord(c, req)
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

func (ins *RecordController) NoticeUserFinishRecord(c *gin.Context) {
	req := &bo.NoticeUserFinishRecordRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetRecordService().NoticeUserFinishRecord(c, req)
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

func (ins *RecordController) NoticeAllUserNotFinishRecord(c *gin.Context) {
	req := &bo.NoticeAllUserNotFinishRecordRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetRecordService().NoticeAllUserNotFinishRecord(c, req)
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
