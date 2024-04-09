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

type comment struct {
}

func (comment) AddComment(c *gin.Context) {
	rly := app.NewResponse(c)
	req := &request.AddComment{}
	if err := c.ShouldBind(req); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.UserToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	err := logic.Group.Comment.AddComment(c, req, content.ID)
	rly.Reply(err, nil)
}
func (comment) DeleteComment(c *gin.Context) {
	rly := app.NewResponse(c)
	req := &request.DeleteComment{}
	if err := c.ShouldBind(req); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.UserToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	err := logic.Group.Comment.DeleteComment(c, req.ID, content.ID)
	rly.Reply(err)
}

func (comment) GetProductComment(c *gin.Context) {
	rly := app.NewResponse(c)
	req := &request.GetProductComment{}
	if err := c.ShouldBindQuery(req); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.UserToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	data, err := logic.Group.Comment.GetProductComment(c, req.ID)
	rly.Reply(err, data)
}
