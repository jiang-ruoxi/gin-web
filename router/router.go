package router

import (
	"web/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	//后台路由组
	admin := router.Group("v1")
	admin.Use()
	{
		admin.GET("/index", controller.AdminIndex)
	}

	router.LoadHTMLGlob("view/**/*")
	router.Static("/static", "./static")
	return router
}
