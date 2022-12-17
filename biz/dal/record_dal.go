package dal

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/check_in_sys/biz/common"
	"github.com/lutasam/check_in_sys/biz/model"
	"github.com/lutasam/check_in_sys/biz/repository"
	"sync"
)

type RecordDal struct{}

var (
	recordDal     *RecordDal
	recordDalOnce sync.Once
)

func GetRecordDal() *RecordDal {
	recordDalOnce.Do(func() {
		recordDal = &RecordDal{}
	})
	return recordDal
}

func (ins *RecordDal) FindRecords(c *gin.Context, currentPage, pageSize int, userID uint64) ([]*model.Record, int64, error) {
	var records []*model.Record
	var count int64
	err := repository.GetDB().WithContext(c).Table(model.Record{}.TableName()).Where("user_id = ? and deleted_at is null", userID).Count(&count).
		Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&records).Error
	if err != nil {
		return nil, 0, common.DATABASEERROR
	}
	return records, count, nil
}

func (ins *RecordDal) CountUserRecords(c *gin.Context, userID uint64) (int64, error) {
	var count int64
	err := repository.GetDB().WithContext(c).Table(model.Record{}.TableName()).Where("user_id = ? and deleted_at is null", userID).Count(&count).Error
	if err != nil {
		return 0, common.DATABASEERROR
	}
	return count, nil
}

func (ins *RecordDal) CreateRecord(c *gin.Context, record *model.Record) error {
	err := repository.GetDB().WithContext(c).Table(record.TableName()).Create(record).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}

func (ins *RecordDal) TakeRecentUserRecord(c *gin.Context, userID uint64) (*model.Record, error) {
	record := &model.Record{}
	err := repository.GetDB().WithContext(c).Table(record.TableName()).
		Where("user_id = ?", userID).Order("created_at desc").Limit(1).Find(record).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	return record, nil
}

func (ins *RecordDal) DeleteUserAllRecords(c *gin.Context, user *model.User) error {
	err := repository.GetDB().WithContext(c).Table(model.Record{}.TableName()).
		Where("user_id = ?", user.ID).Delete(&model.Record{}).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}
