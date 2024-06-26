package middleware

import (
	"GraduationDesign/src/dao"
	"context"
	"github.com/0RAJA/Rutils/pkg/app"
	"net/http"
	"strings"

	"GraduationDesign/src/global"
	"GraduationDesign/src/model"
	"GraduationDesign/src/myerr"
	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/0RAJA/Rutils/pkg/token"
	"github.com/gin-gonic/gin"
)

// GetToken 从当前头部获取token
func GetToken(header http.Header) (string, errcode.Err) {
	authorizationHeader := header.Get(global.PvSettings.Token.AuthorizationKey)
	if len(authorizationHeader) == 0 {
		return "", myerr.AuthNotExist
	}
	fields := strings.SplitN(authorizationHeader, " ", 2)
	if len(fields) != 2 || strings.ToLower(fields[0]) != global.PvSettings.Token.AuthorizationType {
		return "", myerr.AuthenticationFailed
	}
	return fields[1], nil
}

// ParseHeader 获取并解析header中token
// 返回 payload,token,err
func ParseHeader(accessToken string) (*token.Payload, string, errcode.Err) {
	payload, err := global.Maker.VerifyToken(accessToken)
	if err != nil {
		if err.Error() == "超时错误" {
			return nil, "", myerr.AuthOverTime
		}
		return nil, "", myerr.AuthenticationFailed
	}
	return payload, accessToken, nil
}

// Auth 鉴权中间件,用于解析并写入token
func Auth() func(c *gin.Context) {
	return func(c *gin.Context) {
		accessToken, merr := GetToken(c.Request.Header)
		if merr != nil {
			c.Next()
			return
		}
		payload, _, merr := ParseHeader(accessToken)
		if merr != nil {
			c.Next()
			return
		}
		content := &model.Content{}
		if err := content.Unmarshal(payload.Content); err != nil {
			c.Next()
			return
		}
		c.Set(global.PvSettings.Token.AuthorizationKey, content)
		c.Next()
	}
}

// GetTokenContent 从当前上下文中获取保存的content内容
func GetTokenContent(c *gin.Context) (*model.Content, bool) {
	val, ok := c.Get(global.PvSettings.Token.AuthorizationKey)
	if !ok {
		return nil, false
	}
	return val.(*model.Content), true
}

//MustUser 必须是用户
func MustUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		rly := app.NewResponse(c)
		val, ok := c.Get(global.PvSettings.Token.AuthorizationKey)
		if !ok {
			rly.Reply(myerr.AuthNotExist)
			c.Abort()
			return
		}
		data := val.(*model.Content)
		if data.Type != model.UserToken {
			rly.Reply(myerr.AuthenticationFailed)
			c.Abort()
			return
		}
		ok, err := dao.Group.Mysql.ExistsUserByID(c, data.ID)
		if err != nil {
			global.Logger.Error(err.Error(), ErrLogMsg(c)...)
			rly.Reply(errcode.ErrServer)
			c.Abort()
			return
		}
		if !ok {
			rly.Reply(myerr.UserNotFound)
			c.Abort()
			return
		}
		c.Next()
	}
}

func MustAccount(accessToken string) (*model.Token, errcode.Err) {
	payload, _, merr := ParseHeader(accessToken)
	if merr != nil {
		return nil, merr
	}
	content := &model.Content{}
	if err := content.Unmarshal(payload.Content); err != nil {
		return nil, myerr.AuthenticationFailed
	}
	if content.Type != model.UserToken {
		return nil, myerr.AuthenticationFailed
	}
	ok, err := dao.Group.Mysql.ExistsUserByID(context.Background(), content.ID)
	if err != nil {
		global.Logger.Error(err.Error())
		return nil, errcode.ErrServer
	}
	if !ok {
		return nil, myerr.UserNotFound
	}
	return &model.Token{
		AccessToken: accessToken,
		Payload:     payload,
		Content:     content,
	}, nil
}
