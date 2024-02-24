package router

import (
	_ "GraduationDesign/docs"
	"GraduationDesign/src/global"
	mid "GraduationDesign/src/middleware"
	"GraduationDesign/src/routing"
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
		rg := routing.Group
		rg.Email.Init(root)
		rg.User.Init(root)
		rg.File.Init(root)
		rg.Product.Init(root)
		rg.Tags.Init(root)
	}
	return r
}
