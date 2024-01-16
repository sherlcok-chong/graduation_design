package router

import (
	_ "GraduationDesign/docs"
	"GraduationDesign/src/global"
	mid "GraduationDesign/src/middleware"
	"github.com/0RAJA/Rutils/pkg/app"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(mid.Cors(), mid.GinLogger(), mid.Recovery(true))
	root := r.Group("api", mid.LogBody(), mid.Auth())
	{
		root.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		root.GET("ping", func(c *gin.Context) {
			rly := app.NewResponse(c)
			global.Logger.Info("ping", mid.ErrLogMsg(c)...)
			rly.Reply(nil, "pang")
		})
	}
	return r
}
