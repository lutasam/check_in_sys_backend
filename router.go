package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/check_in_sys/biz/handler"
	"github.com/lutasam/check_in_sys/biz/middleware"
	"io"
	"os"
)

func InitRouterAndMiddleware(r *gin.Engine) {
	// 设置log文件输出
	logFile, err := os.Create("log/log.log")
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)

	// 注册全局中间件Recovery和Logger
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 限制传输文件 最大32MB
	r.MaxMultipartMemory = 32 << 20

	// 注册分组路由
	// 登录模块
	login := r.Group("/login")
	handler.RegisterLoginRouter(login)

	// 用户模块
	user := r.Group("/user", middleware.JWTAuth())
	handler.RegisterUserRouter(user)

	// 文件模块
	file := r.Group("/file", middleware.JWTAuth())
	handler.RegisterFileRouter(file)
}
