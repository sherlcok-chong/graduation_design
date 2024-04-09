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
	err := dao.Group.Mysql.CreateProduct(c, db.CreateProductParams{
		Name:   req.Name,
		UserID: userID,
		Price:  req.Price,
		Texts:  req.Describe,
		IsFree: req.IsFree,
		IsLend: req.IsLend == 1,
	})
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	pID, _ := dao.Group.Mysql.GetLastProductID(c)
	for _, v := range req.Media {
		if v != nil {
			_, err := uploadFile(c, v, userID)
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
	tgs := strings.Split(req.Tags, ",")
	for _, v := range tgs {
		t, _ := strconv.ParseInt(v, 10, 64)
		err = dao.Group.Mysql.CreateNewTagProduct(c,
			db.CreateNewTagProductParams{
				ProductID: pID,
				TagID:     t,
			})
	}
	return nil
}

func (product) GetUserProduct(c *gin.Context, userId int64) ([]reply.ProductInfo, errcode.Err) {
	data, err := dao.Group.Mysql.GetUserLendProductTx(c, userId)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	return data, nil
}
func (product) GetUserNeed(c *gin.Context, userId int64) ([]reply.ProductInfo, errcode.Err) {
	data, err := dao.Group.Mysql.GetUserNeedProductTx(c, userId)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	return data, nil
}
func (product) GetProductDetails(c *gin.Context, pId, uID int64) (reply.Product, errcode.Err) {
	data, err := dao.Group.Mysql.GetProductDetailsTX(c, pId, uID)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return reply.Product{}, errcode.ErrServer
	}
	return data, nil
}

func (product) GetProductInfo(c *gin.Context, offset int32) ([]reply.ProductInfo, errcode.Err) {
	data, err := dao.Group.Mysql.GetProductInfoTx(c, offset)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	return data, nil
}

func (product) DeleteProduct(c *gin.Context, productID, userID int64) errcode.Err {
	if err := checkSignP(c, productID, userID); err != nil {
		return err
	}
	err := dao.Group.Mysql.DeleteProduct(c, productID)
	err = dao.Group.Mysql.DeleteLike(c, productID)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	return nil
}

func (product) UpdateProduct(c *gin.Context, req *request.UpdateProduct, userID int64) errcode.Err {
	if err := checkSignP(c, req.ID, userID); err != nil {
		return err
	}
	if err := deleteProductFileWithOSS(c, req.ID); err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	err := dao.Group.Mysql.UpdateProduct(c, db.UpdateProductParams{
		Name:   req.Name,
		Price:  req.Price,
		Texts:  req.Describe,
		IsFree: req.IsFree,
		ID:     req.ID,
	})
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	for _, v := range req.Media {
		if v != nil {
			_, err := uploadFile(c, v, userID)
			if err != nil {
				global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
				return errcode.ErrServer
			}
			fID, _ := dao.Group.Mysql.GetLastFileID(c)
			err = dao.Group.Mysql.CreateNewMediaProduct(c,
				db.CreateNewMediaProductParams{
					CommodityID: req.ID,
					FileID:      fID,
				})
			if err != nil {
				global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
				return errcode.ErrServer
			}
		}
	}
	tgs := strings.Split(req.Tags, ",")
	for _, v := range tgs {
		t, _ := strconv.ParseInt(v, 10, 64)
		err = dao.Group.Mysql.CreateNewTagProduct(c,
			db.CreateNewTagProductParams{
				ProductID: req.ID,
				TagID:     t,
			})
	}
	return nil
}

func (product) ChangeLikeProduct(c *gin.Context, uID, pID int64) errcode.Err {
	f, err := dao.Group.Mysql.CheckUserLike(c, db.CheckUserLikeParams{
		UserID:    uID,
		ProductID: pID,
	})
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	if f {
		err = dao.Group.Mysql.DisLikeProduct(c, db.DisLikeProductParams{
			UserID:    uID,
			ProductID: pID,
		})
	} else {
		err = dao.Group.Mysql.LikeProduct(c, db.LikeProductParams{
			UserID:    uID,
			ProductID: pID,
		})
	}
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	return nil
}

func (product) GetLikeList(c *gin.Context, uID int64) ([]reply.ProductInfo, errcode.Err) {
	pids, err := dao.Group.Mysql.GetLikeList(c, uID)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	rsp := make([]reply.ProductInfo, 0, len(pids))
	for _, v := range pids {
		p := reply.ProductInfo{}
		data, err := dao.Group.Mysql.GetProductLike(c, v)
		if err != nil {
			global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
			return nil, errcode.ErrServer
		}
		p = reply.ProductInfo{
			ID:       v,
			Name:     data.Name,
			Price:    data.Price,
			Media:    "",
			UserName: "",
			Avatar:   "",
			IsFree:   data.IsFree,
		}
		ud, err := dao.Group.Mysql.GetUserInfoById(c, data.UserID)
		p.UserName = ud.Name
		p.Avatar = ud.Avatar
		d, err := dao.Group.Mysql.GetProductFirstMedia(c, v)
		if err != nil {
			global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
			return nil, errcode.ErrServer
		}
		id := d.(int64)
		media, err := dao.Group.Mysql.GetFileByID(c, id)
		if err != nil {
			global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
			return nil, errcode.ErrServer
		}
		p.Media = media
		rsp = append(rsp, p)
	}
	return rsp, nil
}

func (product) SearchTag(c *gin.Context, tagID int64) (reply.SearchTags, errcode.Err) {
	exsit, err := dao.Group.Mysql.ExistsTags(c, tagID)
	if err != nil || !exsit {
		return reply.SearchTags{}, errcode.ErrServer
	}
	data, err := dao.Group.Mysql.GetTagsProduct(c, tagID)
	rsp := make([]reply.ProductInfo, 0, len(data))
	for _, v := range data {
		t := reply.ProductInfo{
			ID:       v.CommodityID,
			Name:     v.CommodityName,
			Price:    v.CommodityPrice,
			Media:    v.MediaUrl,
			UserName: v.Username,
			Avatar:   v.Avatar,
			IsFree:   v.IsFree,
		}
		rsp = append(rsp, t)
	}
	tagname, _ := dao.Group.Mysql.GetTagName(c, tagID)
	return reply.SearchTags{
		TagName:     tagname,
		Commodities: rsp,
	}, nil
}

func (product) SearchText(c *gin.Context, text string) ([]reply.ProductInfo, errcode.Err) {
	data, err := dao.Group.Mysql.SearchLikeText(c, text)
	if err != nil {
		return nil, errcode.ErrServer
	}
	rsp := make([]reply.ProductInfo, 0, len(data))
	for _, v := range data {
		t := reply.ProductInfo{
			ID:       v.CommodityID,
			Name:     v.CommodityName,
			Price:    v.CommodityPrice,
			Media:    v.MediaUrl,
			UserName: v.Username,
			Avatar:   v.Avatar,
			IsFree:   v.IsFree,
		}
		rsp = append(rsp, t)
	}
	return rsp, nil
}
func deleteProductFileWithOSS(c *gin.Context, ID int64) error {
	data, err := dao.Group.Mysql.GetProductMediaId(c, ID)
	if err != nil {
		return err
	}
	keys := make([]string, len(data))
	for _, v := range data {
		key, err := dao.Group.Mysql.GetKeyByID(c, v)
		if err != nil {
			return err
		}
		keys = append(keys, key)
	}
	for _, v := range keys {
		_, err := global.OSS.DeleteFile(v)
		if err != nil {
			global.Logger.Error("delete oss error", mid.ErrLogMsg(c)...)
		}
	}
	for _, v := range data {
		if err = dao.Group.Mysql.DeleteFileByID(c, v); err != nil {
			return err
		}
	}
	return dao.Group.Mysql.DeleteFileMedia(c, ID)
}

func checkSignP(c *gin.Context, cID, uID int64) errcode.Err {
	v, err := dao.Group.Mysql.GetProductByID(c, cID)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	if v.UserID != uID {
		return myerr.NoPermissionToDelete
	}
	return nil
}
