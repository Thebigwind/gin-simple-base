package routers

import (
	"gin-simple-base/middleware/jwt"
	v1 "gin-simple-base/routers/api/v1"
	"gin-simple-base/routers/rest"
	"github.com/gin-gonic/gin"
	"time"

	"gin-simple-base/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":   "test",
			"timestamp": time.Now().Unix(),
		})
	})
	r.GET("/auth", rest.GetAuth)
	//r.Use(jwt.JWT())

	r.POST("/user/add", rest.AddUser)
	r.GET("/user/info", jwt.JWT(), rest.GetUser)
	r.GET("/user/list", rest.GetUserList)

	apiv1 := r.Group("/api/v1")
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}

	return r
}
