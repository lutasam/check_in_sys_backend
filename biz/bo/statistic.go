package bo

import "github.com/lutasam/check_in_sys/biz/vo"

type TakeUserStatisticRequest struct{}

type TakeUserStatisticResponse struct {
	IsTodayRecordFinished bool           `json:"is_today_record_finish"`
	RecordTimes           int            `json:"record_times"`
	Notices               []*vo.NoticeVO `json:"notices"`
}

type TakeAdminStatisticRequest struct{}

type TakeAdminStatisticResponse struct {
	PeopleNum         int                    `json:"people_num"`
	AbnormalNum       int                    `json:"abnormal_num"`
	FinishRecordNum   int                    `json:"finish_record_num"`
	UnFinishRecordNum int                    `json:"unfinish_record_num"`
	FinishPercentage  int                    `json:"finished_percentage"`
	HealthCodeParts   []*vo.HealthCodePartVO `json:"health_code_parts"`
}
