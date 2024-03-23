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

// 定义产品结构体
type product struct {
}

// UploadProduct 上传产品
func (product) UploadProduct(c *gin.Context) {
	// 初始化响应
	rly := app.NewResponse(c)
	// 解析请求参数
	params := &request.Product{}
	if err := c.ShouldBind(params); err != nil {
		// 参数错误处理
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	// 获取令牌内容
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.UserToken {
		// 未授权处理
		rly.Reply(myerr.AuthNotExist)
		return
	}
	// 调用逻辑层处理上传产品逻辑
	err := logic.Group.Product.UploadProduct(c, params, content.ID)
	// 返回处理结果
	rly.Reply(err, nil)
}

// GetUserLendProduct 获取用户出借产品
func (product) GetUserLendProduct(c *gin.Context) {
	// 初始化响应
	rly := app.NewResponse(c)
	// 解析查询参数
	p := &request.ProductInfo{}
	if err := c.ShouldBindQuery(p); err != nil {
		// 参数错误处理
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	// 获取令牌内容
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.UserToken {
		// 未授权处理
		rly.Reply(myerr.AuthNotExist)
		return
	}
	// 调用逻辑层获取用户出借产品信息
	rsp, err := logic.Group.Product.GetUserProduct(c, content.ID)
	// 返回处理结果
	rly.Reply(err, rsp)
}

// GetUserNeedProduct 获取用户需求产品
func (product) GetUserNeedProduct(c *gin.Context) {
	// 初始化响应
	rly := app.NewResponse(c)
	// 解析查询参数
	p := &request.ProductInfo{}
	if err := c.ShouldBindQuery(p); err != nil {
		// 参数错误处理
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	// 获取令牌内容
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.UserToken {
		// 未授权处理
		rly.Reply(myerr.AuthNotExist)
		return
	}
	// 调用逻辑层获取用户需求产品信息
	rsp, err := logic.Group.Product.GetUserNeed(c, content.ID)
	// 返回处理结果
	rly.Reply(err, rsp)
}

// GetProductDetails 获取产品详情
func (product) GetProductDetails(c *gin.Context) {
	// 初始化响应
	rly := app.NewResponse(c)
	// 解析查询参数
	p := &request.ProductDetails{}
	if err := c.ShouldBindQuery(p); err != nil {
		// 参数错误处理
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	// 获取令牌内容
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.UserToken {
		// 未授权处理
		rly.Reply(myerr.AuthNotExist)
		return
	}
	// 调用逻辑层获取产品详情信息
	rsp, err := logic.Group.Product.GetProductDetails(c, p.ID, content.ID)
	// 返回处理结果
	rly.Reply(err, rsp)
}

// GetProductInfo 获取产品信息
func (product) GetProductInfo(c *gin.Context) {
	// 初始化响应
	rly := app.NewResponse(c)
	// 解析查询参数
	p := &request.ProductInfo{}
	if err := c.ShouldBindQuery(p); err != nil {
		// 参数错误处理
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	// 获取令牌内容
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.UserToken {
		// 未授权处理
		rly.Reply(myerr.AuthNotExist)
		return
	}
	// 调用逻辑层获取产品信息
	rsp, err := logic.Group.Product.GetProductInfo(c, p.Offset)
	// 返回处理结果
	rly.Reply(err, rsp)
}

// DeleteProduct 删除产品
func (product) DeleteProduct(c *gin.Context) {
	// 初始化响应
	rly := app.NewResponse(c)
	// 解析JSON参数
	p := &request.ProductID{}
	if err := c.ShouldBindJSON(p); err != nil {
		// 参数错误处理
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	// 获取令牌内容
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.UserToken {
		// 未授权处理
		rly.Reply(myerr.AuthNotExist)
		return
	}
	// 调用逻辑层处理删除产品逻辑
	err := logic.Group.Product.DeleteProduct(c, p.ID, content.ID)
	// 返回处理结果
	rly.Reply(err, nil)
}

// UpdateProduct 更新产品
func (product) UpdateProduct(c *gin.Context) {
	// 初始化响应
	rly := app.NewResponse(c)
	// 解析参数
	p := &request.UpdateProduct{}
	if err := c.ShouldBind(p); err != nil {
		// 参数错误处理
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	// 获取令牌内容
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.UserToken {
		// 未授权处理
		rly.Reply(myerr.AuthNotExist)
		return
	}
	// 调用逻辑层处理更新产品逻辑
	err := logic.Group.Product.UpdateProduct(c, p, content.ID)
	// 返回处理结果
	rly.Reply(err, nil)
}

// ChangeLikeProduct 改变产品喜欢状态
func (product) ChangeLikeProduct(c *gin.Context) {
	// 初始化响应
	rly := app.NewResponse(c)
	// 解析参数
	p := &request.ProductID{}
	if err := c.ShouldBind(p); err != nil {
		// 参数错误处理
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	// 获取令牌内容
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.UserToken {
		// 未授权处理
		rly.Reply(myerr.AuthNotExist)
		return
	}
	// 调用逻辑层处理改变产品喜欢状态逻辑

	err := logic.Group.Product.ChangeLikeProduct(c, content.ID, p.ID)
	rly.Reply(err)
}

func (product) GetLikeList(c *gin.Context) {
	rly := app.NewResponse(c)
	// 获取令牌内容
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.UserToken {
		// 未授权处理
		rly.Reply(myerr.AuthNotExist)
		return
	}
	data, err := logic.Group.Product.GetLikeList(c, content.ID)
	rly.Reply(err, data)
}

func (product) GetProductBusyTime(c *gin.Context) {
	rly := app.NewResponse(c)
	pID := &request.ProductID{}
	if err := c.ShouldBindQuery(pID); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.UserToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	rsp, err := logic.Group.Order.LendBusyTime(c, pID.ID)
	rly.Reply(err, rsp)
}
