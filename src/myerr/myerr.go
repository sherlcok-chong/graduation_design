package myerr

import (
	"github.com/0RAJA/Rutils/pkg/app/errcode"
)

var (
	UserNotFound         = errcode.NewErr(2001, "用户不存在")
	PasswordNotValid     = errcode.NewErr(2002, "密码错误")
	EmailSendMany        = errcode.NewErr(2003, "邮件发送频繁，请稍后再试")
	EmailCodeNotValid    = errcode.NewErr(2004, "邮箱验证码校验失败")
	EmailSame            = errcode.NewErr(2005, "邮箱相同")
	EmailExists          = errcode.NewErr(2006, "邮箱已经注册")
	AuthNotExist         = errcode.NewErr(2007, "身份不存在")  // 没找到header上的token
	AuthenticationFailed = errcode.NewErr(2008, "身份验证失败") // token验证失败
	AccountNotFound      = errcode.NewErr(2010, "账号不存在")
	AuthOverTime         = errcode.NewErr(2013, "身份过期")   // token过期
	UserHasAccount       = errcode.NewErr(3005, "用户存在账户") // 删除用户时用户仍存在账户
	FiledStore           = errcode.NewErr(4001, "文件存储失败")
	NoPermissionToDelete = errcode.NewErr(5001, "无权删除")
)
