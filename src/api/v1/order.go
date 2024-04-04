package v1

import (
	"GraduationDesign/src/logic"
	mid "GraduationDesign/src/middleware"
	"GraduationDesign/src/model"
	"GraduationDesign/src/model/reply"
	"GraduationDesign/src/model/request"
	"GraduationDesign/src/myerr"
	"github.com/0RAJA/Rutils/pkg/app"
	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/xid"
	"strconv"
)

type orders struct {
}

func (orders) CreatOrder(c *gin.Context) {
	rly := app.NewResponse(c)
	p := &request.Order{}
	if err := c.ShouldBindJSON(p); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.UserToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	orderID := strconv.FormatInt(xid.Next(), 10)
	err := logic.Group.Order.CreateOrder(c, *p, orderID)
	rly.Reply(err, reply.OrderID{ID: orderID})
}

func (orders) GetOrderList(c *gin.Context) {
	rly := app.NewResponse(c)
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.UserToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	rsp, err := logic.Group.Order.GetOrderList(c, content.ID)
	rly.Reply(err, rsp)
}

func (orders) ChangeOrderStatus(c *gin.Context) {
	rly := app.NewResponse(c)
	req := &request.ChangeOrderStatus{}
	if err := c.ShouldBindJSON(req); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.UserToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	err := logic.Group.Order.ChangeOrderStatus(c, req)
	rly.Reply(err, nil)
}
