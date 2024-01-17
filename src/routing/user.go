package routing

type user struct {
}

//func (user) Init(router *gin.RouterGroup) {
//	ug := router.Group("user")
//	{
//		ug.POST("register", v1.Group.User.Register)
//		ug.POST("login", v1.Group.User.Login)
//		updateGroup := ug.Group("update").Use(mid.MustUser())
//		{
//			updateGroup.PUT("email", v1.Group.User.UpdateUserEmail)
//			updateGroup.PUT("pwd", v1.Group.User.UpdateUserPassword)
//		}
//		ug.DELETE("delete", mid.MustUser(), v1.Group.User.DeleteUser)
//	}
//}
