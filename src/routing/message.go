package routing

import (
	v1 "GraduationDesign/src/api/v1"
	mid "GraduationDesign/src/middleware"
	"github.com/gin-gonic/gin"
)

type msg struct {
}

func (msg) Init(router *gin.RouterGroup) {
	mg := router.Group("msg", mid.MustUser())
	{
		mg.GET("not_read", v1.Group.Ws.GetNotReadMsg)
	}
}
