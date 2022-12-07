package bo

import "github.com/lutasam/check_in_sys/biz/vo"

type FindAllNoticesRequest struct {
	CurrentPage int `json:"current_page" binding:"required"`
	PageSize    int `json:"page_size" binding:"required"`
}

type FindAllNoticesResponse struct {
	Total   int            `json:"total"`
	Notices []*vo.NoticeVO `json:"notices"`
}

type UploadNoticeRequest struct {
	Content string `json:"content" binding:"required"`
}

type UploadNoticeResponse struct{}

type DeleteNoticeRequest struct {
	NoticeID string `json:"notice_id"`
}

type DeleteNoticeResponse struct{}
