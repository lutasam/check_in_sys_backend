package dal

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/check_in_sys/biz/common"
	"github.com/lutasam/check_in_sys/biz/model"
	"github.com/lutasam/check_in_sys/biz/repository"
	"sync"
)

type NoticeDal struct{}

var (
	noticeDal     *NoticeDal
	noticeDalOnce sync.Once
)

func GetNoticeDal() *NoticeDal {
	noticeDalOnce.Do(func() {
		noticeDal = &NoticeDal{}
	})
	return noticeDal
}

// FindAllThreeDaysNotice only return notices in three days
func (ins *NoticeDal) FindAllThreeDaysNotice(c *gin.Context) ([]*model.Notice, error) {
	var notices []*model.Notice
	err := repository.GetDB().WithContext(c).Table(model.Notice{}.TableName()).Preload("User").
		Where("to_days(now())-to_days(created_at) <= 3").Find(&notices).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	return notices, nil
}

func (ins *NoticeDal) FindAllNotices(c *gin.Context, currentPage, pageSize int) ([]*model.Notice, int64, error) {
	var notices []*model.Notice
	var count int64
	err := repository.GetDB().WithContext(c).Table(model.Notice{}.TableName()).Preload("User").Count(&count).
		Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&notices).Error
	if err != nil {
		return nil, 0, common.DATABASEERROR
	}
	return notices, count, err
}

func (ins *NoticeDal) CreateNotice(c *gin.Context, notice *model.Notice) error {
	err := repository.GetDB().WithContext(c).Table(notice.TableName()).Create(notice).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}

func (ins *NoticeDal) DeleteNotice(c *gin.Context, noticeID uint64) error {
	err := repository.GetDB().WithContext(c).Table(model.Notice{}.TableName()).Where("id = ?", noticeID).
		Delete(&model.Notice{}).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}

func (ins *NoticeDal) FindNoticeByID(c *gin.Context, noticeID uint64) (*model.Notice, error) {
	notice := &model.Notice{}
	err := repository.GetDB().WithContext(c).Table(notice.TableName()).Where("id = ?", noticeID).Find(notice).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	if notice.ID == 0 {
		return nil, common.DATANOTFOUND
	}
	return notice, nil
}
