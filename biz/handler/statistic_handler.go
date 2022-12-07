package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/check_in_sys/biz/bo"
	"github.com/lutasam/check_in_sys/biz/common"
	"github.com/lutasam/check_in_sys/biz/service"
	"github.com/lutasam/check_in_sys/biz/utils"
)

type StatisticController struct{}

func RegisterStatisticRouter(r *gin.RouterGroup) {
	statisticController := &StatisticController{}
	{
		r.POST("/take_user_statistic", statisticController.TakeUserStatistic)
		r.POST("/take_admin_statistic", statisticController.TakeAdminStatistic)
	}
}

func (ins *StatisticController) TakeUserStatistic(c *gin.Context) {
	req := &bo.TakeUserStatisticRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetStatisticService().TakeUserStatistic(c, req)
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

func (ins *StatisticController) TakeAdminStatistic(c *gin.Context) {
	req := &bo.TakeAdminStatisticRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetStatisticService().TakeAdminStatistic(c, req)
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
