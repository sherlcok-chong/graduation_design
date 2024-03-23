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
		data, err := dao.Group.Mysql.GetNotReadMsgByUserID(c, db.GetNotReadMsgByUserIDParams{
			Tid: userID,
			Fid: v,
		})
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
			r.Msg = append(r.Msg, t)
		}
		rsp = append(rsp, r)
	}
	return rsp, nil
}
