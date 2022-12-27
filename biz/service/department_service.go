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

func (ins *DepartmentService) FindAllDepartments(c *gin.Context, req *bo.FindAllDepartmentsRequest) (*bo.FindAllDepartmentsResponse, error) {
	if req.CurrentPage < 0 || req.PageSize < 0 || req.PageSize > 100 {
		return nil, common.USERINPUTERROR
	}
	departments, total, err := dal.GetDepartmentDal().FindDepartments(c, req.CurrentPage, req.PageSize)
	if err != nil {
		return nil, err
	}
	return &bo.FindAllDepartmentsResponse{
		Total:       int(total),
		Departments: convertToDepartmentVOs(departments),
	}, nil
}

func (ins *DepartmentService) AddDepartmentPermission(c *gin.Context, req *bo.AddDepartmentPermissionRequest) (*bo.AddDepartmentPermissionResponse, error) {
	userID, err := utils.StringToUint64(req.AdminID)
	departmentID, err := utils.StringToUint64(req.DepartmentID)
	if err != nil {
		return nil, common.USERINPUTERROR
	}
	user, err := dal.GetUserDal().TakeUserByID(c, userID)
	if err != nil {
		return nil, err
	}
	if user.DepartmentID != departmentID {
		return nil, common.NOTINDEPARTMENT
	}
	department, err := dal.GetDepartmentDal().TakeDepartmentByID(c, departmentID)
	if err != nil {
		return nil, err
	}
	user.DepartmentID = departmentID
	err = dal.GetUserDal().UpdateUser(c, user)
	if err != nil {
		return nil, err
	}
	department.Admin = *user
	err = dal.GetDepartmentDal().UpdateDepartment(c, department)
	if err != nil {
		return nil, err
	}
	return &bo.AddDepartmentPermissionResponse{}, nil
}

func (ins *DepartmentService) DeleteDepartmentPermission(c *gin.Context, req *bo.DeleteDepartmentPermissionRequest) (*bo.DeleteDepartmentPermissionResponse, error) {
	departmentID, err := utils.StringToUint64(req.DepartmentID)
	if err != nil {
		return nil, common.USERINPUTERROR
	}
	department, err := dal.GetDepartmentDal().TakeDepartmentByID(c, departmentID)
	if err != nil {
		return nil, err
	}
	superAdmin, err := dal.GetUserDal().TakeSuperAdmin(c)
	if err != nil {
		return nil, err
	}
	department.Admin = *superAdmin
	err = dal.GetDepartmentDal().UpdateDepartment(c, department)
	if err != nil {
		return nil, err
	}
	return &bo.DeleteDepartmentPermissionResponse{}, nil
}

func convertToDepartmentVOs(departments []*model.Department) []*vo.DepartmentVO {
	var vos []*vo.DepartmentVO
	for _, department := range departments {
		vos = append(vos, &vo.DepartmentVO{
			ID:        utils.Uint64ToString(department.ID),
			Name:      department.Name,
			AdminName: department.Admin.Name,
			CreatedAt: utils.TimeToString(department.CreatedAt),
		})
	}
	return vos
}
