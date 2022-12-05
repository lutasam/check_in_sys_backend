package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/check_in_sys/biz/bo"
	"github.com/lutasam/check_in_sys/biz/common"
	"github.com/lutasam/check_in_sys/biz/service"
	"github.com/lutasam/check_in_sys/biz/utils"
)

type DepartmentController struct{}

func RegisterDepartmentRouter(r *gin.RouterGroup) {
	departmentController := &DepartmentController{}
	{
		r.POST("/find_all_departments", departmentController.FindAllDepartments)
		r.POST("/add_department_permission", departmentController.AddDepartmentPermission)
		r.POST("/delete_department_permission", departmentController.DeleteDepartmentPermission)
	}
}

func (ins *DepartmentController) FindAllDepartments(c *gin.Context) {
	req := &bo.FindAllDepartmentsRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetDepartmentService().FindAllDepartments(c, req)
	if err != nil {
		if utils.IsClientError(err) {
			utils.ResponseClientError(c, err.(common.Error))
			return
		} else {
			utils.ResponseServerError(c, err.(common.Error))
			return
		}
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *DepartmentController) AddDepartmentPermission(c *gin.Context) {
	req := &bo.AddDepartmentPermissionRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetDepartmentService().AddDepartmentPermission(c, req)
	if err != nil {
		if utils.IsClientError(err) {
			utils.ResponseClientError(c, err.(common.Error))
			return
		} else {
			utils.ResponseServerError(c, err.(common.Error))
			return
		}
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *DepartmentController) DeleteDepartmentPermission(c *gin.Context) {
	req := &bo.DeleteDepartmentPermissionRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetDepartmentService().DeleteDepartmentPermission(c, req)
	if err != nil {
		if utils.IsClientError(err) {
			utils.ResponseClientError(c, err.(common.Error))
			return
		} else {
			utils.ResponseServerError(c, err.(common.Error))
			return
		}
	}
	utils.ResponseSuccess(c, resp)
}
