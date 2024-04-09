package logic

import (
	"GraduationDesign/src/dao"
	db "GraduationDesign/src/dao/mysql/sqlc"
	"GraduationDesign/src/global"
	mid "GraduationDesign/src/middleware"
	"GraduationDesign/src/model/reply"
	"GraduationDesign/src/model/request"
	"GraduationDesign/src/myerr"
	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type file struct {
}

func (file) UpdateUserAvatar(c *gin.Context, req *request.UpdateUserAvatar, userId int64) (*reply.UpdateUserAvatar, errcode.Err) {
	var url, key string
	var err error
	options := []oss.Option{
		oss.ContentDisposition("inline"),
	}
	if req.Avatar != nil {
		url, key, err = global.OSS.UploadFile(req.Avatar, options)
		if err != nil {
			return nil, myerr.FiledStore
		}
	}
	err = dao.Group.Mysql.UpdateUserAvatarTx(c, key, url, userId)
	if err != nil {
		return nil, errcode.ErrServer
	}
	return &reply.UpdateUserAvatar{
		UserID: userId,
		Url:    url,
	}, nil
}

func (file) UploadFile(c *gin.Context, f *multipart.FileHeader, userID int64) (string, errcode.Err) {
	url, err := uploadFile(c, f, userID)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return "", errcode.ErrServer
	}
	return url, nil
}
func uploadFile(c *gin.Context, f *multipart.FileHeader, userID int64) (string, error) {
	var url, key string
	var err error
	options := []oss.Option{
		oss.ContentDisposition("inline"),
	}
	if f != nil {
		url, key, err = global.OSS.UploadFile(f, options)
		if err != nil {
			global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
			return "nil", myerr.FiledStore
		}
	}
	err = dao.Group.Mysql.CreateFile(c, db.CreateFileParams{
		Filename: f.Filename,
		FileKey:  key,
		Url:      url,
		Userid:   userID,
	})
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return "nil", myerr.FiledStore
	}
	return url, nil
}
