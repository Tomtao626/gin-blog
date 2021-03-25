package routes

import (
	"gin-blog/api/v1"
	"gin-blog/utils"
	"github.com/gin-gonic/gin"
)

func InitRoute() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	route := r.Group("api/v1")
	{
		//用户模块的路由接口
		route.POST("user/add", v1.AddUser)
		route.GET("users", v1.GetUsers)
		route.PUT("user/:id", v1.EditUser)
		route.DELETE("user/:id", v1.DelUser)
		//分类模块的路由接口

		//文章模块的路由接口
	}
	r.Run(utils.HttpPort)
}
