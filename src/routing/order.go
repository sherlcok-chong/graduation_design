package routing

import (
	v1 "GraduationDesign/src/api/v1"
	mid "GraduationDesign/src/middleware"
	"github.com/gin-gonic/gin"
)

type order struct {
}

func (order) Init(router *gin.RouterGroup) {
	og := router.Group("order", mid.MustUser())
	{
		og.POST("add", v1.Group.Orders.CreatOrder)
		og.POST("change_status", v1.Group.Orders.ChangeOrderStatus)
		og.GET("list", v1.Group.Orders.GetOrderList)

	}
}
