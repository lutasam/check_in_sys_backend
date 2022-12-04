package dal

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/check_in_sys/biz/model"
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

}

func (ins *DepartmentDal) TakeDepartmentByID(c *gin.Context, id uint64) (*model.Department, error) {

}
