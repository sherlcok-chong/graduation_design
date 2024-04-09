package routing

import (
	v1 "GraduationDesign/src/api/v1"
	mid "GraduationDesign/src/middleware"
	"github.com/gin-gonic/gin"
)

type comment struct {
}

func (comment) Init(router *gin.RouterGroup) {
	tg := router.Group("comment", mid.MustUser())
	{
		tg.POST("add", v1.Group.Comment.AddComment)
		tg.POST("delete", v1.Group.Comment.DeleteComment)
		tg.GET("get", v1.Group.Comment.GetProductComment)
	}
}
