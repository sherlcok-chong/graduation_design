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
	"github.com/gin-gonic/gin"
)

type comment struct {
}

func (comment) AddComment(c *gin.Context, req *request.AddComment, userID int64) errcode.Err {
	err := dao.Group.Mysql.CreateNewComment(c, db.CreateNewCommentParams{
		UserID:    userID,
		ProductID: req.ProductID,
		Texts:     req.Comment,
	})
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	cID, _ := dao.Group.Mysql.GetLastCommentID(c)
	for _, v := range req.Media {
		if v == nil {
			continue
		}
		_, err := uploadFile(c, v, userID)
		if err != nil {
			global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
			return errcode.ErrServer
		}
		fID, _ := dao.Group.Mysql.GetLastFileID(c)
		err = dao.Group.Mysql.CreateCommentMedias(c, db.CreateCommentMediasParams{
			CommentID: cID,
			FileID:    fID,
		})
		if err != nil {
			global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
			return errcode.ErrServer
		}
	}
	return nil
}

func (comment) DeleteComment(c *gin.Context, cID, userID int64) errcode.Err {
	if err := checkSignC(c, cID, userID); err != nil {
		return err
	}
	err := dao.Group.Mysql.DeleteCommentID(c, cID)
	err = dao.Group.Mysql.DeleteCommentMedia(c, cID)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	return nil
}
func (comment) GetProductComment(c *gin.Context, cID int64) ([]reply.Comment, errcode.Err) {
	data, err := dao.Group.Mysql.GetProductComment(c, cID)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	rsp := make([]reply.Comment, 0, len(data))
	for _, v := range data {
		r := reply.Comment{
			ID:     v.ID,
			UserID: v.UserID,
			Text:   v.Texts,
		}
		fids, err := dao.Group.Mysql.GetCommentMedia(c, v.ID)
		if err != nil {
			global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
			return nil, errcode.ErrServer
		}
		avatar, err := dao.Group.Mysql.GetUserInfoById(c, v.UserID)
		if err != nil {
			global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
			return nil, errcode.ErrServer
		}
		r.Avatar = avatar.Avatar
		r.Username = avatar.Name
		urls := make([]string, 0, len(fids))
		for _, v := range fids {
			url, err := dao.Group.Mysql.GetFileByID(c, v)
			if err != nil {
				global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
				return nil, errcode.ErrServer
			}
			urls = append(urls, url)
		}
		r.Media = urls
		rsp = append(rsp, r)
	}
	return rsp, nil
}
func checkSignC(c *gin.Context, cID, uID int64) errcode.Err {
	v, err := dao.Group.Mysql.GetCommentUser(c, cID)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	if v != uID {
		return myerr.NoPermissionToDelete
	}
	return nil
}
