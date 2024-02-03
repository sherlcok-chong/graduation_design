package logic

import (
	"GraduationDesign/src/dao"
	db "GraduationDesign/src/dao/mysql/sqlc"
	"GraduationDesign/src/dao/mysql/tx"
	"GraduationDesign/src/global"
	mid "GraduationDesign/src/middleware"
	"GraduationDesign/src/model/request"
	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type product struct {
}

func (product) UploadProduct(c *gin.Context, req *request.Product, userID int64) errcode.Err {
	//存信息到商品表
	//文件到媒体表
	//每次存文件需要存关系到商品媒体表
	//商品标签表
	p, _ := strconv.Atoi(req.Price)
	err := dao.Group.Mysql.CreateProduct(c, db.CreateProductParams{
		UserID: userID,
		Price:  int32(p),
		Texts:  req.Describe,
		IsFree: req.IsFree,
		IsLend: req.IsLend,
	})
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	pID, _ := dao.Group.Mysql.GetLastProductID(c)
	for _, v := range req.Media {
		if v != nil {
			_, err := UploadFile(c, v, userID)
			if err != nil {
				global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
				return errcode.ErrServer
			}
			fID, _ := dao.Group.Mysql.GetLastFileID(c)
			err = dao.Group.Mysql.CreateNewMediaProduct(c,
				db.CreateNewMediaProductParams{
					CommodityID: pID,
					FileID:      fID,
				})
			if err != nil {
				global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
				return errcode.ErrServer
			}
		}
	}
	tags := strings.Split(req.Tags, ",")
	for _, v := range tags {
		err := dao.Group.Mysql.CreateTag(c, v)
		if err != nil {
			global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
			return errcode.ErrServer
		}
		tID, _ := dao.Group.Mysql.GetLastTag(c)
		err = dao.Group.Mysql.CreateNewTagProduct(c,
			db.CreateNewTagProductParams{
				ProductID: pID,
				TagID:     tID,
			})
	}
	return nil
}

func (product) GetUserProduct(c *gin.Context, userId int64) ([]tx.Product, errcode.Err) {
	data, err := dao.Group.Mysql.GetUserLendProductTx(c, userId)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	return data, nil
}

func (product) GetInfoProduct(c *gin.Context, userId int64) ([]tx.Product, errcode.Err) {
	data, err := dao.Group.Mysql.GetUserLendProductTx(c, userId)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	return data, nil
}
