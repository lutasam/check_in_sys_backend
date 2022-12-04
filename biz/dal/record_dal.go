package dal

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/check_in_sys/biz/model"
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

func (ins *RecordDal) FindRecords(c *gin.Context, currentPage, pageSize int) ([]*model.Record, error) {

}

func (ins *RecordDal) CreateRecord(c *gin.Context, record *model.Record) error {

}

func (ins *RecordDal) TakeRecentUserRecord(c *gin.Context, userID uint64) (*model.Record, error) {

}