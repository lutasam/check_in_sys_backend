package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/check_in_sys/biz/bo"
	"github.com/lutasam/check_in_sys/biz/common"
	"github.com/lutasam/check_in_sys/biz/service"
	"github.com/lutasam/check_in_sys/biz/utils"
)

type UserController struct{}

func RegisterUserRouter(r *gin.RouterGroup) {
	userController := &UserController{}
	{
		r.POST("/update_user_info", userController.UpdateUserInfo)
		r.POST("/find_all_user_status", userController.FindAllUserStatus)
		r.POST("/find_all_users", userController.FindAllUsers)
		r.POST("/delete_user", userController.DeleteUser)
		r.POST("/add_user", userController.AddUser)
	}
}

func (ins *UserController) UpdateUserInfo(c *gin.Context) {
	req := &bo.UpdateUserInfoRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetUserService().UpdateUserInfo(c, req)
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

func (ins *UserController) FindAllUserStatus(c *gin.Context) {
	req := &bo.FindAllUserStatusRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetUserService().FindAllUserStatus(c, req)
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

func (ins *UserController) FindAllUsers(c *gin.Context) {
	req := &bo.FindAllUsersRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetUserService().FindAllUsers(c, req)
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

func (ins *UserController) DeleteUser(c *gin.Context) {
	req := &bo.DeleteUserRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetUserService().DeleteUser(c, req)
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

func (ins *UserController) AddUser(c *gin.Context) {
	req := &bo.AddUserRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetUserService().AddUser(c, req)
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
