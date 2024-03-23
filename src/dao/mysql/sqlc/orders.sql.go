// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: orders.sql

package db

import (
	"context"
)

const createOrder = `-- name: CreateOrder :exec
insert into orders (lend_user_id, borrow_user_id, product_id, unit_price,
                    total_price, completion_time, product_status,
                    start_time, end_time)
values (?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type CreateOrderParams struct {
	LendUserID     int64  `json:"lend_user_id"`
	BorrowUserID   int64  `json:"borrow_user_id"`
	ProductID      int64  `json:"product_id"`
	UnitPrice      string `json:"unit_price"`
	TotalPrice     string `json:"total_price"`
	CompletionTime string `json:"completion_time"`
	ProductStatus  int32  `json:"product_status"`
	StartTime      string `json:"start_time"`
	EndTime        string `json:"end_time"`
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) error {
	_, err := q.db.ExecContext(ctx, createOrder,
		arg.LendUserID,
		arg.BorrowUserID,
		arg.ProductID,
		arg.UnitPrice,
		arg.TotalPrice,
		arg.CompletionTime,
		arg.ProductStatus,
		arg.StartTime,
		arg.EndTime,
	)
	return err
}

const deleteOrder = `-- name: DeleteOrder :exec
delete
from orders
where id = ?
`

func (q *Queries) DeleteOrder(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteOrder, id)
	return err
}

const ensureExpress = `-- name: EnsureExpress :exec
update orders
set product_status = 2
where id = ?
`

func (q *Queries) EnsureExpress(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, ensureExpress, id)
	return err
}

const ensureRec = `-- name: EnsureRec :exec
update orders
set product_status = 3
where id = ?
`

func (q *Queries) EnsureRec(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, ensureRec, id)
	return err
}

const getOrderDetail = `-- name: GetOrderDetail :one
select id, lend_user_id, borrow_user_id, product_id, unit_price, total_price, completion_time, product_status, express_number, start_time, end_time
from orders
where id = ?
`

func (q *Queries) GetOrderDetail(ctx context.Context, id int64) (Order, error) {
	row := q.db.QueryRowContext(ctx, getOrderDetail, id)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.LendUserID,
		&i.BorrowUserID,
		&i.ProductID,
		&i.UnitPrice,
		&i.TotalPrice,
		&i.CompletionTime,
		&i.ProductStatus,
		&i.ExpressNumber,
		&i.StartTime,
		&i.EndTime,
	)
	return i, err
}

const getProductNotFreeTime = `-- name: GetProductNotFreeTime :many
select start_time, end_time
from orders
where product_id = ?
`

type GetProductNotFreeTimeRow struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

func (q *Queries) GetProductNotFreeTime(ctx context.Context, productID int64) ([]GetProductNotFreeTimeRow, error) {
	rows, err := q.db.QueryContext(ctx, getProductNotFreeTime, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetProductNotFreeTimeRow{}
	for rows.Next() {
		var i GetProductNotFreeTimeRow
		if err := rows.Scan(&i.StartTime, &i.EndTime); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserBorrowOrder = `-- name: GetUserBorrowOrder :many
select id, lend_user_id, borrow_user_id, product_id, unit_price, total_price, completion_time, product_status, express_number, start_time, end_time
from orders
where lend_user_id = ?
`

func (q *Queries) GetUserBorrowOrder(ctx context.Context, lendUserID int64) ([]Order, error) {
	rows, err := q.db.QueryContext(ctx, getUserBorrowOrder, lendUserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Order{}
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.ID,
			&i.LendUserID,
			&i.BorrowUserID,
			&i.ProductID,
			&i.UnitPrice,
			&i.TotalPrice,
			&i.CompletionTime,
			&i.ProductStatus,
			&i.ExpressNumber,
			&i.StartTime,
			&i.EndTime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserLendOrder = `-- name: GetUserLendOrder :many
select id, lend_user_id, borrow_user_id, product_id, unit_price, total_price, completion_time, product_status, express_number, start_time, end_time
from orders
where borrow_user_id = ?
`

func (q *Queries) GetUserLendOrder(ctx context.Context, borrowUserID int64) ([]Order, error) {
	rows, err := q.db.QueryContext(ctx, getUserLendOrder, borrowUserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Order{}
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.ID,
			&i.LendUserID,
			&i.BorrowUserID,
			&i.ProductID,
			&i.UnitPrice,
			&i.TotalPrice,
			&i.CompletionTime,
			&i.ProductStatus,
			&i.ExpressNumber,
			&i.StartTime,
			&i.EndTime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateOrderExpress = `-- name: UpdateOrderExpress :exec
update orders
set express_number = ?,
    product_status = 1
where id = ?
`

type UpdateOrderExpressParams struct {
	ExpressNumber string `json:"express_number"`
	ID            int64  `json:"id"`
}

func (q *Queries) UpdateOrderExpress(ctx context.Context, arg UpdateOrderExpressParams) error {
	_, err := q.db.ExecContext(ctx, updateOrderExpress, arg.ExpressNumber, arg.ID)
	return err
}
