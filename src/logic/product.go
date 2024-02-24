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
	for _, v := range req.Tags {
		err = dao.Group.Mysql.CreateNewTagProduct(c,
			db.CreateNewTagProductParams{
				ProductID: pID,
				TagID:     v,
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
func (product) GetProductDetails(c *gin.Context, userId int64) (reply.Product, errcode.Err) {
	data, err := dao.Group.Mysql.GetProductDetailsTX(c, userId)
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
	for _, v := range req.Tags {
		err = dao.Group.Mysql.CreateNewTagProduct(c,
			db.CreateNewTagProductParams{
				ProductID: req.ID,
				TagID:     v,
			})
	}
	return nil
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
