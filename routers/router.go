package routers

import (
	_ "gin-blog/docs"
	"gin-blog/pkg/setting"
	"gin-blog/routers/api"
	v1 "gin-blog/routers/api/v1"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	r.GET("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiv1 := r.Group("/api/v1")
	{
		// 获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		// 新建标签
		apiv1.POST("/tags", v1.AddTag)
		// 编辑标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		// 获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		// 获取指定文章
		apiv1.GET("/article/:id", v1.GetArticle)
		// 新建文章
		apiv1.POST("/articles", v1.AddArticle)
		// 编辑文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		// 删除文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}
	//r.GET("/test", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"Message": "Test",
	//	})
	//})
	return r
}

/*
    获取权限： GET /auth
	获取标签列表：GET("/tags")
	新建标签：POST("/tags")
	更新指定标签：PUT("/tags/:id")
	删除指定标签：DELETE("/tags/:id")
	获取文章列表：GET("/articles")
	获取指定文章：GET("/articles/:id")
	新建文章：POST("/articles")
	更新指定文章：PUT("/articles/:id")
	删除指定文章：DELETE("/articles/:id")
*/
