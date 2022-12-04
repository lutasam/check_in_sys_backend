package service

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/check_in_sys/biz/bo"
	"sync"
)

type DepartmentService struct{}

var (
	departmentService     *DepartmentService
	departmentServiceOnce sync.Once
)

func GetDepartmentService() *DepartmentService {
	departmentServiceOnce.Do(func() {
		departmentService = &DepartmentService{}
	})
	return departmentService
}

func (ins *DepartmentService) FindAllDepartment(c *gin.Context, req *bo.FindAllDepartmentsRequest) (*bo.FindAllDepartmentsResponse, error) {

}

func (ins *DepartmentService) AddDepartmentPermission(c *gin.Context, req *bo.AddDepartmentPermissionRequest) (*bo.AddDepartmentPermissionResponse, error) {

}

func (ins *DepartmentService) DeleteDepartmentPermission(c *gin.Context, req *bo.DeleteDepartmentPermissionRequest) (*bo.DeleteDepartmentPermissionResponse, error) {

}