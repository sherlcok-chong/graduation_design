package logic

import (
	"GraduationDesign/src/dao"
	db "GraduationDesign/src/dao/mysql/sqlc"
	"GraduationDesign/src/global"
	mid "GraduationDesign/src/middleware"
	"GraduationDesign/src/model/reply"
	"GraduationDesign/src/model/request"
	"GraduationDesign/src/pkg/tool"
	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/0RAJA/Rutils/pkg/times"
	"github.com/gin-gonic/gin"
)

type order struct {
}

func (order) CreateOrder(c *gin.Context, req request.Order, orderID string) errcode.Err {
	err := dao.Group.Mysql.CreateOrder(c, db.CreateOrderParams{
		OrderID:       orderID,
		LendUserID:    req.LendUserID,
		BorrowUserID:  req.BorrowUserID,
		ProductID:     req.ProductID,
		UnitPrice:     req.UintPrice,
		TotalPrice:    req.TotalPrice,
		ProductStatus: -1,
		StartTime:     req.StartTime,
		EndTime:       req.EndTime,
	})
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	return nil
}

func (order) GetOrderList(c *gin.Context, userID int64) ([]reply.Order, errcode.Err) {
	data, err := dao.Group.Mysql.GetUserLendOrder(c, userID)
	d2, err := dao.Group.Mysql.GetUserBorrowOrder(c, userID)
	data = append(data, d2...)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	rsp := make([]reply.Order, 0, len(data))
	for _, v := range data {
		r := reply.Order{
			ID:           v.ID,
			OrderID:      v.OrderID,
			ProductID:    v.ProductID,
			LendUserID:   v.LendUserID,
			BorrowUserID: v.BorrowUserID,
			UintPrice:    v.UnitPrice,
			TotalPrice:   v.TotalPrice,
			Status:       v.ProductStatus,
		}
		d, err := dao.Group.Mysql.GetUserInfoById(c, v.LendUserID)
		if err != nil {
			return nil, errcode.ErrServer
		}
		r.LendUserName = d.Name
		d, err = dao.Group.Mysql.GetUserInfoById(c, v.BorrowUserID)
		if err != nil {
			return nil, errcode.ErrServer
		}
		r.BorrowUserName = d.Name
		med, err := dao.Group.Mysql.GetProductFirstMedia(c, v.ProductID)
		if err != nil {
			return nil, errcode.ErrServer
		}
		id := med.(int64)
		media, err := dao.Group.Mysql.GetFileByID(c, id)
		if err != nil {
			return nil, errcode.ErrServer
		}
		r.ProductMedia = media
		r.StartTime = times.ParseDataToStr(v.StartTime)
		r.EndTime = times.ParseDataToStr(v.EndTime)
		rsp = append(rsp, r)
	}
	return rsp, nil
}

func (order) LendBusyTime(c *gin.Context, pid int64) (reply.BusyTime, errcode.Err) {
	data, err := dao.Group.Mysql.GetProductNotFreeTime(c, pid)
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return reply.BusyTime{}, errcode.ErrServer
	}
	rsp := reply.BusyTime{
		ID:    pid,
		Times: nil,
	}
	tim := make([]reply.LendTime, 0, len(data))
	for _, v := range data {
		t := reply.LendTime{
			Start: times.ParseDataToStr(v.StartTime),
			End:   times.ParseDataToStr(v.EndTime),
		}
		tim = append(tim, t)
	}
	rsp.Times = tim
	return rsp, nil
}
func (order) ChangeOrderStatus(c *gin.Context, req *request.ChangeOrderStatus) errcode.Err {
	var err error
	if req.Status == 1 {
		err = dao.Group.Mysql.UpdateOrderExpress(c, db.UpdateOrderExpressParams{
			ExpressNumber: req.ExpressNum,
			ID:            req.ID,
		})
	} else if req.Status == 2 {
		err = dao.Group.Mysql.EnsureExpress(c, req.ID)
	} else if req.Status == 3 {
		err = dao.Group.Mysql.EnsureRec(c, req.ID)
	}
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	return nil
}

func (order) QueryExpress(c *gin.Context, orderID int64) ([]reply.TraceItem, errcode.Err) {
	expressNum, err := dao.Group.Mysql.GetOrderExpressNum(c, orderID)
	if err != nil {
		return nil, errcode.ErrServer
	}
	data, err := tool.KdnTraces(tool.SHIP_CODE_YUNDA, expressNum)
	if err != nil {
		return nil, errcode.ErrServer
	}
	return data.Traces, nil
}
