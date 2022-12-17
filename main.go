package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/check_in_sys/biz/utils"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.New()

	// 路由和中间件注册
	InitRouterAndMiddleware(r)

	// 定时任务注册
	InitCronJob()

	err := r.Run(":" + utils.GetConfigString("server.port"))
	if err != nil {
		panic(err)
	}
}
