package logic

import (
	"GraduationDesign/src/dao"
	db "GraduationDesign/src/dao/mysql/sqlc"
	"GraduationDesign/src/global"
	mid "GraduationDesign/src/middleware"
	"GraduationDesign/src/model"
	"GraduationDesign/src/model/common"
	"GraduationDesign/src/model/reply"
	"GraduationDesign/src/model/request"
	"GraduationDesign/src/myerr"
	"database/sql"
	"errors"
	"github.com/0RAJA/Rutils/pkg/app/errcode"
	encode "github.com/0RAJA/Rutils/pkg/password"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

type user struct {
}

// getUserInfoByID 通过ID获取用户信息
// 参数：userID 用户ID
// 成功: 用户信息,nil
// 失败: 打印日志 myerr.UserNotFound,errcode.ErrServer
//func getUserInfoByID(c *gin.Context, userID int64) (*db.User, errcode.Err) {
//	userInfo, err := dao.Group.DB.GetUserByID(c, userID)
//	if err != nil {
//		if errors.Is(err, pgx.ErrNoRows) {
//			return nil, myerr.UserNotFound
//		}
//		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
//		return nil, errcode.ErrServer
//	}
//	return userInfo, nil
//}

//getUserInfoByEmail 通过email获取用户信息
//参数：email 邮箱
//成功: 用户信息,nil
//失败: 打印日志 myerr.UserNotFound,errcode.ErrServer
func getUserInfoByEmail(c *gin.Context, email string) (db.User, errcode.Err) {
	userInfo, err := dao.Group.Mysql.GetUserByEmail(c, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return db.User{}, myerr.UserNotFound
		}
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return db.User{}, errcode.ErrServer
	}
	return userInfo, nil
}

func (user) Register(c *gin.Context, emailStr, pwd, code string) (*reply.Register, errcode.Err) {
	// 判断邮箱没有被注册
	if err := CheckEmailNotExists(c, emailStr); err != nil {
		return nil, err
	}
	// 校验验证码
	if !global.EmailMark.CheckCode(emailStr, code) {
		return nil, myerr.EmailCodeNotValid
	}
	hashPassword, err := encode.HashPassword(pwd)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	err = dao.Group.Mysql.CreateUser(c, db.CreateUserParams{
		Name:     emailStr,
		Password: hashPassword,
		Email:    emailStr,
	})
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	// 添加邮箱到redis
	if err := dao.Group.Redis.AddEmails(c, emailStr); err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		reTry("addEmail:"+emailStr, func() error { return dao.Group.Redis.AddEmails(c, emailStr) })
	}
	u, err := dao.Group.Mysql.GetUserByUsername(c, emailStr)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	// 创建token
	userToken, payload, err := newToken(model.UserToken, u.ID)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	return &reply.Register{
		UserInfo: reply.UserInfo{
			ID:    u.ID,
			Email: u.Email,
		},
		UserToken: common.Token{
			Token:     userToken,
			ExpiredAt: payload.ExpiredAt,
		},
	}, nil
}

//func (user) DeleteUser(c *gin.Context, userID int64) errcode.Err {
//	userInfo, merr := getUserInfoByID(c, userID)
//	if merr != nil {
//		return merr
//	}
//	accountNum, err := dao.Group.DB.CountAccountByUserID(c, userID)
//	if err != nil {
//		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
//		return errcode.ErrServer
//	}
//	if accountNum > 0 {
//		return myerr.UserHasAccount
//	}
//	if err := dao.Group.DB.DeleteUser(c, userID); err != nil {
//		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
//		return errcode.ErrServer
//	}
//	// 从redis删除邮箱
//	if err := dao.Group.Redis.DeleteEmail(c, userInfo.Email); err != nil {
//		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
//		reTry("delEmail:"+userInfo.Email, func() error { return dao.Group.Redis.DeleteEmail(c, userInfo.Email) })
//	}
//	return nil
//}

//func (user) UpdateUserEmail(c *gin.Context, userID int64, newEmail, code string) errcode.Err {
//	// 判断邮箱是否更改
//	userInfo, err := getUserInfoByID(c, userID)
//	if err != nil {
//		return err
//	}
//	if userInfo.Email == newEmail {
//		return myerr.EmailSame
//	}
//	// 判断邮箱没有被注册
//	if err := CheckEmailNotExists(c, newEmail); err != nil {
//		return err
//	}
//	// 校验验证码
//	if !global.EmailMark.CheckCode(newEmail, code) {
//		return myerr.EmailCodeNotValid
//	}
//	if err := dao.Group.DB.UpdateUser(c, &db.UpdateUserParams{
//		Email:    newEmail,
//		Password: userInfo.Password,
//		ID:       userID,
//	}); err != nil {
//		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
//		return errcode.ErrServer
//	}
//	// 更新邮箱到redis
//	if err := dao.Group.Redis.UpdateEmail(c, userInfo.Email, newEmail); err != nil {
//		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
//		reTry(fmt.Sprintf("updateEmail:%s,%s", userInfo.Email, newEmail), func() error { return dao.Group.Redis.UpdateEmail(c, userInfo.Email, newEmail) })
//	}
//	// 推送更改邮箱通知
//	accessToken, _ := mid.GetToken(c.Request.Header)
//	global.Worker.SendTask(task.UpdateEmail(accessToken, userID, newEmail))
//	return nil
//}

//func (user) UpdateUserPassword(c *gin.Context, userID int64, code, newPwd string) errcode.Err {
//	userInfo, merr := getUserInfoByID(c, userID)
//	if merr != nil {
//		return merr
//	}
//	// 校验验证码
//	if !global.EmailMark.CheckCode(userInfo.Email, code) {
//		return myerr.EmailCodeNotValid
//	}
//	hashPassword, err := encode.HashPassword(newPwd)
//	if err != nil {
//		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
//		return errcode.ErrServer
//	}
//	// 更新
//	if err := dao.Group.DB.UpdateUser(c, &db.UpdateUserParams{
//		Email:    userInfo.Email,
//		Password: hashPassword,
//		ID:       userID,
//	}); err != nil {
//		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
//		return errcode.ErrServer
//	}
//	return nil
//}

func (user) Login(c *gin.Context, email, pwd string) (*reply.Login, errcode.Err) {
	userInfo, merr := getUserInfoByEmail(c, email)
	if merr != nil {
		return nil, merr
	}
	if err := encode.CheckPassword(pwd, userInfo.Password); err != nil {
		return nil, myerr.PasswordNotValid
	}
	// 创建token
	userToken, payload, err := newToken(model.UserToken, userInfo.ID)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	return &reply.Login{
		UserInfo: reply.UserInfo{
			ID:    userInfo.ID,
			Email: userInfo.Email,
		},
		UserToken: common.Token{
			Token:     userToken,
			ExpiredAt: payload.ExpiredAt,
		},
	}, nil
}

func (user) UpdateUserInfo(c *gin.Context, userinfo *request.UpdateUserInfo, userId int64) (myErr errcode.Err) {
	var err error
	err = dao.Group.Mysql.UpdateUserInfo(c, db.UpdateUserInfoParams{
		Name:     userinfo.Name,
		Sign:     userinfo.Sign,
		Gender:   userinfo.Gender,
		Birthday: userinfo.Birthday,
		ID:       userId,
	})
	if err != nil {
		myErr = errcode.ErrServer
	}
	return
}

func (user) GetUserInfo(c *gin.Context, userId int64) (rsp *reply.UserData, myErr errcode.Err) {
	data, err := dao.Group.Mysql.GetUserInfoById(c, userId)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	return &reply.UserData{
		UserInfo: reply.UserInfo{
			ID:    data.ID,
			Email: data.Email,
		},
		Name:     data.Name,
		Avatar:   data.Avatar,
		Sign:     data.Sign,
		Gender:   data.Gender,
		Birthday: data.Birthday,
		Address:  data.Address.String,
	}, nil
}

func (user) UpdateUserAddress(c *gin.Context, uID int64, address string) errcode.Err {
	err := dao.Group.Mysql.AddAddressByID(c, db.AddAddressByIDParams{
		Address: sql.NullString{
			String: address,
			Valid:  true,
		},
		ID: uID,
	})
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	return nil
}
