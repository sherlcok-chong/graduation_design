package routing

import (
	v1 "GraduationDesign/src/api/v1"
	mid "GraduationDesign/src/middleware"
	"github.com/gin-gonic/gin"
)

type user struct {
}

func (user) Init(router *gin.RouterGroup) {
	ug := router.Group("user")
	{
		ug.POST("register", v1.Group.User.Register)
		ug.POST("login", v1.Group.User.Login)
		updateGroup := ug.Group("update").Use(mid.MustUser())
		{
			updateGroup.POST("info", v1.Group.User.UpdateUserInfo)
			//	updateGroup.PUT("pwd", v1.Group.User.UpdateUserPassword)
		}
		ug.GET("info", v1.Group.User.GetUserInfo).Use(mid.MustUser())
		//ug.DELETE("delete", mid.MustUser(), v1.Group.User.DeleteUser)
	}
}
