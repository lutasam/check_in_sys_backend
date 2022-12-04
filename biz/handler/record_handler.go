package handler

import (
	"github.com/gin-gonic/gin"
)

type RecordController struct{}

func RegisterRecordRouter(r *gin.RouterGroup) {
	recordController := &RecordController{}
	{
		r.GET("/find_user_all_records", recordController.FindUserAllRecords)
		r.POST("/upload_user_record", recordController.UploadUserRecord)
		r.GET("/find_user_today_record", recordController.FindUserTodayRecord)
	}
}

func (ins *RecordController) UploadUserRecord(c *gin.Context) {

}

func (ins *RecordController) FindUserAllRecords(c *gin.Context) {

}

func (ins *RecordController) FindUserTodayRecord(c *gin.Context) {

}
