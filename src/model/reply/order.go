package reply

type Order struct {
	ID           int64  `json:"id"`
	ProductID    int64  `json:"product_id,omitempty"`
	LendUserID   int64  `json:"lend_user_id,omitempty"`
	BorrowUserID int64  `json:"borrow_user_id,omitempty"`
	UintPrice    string `json:"uint_price,omitempty"`
	TotalPrice   string `json:"total_price,omitempty"`
	StartTime    string `json:"start_time,omitempty"`
	EndTime      string `json:"end_time,omitempty"`
	Status       int32  `json:"status"`
}

type BusyTime struct {
	ID    int64      `json:"id,omitempty"`
	Times []LendTime `json:"times,omitempty"`
}

type LendTime struct {
	Start string `json:"start,omitempty"`
	End   string `json:"end,omitempty"`
}
