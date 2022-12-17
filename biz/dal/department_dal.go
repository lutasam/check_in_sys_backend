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

type DepartmentDal struct{}

var (
	departmentDal     *DepartmentDal
	departmentDalOnce sync.Once
)

func GetDepartmentDal() *DepartmentDal {
	departmentDalOnce.Do(func() {
		departmentDal = &DepartmentDal{}
	})
	return departmentDal
}

func (ins *DepartmentDal) TakeDepartmentByName(c *gin.Context, name string) (*model.Department, error) {
	department := &model.Department{}
	departmentJSON, err := repository.GetRedis().Get(c, name+common.DEPARTMENTNAMESUFFIX).Result()
	if err != nil && errors.Is(err, redis.Nil) {
		err := repository.GetDB().WithContext(c).Table(department.TableName()).Preload("Admin").Where("name = ?", name).Find(department).Error
		if err != nil {
			return nil, common.DATABASEERROR
		}
		if department.Name == "" {
			return nil, common.DATANOTFOUND
		}
		go func() {
			j, err := json.Marshal(department)
			if err != nil {
				panic(err)
			}
			err = repository.GetRedis().Set(c, name+common.DEPARTMENTNAMESUFFIX, j, common.REDISEXPIRETIME).Err()
			if err != nil {
				panic(err)
			}
		}()
		return department, nil
	}
	err = json.Unmarshal([]byte(departmentJSON), department)
	if err != nil {
		return nil, common.UNKNOWNERROR
	}
	return department, nil
	//err := repository.GetDB().WithContext(c).Table(department.TableName()).Preload("Admin").Where("name = ?", name).Find(department).Error
	//if err != nil {
	//	return nil, common.DATABASEERROR
	//}
	//if department.Name == "" {
	//	return nil, common.DATANOTFOUND
	//}
	//return department, nil
}

func (ins *DepartmentDal) TakeDepartmentByID(c *gin.Context, id uint64) (*model.Department, error) {
	department := &model.Department{}
	departmentJSON, err := repository.GetRedis().Get(c, utils.Uint64ToString(id)+common.DEPARTMENTIDSUFFIX).Result()
	if err != nil && errors.Is(err, redis.Nil) {
		err := repository.GetDB().WithContext(c).Table(department.TableName()).Preload("Admin").Where("id = ?", id).Find(department).Error
		if err != nil {
			return nil, common.DATABASEERROR
		}
		if department.Name == "" {
			return nil, common.DATANOTFOUND
		}
		go func() {
			j, err := json.Marshal(department)
			if err != nil {
				panic(err)
			}
			err = repository.GetRedis().Set(c, utils.Uint64ToString(id)+common.DEPARTMENTIDSUFFIX, j, common.REDISEXPIRETIME).Err()
			if err != nil {
				panic(err)
			}
		}()
		return department, nil
	}
	err = json.Unmarshal([]byte(departmentJSON), department)
	if err != nil {
		return nil, common.UNKNOWNERROR
	}
	return department, nil
}

func (ins *DepartmentDal) FindDepartments(c *gin.Context, currentPage, pageSize int) ([]*model.Department, int64, error) {
	var departments []*model.Department
	var count int64
	err := repository.GetDB().WithContext(c).Table(model.Department{}.TableName()).Where("id != ? and deleted_at is null", common.ALLDEPARTMENTS).Count(&count).Preload("Admin").
		Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&departments).Error
	if err != nil {
		return nil, 0, common.DATABASEERROR
	}
	return departments, count, nil
}

func (ins *DepartmentDal) UpdateDepartment(c *gin.Context, department *model.Department) error {
	err := repository.GetDB().WithContext(c).Model(department).Updates(department).Error
	if err != nil {
		return common.DATABASEERROR
	}
	go func() {
		err := repository.GetRedis().Del(c, utils.Uint64ToString(department.ID)+common.DEPARTMENTIDSUFFIX).Err()
		if err != nil && !errors.Is(err, redis.Nil) {
			panic(err)
		}
		err = repository.GetRedis().Del(c, department.Name+common.DEPARTMENTNAMESUFFIX).Err()
		if err != nil && !errors.Is(err, redis.Nil) {
			panic(err)
		}
	}()
	return nil
}

func (ins *DepartmentDal) FindAllDepartmentsByAdminID(c *gin.Context, adminID uint64) ([]*model.Department, error) {
	var departments []*model.Department
	err := repository.GetDB().WithContext(c).Table(model.Department{}.TableName()).Where("admin_id = ?", adminID).Preload("Admin").Find(&departments).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	return departments, nil
}
