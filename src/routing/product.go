package routing

import (
	v1 "GraduationDesign/src/api/v1"
	mid "GraduationDesign/src/middleware"
	"github.com/gin-gonic/gin"
)

type product struct {
}

func (product) Init(router *gin.RouterGroup) {
	pg := router.Group("product", mid.MustUser())
	{
		pg.POST("upload", v1.Group.Product.UploadProduct)
		pgGroup := pg.Group("ps")
		{
			pgGroup.GET("user", v1.Group.Product.GetUserLendProduct)
			pgGroup.GET("info", v1.Group.Product.GetUserInfoProduct)
		}
	}
}
