package logic

import (
	"GraduationDesign/src/dao"
	"GraduationDesign/src/global"
	mid "GraduationDesign/src/middleware"
	"GraduationDesign/src/model/reply"
	"GraduationDesign/src/myerr"
	"GraduationDesign/src/pkg/mark"
	"errors"
	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/0RAJA/Rutils/pkg/utils"
	"github.com/gin-gonic/gin"
)

type email struct {
}

// ExistEmail 是否存在email
func (email) ExistEmail(c *gin.Context, emailStr string) (*reply.ExistEmail, errcode.Err) {
	//ok, err := dao.Group.Redis.ExistEmail(c, emailStr)
	//if err == nil {
	//	return &reply.ExistEmail{Exist: ok}, nil
	//}
	//global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
	ok, err := dao.Group.Mysql.ExistEmail(c, emailStr)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	return &reply.ExistEmail{Exist: ok}, nil
}

// CheckEmailNotExists 判断邮箱是否已经注册
func CheckEmailNotExists(c *gin.Context, emailStr string) errcode.Err {
	result, err := email{}.ExistEmail(c, emailStr)
	if err != nil {
		return err
	}
	if result.Exist {
		return myerr.EmailExists
	}
	return nil
}

// SendEmail 发送邮件
func (email) SendEmail(c *gin.Context, emailStr string) errcode.Err {
	return nil
	// 判断发送频率
	if global.EmailMark.CheckUserExist(emailStr) {
		return myerr.EmailSendMany
	}
	// 异步发送邮件
	global.Worker.SendTask(func() {
		code := utils.RandomString(global.PbSettings.Rule.CodeLength)
		if err := global.EmailMark.SendMail(emailStr, code); err != nil && !errors.Is(err, mark.ErrSendTooMany) {
			global.Logger.Error(err.Error())
		}
	})
	return nil
}
