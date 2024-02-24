package logic

import (
	"GraduationDesign/src/dao"
	"GraduationDesign/src/global"
	mid "GraduationDesign/src/middleware"
	"GraduationDesign/src/model/reply"
	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/gin-gonic/gin"
)

type tags struct {
}

func (tags) GetAllTags(c *gin.Context) ([]reply.Tags, errcode.Err) {
	data, err := dao.Group.Mysql.GetAllTags(c)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	ts := make([]reply.Tags, 0, len(data))
	for _, v := range data {
		d := reply.Tags{
			ID:  v.TagID,
			Tag: v.TagName,
		}
		ts = append(ts, d)
	}
	return ts, nil
}
