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

func (ins *RecordService) FindAllUserRecord(c *gin.Context, req *bo.FindUserAllRecordsRequest) (*bo.FindUserAllRecordsResponse, error) {
	if req.CurrentPage < 0 || req.PageSize < 0 || req.PageSize > 100 {
		return nil, common.USERINPUTERROR
	}
	records, err := dal.GetRecordDal().FindRecords(c, req.CurrentPage, req.PageSize)
	if err != nil {
		return nil, err
	}
	return &bo.FindUserAllRecordsResponse{
		Total:   len(records),
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
		TemperatureRange: req.TemperatureRange,
		IsHealthy:        req.IsHealthy,
		HealthCodeStatus: req.HealthCodeStatus,
		Remark:           req.Remark,
		Appendix:         req.Appendix,
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
