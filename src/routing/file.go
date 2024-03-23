package routing

import (
	v1 "GraduationDesign/src/api/v1"
	mid "GraduationDesign/src/middleware"
	"github.com/gin-gonic/gin"
)

type file struct {
}

func (file) Init(router *gin.RouterGroup) {
	fg := router.Group("file")
	{
		uploadGroup := fg.Group("upload").Use(mid.MustUser())
		{
			uploadGroup.POST("avatar", v1.Group.File.UpdateUserAvatar)
			uploadGroup.POST("file", v1.Group.File.UploadFile)
		}

	}
}
