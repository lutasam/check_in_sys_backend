package dal

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/lutasam/check_in_sys/biz/common"
	"github.com/lutasam/check_in_sys/biz/model"
	"github.com/lutasam/check_in_sys/biz/repository"
	"github.com/lutasam/check_in_sys/biz/utils"
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
	noticeJson, err := repository.GetRedis().Get(c, common.USERNOTICEKEY).Result()
	if err != nil && errors.Is(err, redis.Nil) || noticeJson == "" {
		err := repository.GetDB().WithContext(c).Table(model.Notice{}.TableName()).Preload("User").
			Where("to_days(now())-to_days(created_at) <= 3").Find(&notices).Error
		if err != nil {
			return nil, common.DATABASEERROR
		}
		go func() {
			j, err := json.Marshal(notices)
			if err != nil {
				panic(err)
			}
			err = repository.GetRedis().Set(c, common.USERNOTICEKEY, j, common.REDISEXPIRETIME).Err()
			if err != nil {
				panic(err)
			}
		}()
		return notices, nil
	}
	if err != nil {
		return nil, common.REDISERROR
	}
	err = json.Unmarshal([]byte(noticeJson), &notices)
	if err != nil {
		return nil, common.UNKNOWNERROR
	}
	return notices, nil
	//err := repository.GetDB().WithContext(c).Table(model.Notice{}.TableName()).Preload("User").
	//	Where("to_days(now())-to_days(created_at) <= 3").Find(&notices).Error
	//if err != nil {
	//	return nil, common.DATABASEERROR
	//}
	//return notices, nil
}

func (ins *NoticeDal) FindAllNotices(c *gin.Context, currentPage, pageSize int) ([]*model.Notice, int64, error) {
	var notices []*model.Notice
	var count int64
	err := repository.GetDB().WithContext(c).Table(model.Notice{}.TableName()).Where("deleted_at is null").Count(&count).Preload("User").
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
	go func() {
		err := repository.GetRedis().Del(c, common.USERNOTICEKEY).Err()
		if err != nil {
			panic(err)
		}
	}()
	return nil
}

func (ins *NoticeDal) DeleteNotice(c *gin.Context, notice *model.Notice) error {
	err := repository.GetDB().WithContext(c).Table(model.Notice{}.TableName()).Where("id = ?", notice.ID).
		Delete(&model.Notice{}).Error
	if err != nil {
		return common.DATABASEERROR
	}
	go func() {
		err := repository.GetRedis().Del(c, common.USERNOTICEKEY).Err()
		if err != nil && !errors.Is(err, redis.Nil) {
			panic(err)
		}
		err = repository.GetRedis().Del(c, utils.Uint64ToString(notice.ID)+common.NOTICEIDSUFFIX).Err()
		if err != nil && !errors.Is(err, redis.Nil) {
			panic(err)
		}
	}()
	return nil
}

func (ins *NoticeDal) FindNoticeByID(c *gin.Context, noticeID uint64) (*model.Notice, error) {
	notice := &model.Notice{}
	noticeJSON, err := repository.GetRedis().Get(c, utils.Uint64ToString(noticeID)+common.NOTICEIDSUFFIX).Result()
	if err != nil && errors.Is(err, redis.Nil) {
		err := repository.GetDB().WithContext(c).Table(notice.TableName()).Where("id = ?", noticeID).Find(notice).Error
		if err != nil {
			return nil, common.DATABASEERROR
		}
		if notice.ID == 0 {
			return nil, common.DATANOTFOUND
		}
		go func() {
			j, err := json.Marshal(notice)
			if err != nil {
				panic(err)
			}
			err = repository.GetRedis().Set(c, utils.Uint64ToString(noticeID)+common.NOTICEIDSUFFIX, j, common.REDISEXPIRETIME).Err()
			if err != nil {
				panic(err)
			}
		}()
		return notice, nil
	}
	if err != nil {
		return nil, common.REDISERROR
	}
	err = json.Unmarshal([]byte(noticeJSON), notice)
	if err != nil {
		return nil, common.UNKNOWNERROR
	}
	return notice, nil
	//err := repository.GetDB().WithContext(c).Table(notice.TableName()).Where("id = ?", noticeID).Find(notice).Error
	//if err != nil {
	//	return nil, common.DATABASEERROR
	//}
	//if notice.ID == 0 {
	//	return nil, common.DATANOTFOUND
	//}
	//return notice, nil
}
