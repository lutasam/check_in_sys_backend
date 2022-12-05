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

func (ins *RecordDal) FindRecords(c *gin.Context, currentPage, pageSize int, userID uint64) ([]*model.Record, error) {
	var records []*model.Record
	err := repository.GetDB().WithContext(c).Table(model.Record{}.TableName()).Where("user_id = ?", userID).
		Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&records).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	return records, nil
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
