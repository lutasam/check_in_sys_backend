package service

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/check_in_sys/biz/bo"
	"github.com/lutasam/check_in_sys/biz/common"
	"github.com/lutasam/check_in_sys/biz/dal"
	"github.com/lutasam/check_in_sys/biz/model"
	"github.com/lutasam/check_in_sys/biz/utils"
	"github.com/lutasam/check_in_sys/biz/vo"
	"sync"
)

type RecordService struct{}

var (
	recordService     *RecordService
	recordServiceOnce sync.Once
)

func GetRecordService() *RecordService {
	recordServiceOnce.Do(func() {
		recordService = &RecordService{}
	})
	return recordService
}

func (ins *RecordService) FindUserAllRecords(c *gin.Context, req *bo.FindUserAllRecordsRequest) (*bo.FindUserAllRecordsResponse, error) {
	if req.CurrentPage < 0 || req.PageSize < 0 || req.PageSize > 100 {
		return nil, common.USERINPUTERROR
	}
	userInfo, err := utils.GetCtxUserInfoJWT(c)
	if err != nil {
		return nil, err
	}
	records, total, err := dal.GetRecordDal().FindRecords(c, req.CurrentPage, req.PageSize, userInfo.UserID)
	if err != nil {
		return nil, err
	}
	return &bo.FindUserAllRecordsResponse{
		Total:   int(total),
		Records: convertToRecordVOs(records),
	}, nil
}

func (ins *RecordService) UploadUserRecord(c *gin.Context, req *bo.UploadUserRecordRequest) (*bo.UploadUserRecordResponse, error) {
	userInfo, err := utils.GetCtxUserInfoJWT(c)
	if err != nil {
		return nil, err
	}
	err = dal.GetRecordDal().CreateRecord(c, &model.Record{
		ID:               utils.GenerateRecordID(),
		UserID:           userInfo.UserID,
		Address:          req.Address,
		TemperatureRange: *req.TemperatureRange,
		IsHealthy:        *req.IsHealthy,
		HealthCodeStatus: *req.HealthCodeStatus,
		Remark:           req.Remark,
		Appendix:         req.Appendix,
	})
	if err != nil {
		return nil, err
	}
	err = dal.GetUserDal().UpdateUser(c, &model.User{
		ID:                    userInfo.UserID,
		TodayRecordStatus:     true,
		TodayHealthCodeStatus: *req.HealthCodeStatus,
	})
	if err != nil {
		return nil, err
	}
	return &bo.UploadUserRecordResponse{}, nil
}

func (ins *RecordService) FindUserTodayRecord(c *gin.Context, req *bo.FindUserTodayRecordRequest) (*bo.FindUserTodayRecordResponse, error) {
	id, err := utils.StringToUint64(req.UserID)
	if err != nil {
		return nil, err
	}
	_, err = dal.GetUserDal().TakeUserByID(c, id)
	if err != nil {
		return nil, err
	}
	record, err := dal.GetRecordDal().TakeRecentUserRecord(c, id)
	if err != nil {
		return nil, err
	}
	return &bo.FindUserTodayRecordResponse{
		UserRecord: convertToUserRecordVO(record),
	}, nil
}

func (ins *RecordService) NoticeUserFinishRecord(c *gin.Context, req *bo.NoticeUserFinishRecordRequest) (*bo.NoticeUserFinishRecordResponse, error) {
	id, err := utils.StringToUint64(req.UserID)
	if err != nil {
		return nil, err
	}
	user, err := dal.GetUserDal().TakeUserByID(c, id)
	if err != nil {
		return nil, err
	}
	if user.TodayRecordStatus == true {
		return nil, common.NONEEDNOTICE
	}
	go func() {
		// err = sendNoticeEmail(c, user.Email)
		err = sendNoticeEmail(c, "229644572@qq.com") // mock
		if err != nil {
			panic(err)
		}
	}()
	return &bo.NoticeUserFinishRecordResponse{}, nil
}

func (ins *RecordService) NoticeAllUserNotFinishRecord(c *gin.Context, req *bo.NoticeAllUserNotFinishRecordRequest) (*bo.NoticeAllUserNotFinishRecordResponse, error) {
	userInfo, err := utils.GetCtxUserInfoJWT(c)
	if err != nil {
		return nil, err
	}
	user, err := dal.GetUserDal().TakeUserByID(c, userInfo.UserID)
	if err != nil {
		return nil, err
	}
	var departmentIDs []uint64
	if user.Identity == common.SUPER_ADMIN.Ints() {
		departmentIDs = append(departmentIDs, common.ALLDEPARTMENTS)
	} else {
		departments, err := dal.GetDepartmentDal().FindAllDepartmentsByAdminID(c, user.ID)
		if err != nil {
			return nil, err
		}
		for _, department := range departments {
			departmentIDs = append(departmentIDs, department.ID)
		}
	}
	users, err := dal.GetUserDal().FindAllUsersInDepartmentGroup(c, departmentIDs)
	if err != nil {
		return nil, err
	}

	go func() {
		// 模拟的邮箱不是真的邮箱，发送可能会报异常所以采用模拟的方式发送
		for _, user := range users {
			// err := sendNoticeEmail(c, user.Email)
			if user.TodayRecordStatus == true || user.Identity != common.USER.Ints() {
				continue
			}
			err = sendNoticeEmail(c, "229644572@qq.com") // mock
			if err != nil {
				panic(err.Error() + " " + user.Email)
			}
			break
		}
	}()

	return &bo.NoticeAllUserNotFinishRecordResponse{}, nil
}

func convertToRecordVOs(records []*model.Record) []*vo.RecordVO {
	var vos []*vo.RecordVO
	for _, record := range records {
		vos = append(vos, &vo.RecordVO{
			ID:               utils.Uint64ToString(record.ID),
			Address:          record.Address,
			TemperatureRange: common.ParseTemperatureRange(record.TemperatureRange).String(),
			IsHealthy:        record.IsHealthy,
			HealthCodeStatus: common.ParseHealthCodeStatus(record.HealthCodeStatus).String(),
			CreatedAt:        utils.TimeToString(record.CreatedAt),
		})
	}
	return vos
}

func convertToUserRecordVO(record *model.Record) *vo.UserRecordVO {
	return &vo.UserRecordVO{
		ID:               utils.Uint64ToString(record.ID),
		UserID:           utils.Uint64ToString(record.UserID),
		Address:          record.Address,
		TemperatureRange: common.ParseTemperatureRange(record.TemperatureRange).String(),
		IsHealthy:        record.IsHealthy,
		HealthCodeStatus: common.ParseHealthCodeStatus(record.HealthCodeStatus).String(),
		Remark:           record.Remark,
		Appendix:         record.Appendix,
	}
}

func sendNoticeEmail(c *gin.Context, email string) error {
	subject := "[打卡提醒]健康打卡管理系统"
	body := `
今日您的健康打卡还尚未完成，请登陆系统进行打卡！<br>
如果您没有注册本系统，请忽视该邮件。
`
	err := utils.SendMail(email, subject, body)
	if err != nil {
		return err
	}
	return nil
}
