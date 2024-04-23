package request

import "time"

type Order struct {
	ProductID    int64     `json:"product_id,omitempty" binding:"required"`
	LendUserID   int64     `json:"lend_user_id,omitempty" binding:"required"`
	BorrowUserID int64     `json:"borrow_user_id,omitempty" binding:"required"`
	UintPrice    string    `json:"uint_price,omitempty" binding:"required"`
	TotalPrice   string    `json:"total_price,omitempty" binding:"required"`
	StartTime    time.Time `json:"start_time,omitempty" binding:"required"`
	EndTime      time.Time `json:"end_time,omitempty" binding:"required"`
}

type ChangeOrderStatus struct {
	ID         int64  `json:"id"`
	Status     int32  `json:"status"`
	ExpressNum string `json:"express_num"`
}

type Pay struct {
	OrderID string  `form:"order_id"`
	Price   float32 `form:"price"`
}

type OrderID struct {
	ID int64 `form:"id"`
}
