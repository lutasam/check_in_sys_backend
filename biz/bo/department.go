package bo

import "github.com/lutasam/check_in_sys/biz/vo"

type FindAllDepartmentsRequest struct {
	CurrentPage int `json:"current_page" binding:"required"`
	PageSize    int `json:"page_size" binding:"required"`
}

type FindAllDepartmentsResponse struct {
	Total       int                `json:"total"`
	Departments []*vo.DepartmentVO `json:"departments"`
}

type AddDepartmentPermissionRequest struct {
	DepartmentID string `json:"department_id" binding:"required"`
	AdminID      string `json:"admin_id" binding:"required"`
}

type AddDepartmentPermissionResponse struct{}

type DeleteDepartmentPermissionRequest struct {
	DepartmentID string `json:"department_id"`
}

type DeleteDepartmentPermissionResponse struct{}
