package handler

import "github.com/gin-gonic/gin"

type DepartmentController struct{}

func RegisterDepartmentRouter(r *gin.RouterGroup) {
	departmentController := &DepartmentController{}
	{
		r.GET("/find_all_departments", departmentController.FindAllDepartments)
		r.POST("/add_department_permission", departmentController.AddDepartmentPermission)
		r.POST("/delete_department_permission", departmentController.DeleteDepartmentPermission)
	}
}

func (ins *DepartmentController) FindAllDepartments(c *gin.Context) {

}

func (ins *DepartmentController) AddDepartmentPermission(c *gin.Context) {

}

func (ins *DepartmentController) DeleteDepartmentPermission(c *gin.Context) {

}