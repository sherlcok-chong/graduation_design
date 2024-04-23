package reply

type Order struct {
	ID             int64  `json:"id"`
	OrderID        string `json:"order_id"`
	ProductID      int64  `json:"product_id,omitempty"`
	LendUserID     int64  `json:"lend_user_id,omitempty"`
	LendUserName   string `json:"lend_user_name"`
	BorrowUserID   int64  `json:"borrow_user_id,omitempty"`
	BorrowUserName string `json:"borrow_user_name"`
	UintPrice      string `json:"uint_price,omitempty"`
	TotalPrice     string `json:"total_price,omitempty"`
	StartTime      string `json:"start_time,omitempty"`
	EndTime        string `json:"end_time,omitempty"`
	ProductMedia   string `json:"product_media"`
	Status         int32  `json:"status"`
}

type BusyTime struct {
	ID    int64      `json:"id,omitempty"`
	Times []LendTime `json:"times,omitempty"`
}

type LendTime struct {
	Start string `json:"start,omitempty"`
	End   string `json:"end,omitempty"`
}

type OrderID struct {
	ID string `json:"order_id"`
}

type TraceItem struct {
	AcceptTime    string
	AcceptStation string
}
