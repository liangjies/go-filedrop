package cmd

import (
	"go-filedrop/service"

	"github.com/gin-gonic/gin"
)

func Execute() {
	Router := gin.Default()
	Router.StaticFile("/", "./web/index.html") // 前端网页入口页面
	Router.GET("/get", service.Get)
	Router.GET("/ping", service.Ping)
	Router.Run(":8000")
}
