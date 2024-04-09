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

type file struct {
}

func (file) UpdateUserAvatar(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.UpdateUserAvatar{}
	if err := c.ShouldBind(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.UserToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	rsp, err := logic.Group.File.UpdateUserAvatar(c, params, content.ID)
	rly.Reply(err, rsp)
}

func (file) UploadFile(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.UploadFile{}
	if err := c.ShouldBind(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.UserToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	rsp, err := logic.Group.File.UploadFile(c, params.File, content.ID)
	rly.Reply(err, rsp)
}
