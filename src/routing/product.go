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
			pgGroup.GET("need", v1.Group.Product.GetUserNeedProduct)
			pgGroup.GET("info", v1.Group.Product.GetProductInfo)
			pgGroup.GET("detail", v1.Group.Product.GetProductDetails)
		}
		pg.POST("delete", v1.Group.Product.DeleteProduct)
		pg.POST("update", v1.Group.Product.UpdateProduct)
		pg.POST("like_status", v1.Group.Product.ChangeLikeProduct)
		pg.GET("like_list", v1.Group.Product.GetLikeList)
		pg.GET("busy_time", v1.Group.Product.GetProductBusyTime)
	}
}
