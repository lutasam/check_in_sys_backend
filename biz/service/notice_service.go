package service

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/check_in_sys/biz/bo"
	"github.com/lutasam/check_in_sys/biz/common"
	"github.com/lutasam/check_in_sys/biz/dal"
	"github.com/lutasam/check_in_sys/biz/model"
	"github.com/lutasam/check_in_sys/biz/utils"
	"sync"
)

type NoticeService struct{}

var (
	noticeService     *NoticeService
	noticeServiceOnce sync.Once
)

func GetNoticeService() *NoticeService {
	noticeServiceOnce.Do(func() {
		noticeService = &NoticeService{}
	})
	return noticeService
}

func (ins *NoticeService) FindAllNotices(c *gin.Context, req *bo.FindAllNoticesRequest) (*bo.FindAllNoticesResponse, error) {
	if req.CurrentPage < 0 || req.PageSize < 0 || req.PageSize > 100 {
		return nil, common.USERINPUTERROR
	}
	notices, total, err := dal.GetNoticeDal().FindAllNotices(c, req.CurrentPage, req.PageSize)
	if err != nil {
		return nil, err
	}
	return &bo.FindAllNoticesResponse{
		Total:   int(total),
		Notices: convertToNoticeVO(notices),
	}, nil
}

func (ins *NoticeService) UploadNotice(c *gin.Context, req *bo.UploadNoticeRequest) (*bo.UploadNoticeResponse, error) {
	if req.Content == "" {
		return nil, common.USERINPUTERROR
	}
	userInfo, err := utils.GetCtxUserInfoJWT(c)
	if err != nil {
		return nil, err
	}
	err = dal.GetNoticeDal().CreateNotice(c, &model.Notice{
		ID:      utils.GenerateNoticeID(),
		UserID:  userInfo.UserID,
		Content: req.Content,
	})
	if err != nil {
		return nil, err
	}
	return &bo.UploadNoticeResponse{}, nil
}

func (ins *NoticeService) DeleteNotice(c *gin.Context, req *bo.DeleteNoticeRequest) (*bo.DeleteNoticeResponse, error) {
	id, err := utils.StringToUint64(req.NoticeID)
	if err != nil {
		return nil, err
	}
	userInfo, err := utils.GetCtxUserInfoJWT(c)
	if err != nil {
		return nil, err
	}
	notice, err := dal.GetNoticeDal().FindNoticeByID(c, id)
	if err != nil {
		return nil, err
	}
	if notice.UserID != userInfo.UserID {
		return nil, common.HAVENOPERMISSION
	}
	err = dal.GetNoticeDal().DeleteNotice(c, notice)
	if err != nil {
		return nil, err
	}
	return &bo.DeleteNoticeResponse{}, nil
}
