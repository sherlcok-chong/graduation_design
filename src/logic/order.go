package logic

import (
	"GraduationDesign/src/dao"
	db "GraduationDesign/src/dao/mysql/sqlc"
	"GraduationDesign/src/global"
	mid "GraduationDesign/src/middleware"
	"GraduationDesign/src/model/reply"
	"GraduationDesign/src/model/request"
	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/gin-gonic/gin"
)

type order struct {
}

func (order) CreateOrder(c *gin.Context, req request.Order) errcode.Err {
	err := dao.Group.Mysql.CreateOrder(c, db.CreateOrderParams{
		LendUserID:    req.LendUserID,
		BorrowUserID:  req.BorrowUserID,
		ProductID:     req.ProductID,
		UnitPrice:     req.UintPrice,
		TotalPrice:    req.TotalPrice,
		ProductStatus: 0,
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
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return nil, errcode.ErrServer
	}
	rsp := make([]reply.Order, 0, len(data))
	for _, v := range data {
		r := reply.Order{
			ID:           v.ID,
			ProductID:    v.ProductID,
			LendUserID:   v.LendUserID,
			BorrowUserID: v.BorrowUserID,
			UintPrice:    v.UnitPrice,
			TotalPrice:   v.TotalPrice,
			StartTime:    v.StartTime,
			EndTime:      v.EndTime,
			Status:       v.ProductStatus,
		}
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
	times := make([]reply.LendTime, 0, len(data))
	for _, v := range data {
		t := reply.LendTime{
			Start: v.StartTime,
			End:   v.EndTime,
		}
		times = append(times, t)
	}
	rsp.Times = times
	return rsp, nil
}
func (order) ChangeOrderStatus(c *gin.Context, req request.ChangeOrderStatus) errcode.Err {
	var err error
	if req.Status == 1 {
		err = dao.Group.Mysql.UpdateOrderExpress(c, db.UpdateOrderExpressParams{
			ExpressNumber: req.ExpressNum,
			ID:            req.ProductID,
		})
	} else if req.Status == 2 {
		err = dao.Group.Mysql.EnsureExpress(c, req.ProductID)
	} else if req.Status == 3 {
		err = dao.Group.Mysql.EnsureRec(c, req.ProductID)
	}
	if err != nil {
		global.Logger.Error(err.Error(), mid.ErrLogMsg(c)...)
		return errcode.ErrServer
	}
	return nil
}
