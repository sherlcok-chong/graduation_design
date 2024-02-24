package v1

import (
	"GraduationDesign/src/logic"
	mid "GraduationDesign/src/middleware"
	"GraduationDesign/src/model"
	"GraduationDesign/src/model/request"
	"GraduationDesign/src/myerr"
	"github.com/0RAJA/Rutils/pkg/app"
	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/gin-gonic/gin"
)

type product struct {
}

func (product) UploadProduct(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.Product{}
	if err := c.ShouldBind(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.UserToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	err := logic.Group.Product.UploadProduct(c, params, content.ID)
	rly.Reply(err, nil)
}

func (product) GetUserLendProduct(c *gin.Context) {
	rly := app.NewResponse(c)
	p := &request.ProductInfo{}
	if err := c.ShouldBindQuery(p); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.UserToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	rsp, err := logic.Group.Product.GetUserProduct(c, content.ID)
	rly.Reply(err, rsp)
}

func (product) GetUserNeedProduct(c *gin.Context) {
	rly := app.NewResponse(c)
	p := &request.ProductInfo{}
	if err := c.ShouldBindQuery(p); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.UserToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	rsp, err := logic.Group.Product.GetUserNeed(c, content.ID)
	rly.Reply(err, rsp)
}

func (product) GetProductDetails(c *gin.Context) {
	rly := app.NewResponse(c)
	p := &request.ProductDetails{}
	if err := c.ShouldBindQuery(p); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.UserToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	rsp, err := logic.Group.Product.GetProductDetails(c, p.ID)
	rly.Reply(err, rsp)
}
func (product) GetProductInfo(c *gin.Context) {
	rly := app.NewResponse(c)
	p := &request.ProductInfo{}
	if err := c.ShouldBindQuery(p); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.UserToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	rsp, err := logic.Group.Product.GetProductInfo(c, p.Offset)
	rly.Reply(err, rsp)
}

func (product) DeleteProduct(c *gin.Context) {
	rly := app.NewResponse(c)
	p := &request.ProductID{}
	if err := c.ShouldBindJSON(p); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.UserToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	err := logic.Group.Product.DeleteProduct(c, p.ID, content.ID)
	rly.Reply(err, nil)
}

func (product) UpdateProduct(c *gin.Context) {
	rly := app.NewResponse(c)
	p := &request.UpdateProduct{}
	if err := c.ShouldBind(p); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.UserToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	err := logic.Group.Product.UpdateProduct(c, p, content.ID)
	rly.Reply(err, nil)
}
