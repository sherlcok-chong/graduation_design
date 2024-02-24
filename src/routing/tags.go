package routing

import (
	v1 "GraduationDesign/src/api/v1"
	mid "GraduationDesign/src/middleware"
	"github.com/gin-gonic/gin"
)

type tags struct {
}

func (tags) Init(router *gin.RouterGroup) {
	tg := router.Group("tags", mid.MustUser())
	{
		tg.GET("get_all", v1.Group.Tags.GetAllTags)
	}
}
