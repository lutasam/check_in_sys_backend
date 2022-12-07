package dal

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/check_in_sys/biz/common"
	"github.com/lutasam/check_in_sys/biz/model"
	"github.com/lutasam/check_in_sys/biz/repository"
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
	err := repository.GetDB().WithContext(c).Table(department.TableName()).Preload("Admin").Where("name = ?", name).Find(department).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	if department.ID == 0 {
		return nil, common.DATANOTFOUND
	}
	return department, nil
}

func (ins *DepartmentDal) TakeDepartmentByID(c *gin.Context, id uint64) (*model.Department, error) {
	department := &model.Department{}
	err := repository.GetDB().WithContext(c).Table(department.TableName()).Preload("Admin").Where("id = ?", id).Find(department).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	if department.Name == "" {
		return nil, common.DATANOTFOUND
	}
	return department, nil
}

func (ins *DepartmentDal) FindDepartments(c *gin.Context, currentPage, pageSize int) ([]*model.Department, int64, error) {
	var departments []*model.Department
	var count int64
	err := repository.GetDB().WithContext(c).Table(model.Department{}.TableName()).Where("id != ?", common.ALLDEPARTMENTS).Count(&count).Preload("Admin").
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
