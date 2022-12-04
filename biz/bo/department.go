package bo

import "github.com/lutasam/check_in_sys/biz/vo"

type FindAllDepartmentsRequest struct{}

type FindAllDepartmentsResponse struct {
	Total       int                `json:"total"`
	departments []*vo.DepartmentVO `json:"departments"`
}

type AddDepartmentPermissionRequest struct {
	DepartmentID string `json:"department_id"`
	AdminID      string `json:"admin_id"`
}

type AddDepartmentPermissionResponse struct{}

type DeleteDepartmentPermissionRequest struct {
	DepartmentID string `json:"department_id"`
}

type DeleteDepartmentPermissionResponse struct{}
