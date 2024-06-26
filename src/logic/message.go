package logic

import (
	"GraduationDesign/src/dao"
	db "GraduationDesign/src/dao/mysql/sqlc"
	"GraduationDesign/src/global"
	mid "GraduationDesign/src/middleware"
	"GraduationDesign/src/model/chat"
	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/gin-gonic/gin"
)

type ws struct {
}

func (ws) GetNotReadMsg(c *gin.Context, userID int64) ([]chat.NotReadMsg, errcode.Err) {
	tids, err := dao.Group.Mysql.GetUserWhoTalk(c, userID)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	rsp := make([]chat.NotReadMsg, 0, len(tids))
	for _, v := range tids {
		data, _ := dao.Group.Mysql.GetNotReadMsgByUserID(c, db.GetNotReadMsgByUserIDParams{
			Tid: userID,
			Fid: v,
		})
		if len(data) == 0 {
			continue
		}
		userInfo, err := dao.Group.Mysql.GetUserInfoById(c, v)
		if err != nil {
			global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
			return nil, errcode.ErrServer
		}
		r := chat.NotReadMsg{
			UserID:   userInfo.ID,
			UserName: userInfo.Name,
			Avatar:   userInfo.Avatar,
			Msg:      nil,
		}

		r.Msg = make([]chat.MsgSend, 0, len(data))
		for _, v := range data {
			var typ int64
			if v.IsFile {
				typ = 1
			}
			t := chat.MsgSend{
				ID:      v.ID,
				FUserID: v.Fid,
				TUserID: v.Tid,
				MsgType: typ,
				Text:    v.Texts,
				IsRead:  v.IsRead,
			}
			dao.Group.Mysql.ReadUserMessage(c, db.ReadUserMessageParams{
				Fid: v.Fid,
				Tid: v.Tid,
			})
			r.Msg = append(r.Msg, t)
		}
		rsp = append(rsp, r)
	}
	return rsp, nil
}

func (ws) ReadUserMsg(c *gin.Context, fid, tid int64) errcode.Err {
	err := dao.Group.Mysql.ReadUserMessage(c, db.ReadUserMessageParams{
		Fid: fid,
		Tid: tid,
	})
	if err != nil {
		return errcode.ErrServer
	}
	return nil
}
