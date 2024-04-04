package router

import (
	_ "GraduationDesign/docs"
	v1 "GraduationDesign/src/api/v1"
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
		root.GET("/alipay", v1.Group.Alipay.PayUrl).Use(mid.Cors())
		root.GET("/callback", v1.Group.Alipay.Callback)
		root.POST("/notify", v1.Group.Alipay.Notify)
		rg := routing.Group
		rg.Email.Init(root)
		rg.User.Init(root)
		rg.File.Init(root)
		rg.Product.Init(root)
		rg.Tags.Init(root)
		rg.Comment.Init(root)
		rg.Order.Init(root)
		root.GET("/ws", v1.Group.Ws.WebSocket).Use(mid.MustUser())
		rg.Msg.Init(root)
	}
	return r
}
