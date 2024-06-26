// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: product.sql

package db

import (
	"context"
)

const createNewMediaProduct = `-- name: CreateNewMediaProduct :exec
INSERT INTO commodity_media (commodity_id, file_id)
VALUES (?, ?)
`

type CreateNewMediaProductParams struct {
	CommodityID int64 `json:"commodity_id"`
	FileID      int64 `json:"file_id"`
}

func (q *Queries) CreateNewMediaProduct(ctx context.Context, arg CreateNewMediaProductParams) error {
	_, err := q.db.ExecContext(ctx, createNewMediaProduct, arg.CommodityID, arg.FileID)
	return err
}

const createProduct = `-- name: CreateProduct :exec
insert into commodity (user_id, price, texts, is_free, is_lend, name)
values (?, ?, ?, ?, ?, ?)
`

type CreateProductParams struct {
	UserID int64  `json:"user_id"`
	Price  string `json:"price"`
	Texts  string `json:"texts"`
	IsFree bool   `json:"is_free"`
	IsLend bool   `json:"is_lend"`
	Name   string `json:"name"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) error {
	_, err := q.db.ExecContext(ctx, createProduct,
		arg.UserID,
		arg.Price,
		arg.Texts,
		arg.IsFree,
		arg.IsLend,
		arg.Name,
	)
	return err
}

const deleteFileMedia = `-- name: DeleteFileMedia :exec
delete
from commodity_media
where commodity_id = ?
`

func (q *Queries) DeleteFileMedia(ctx context.Context, commodityID int64) error {
	_, err := q.db.ExecContext(ctx, deleteFileMedia, commodityID)
	return err
}

const deleteProduct = `-- name: DeleteProduct :exec
delete
from commodity
where id = ?
`

func (q *Queries) DeleteProduct(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteProduct, id)
	return err
}

const getLastProductID = `-- name: GetLastProductID :one
SELECT LAST_INSERT_ID()
`

func (q *Queries) GetLastProductID(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getLastProductID)
	var last_insert_id int64
	err := row.Scan(&last_insert_id)
	return last_insert_id, err
}

const getProductByID = `-- name: GetProductByID :one
select id, name, user_id, price, texts, is_free, is_lend
from commodity
where id = ?
`

func (q *Queries) GetProductByID(ctx context.Context, id int64) (Commodity, error) {
	row := q.db.QueryRowContext(ctx, getProductByID, id)
	var i Commodity
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.UserID,
		&i.Price,
		&i.Texts,
		&i.IsFree,
		&i.IsLend,
	)
	return i, err
}

const getProductFirstMedia = `-- name: GetProductFirstMedia :one
select MIN(file_id)
from commodity_media
where commodity_id = ?
`

func (q *Queries) GetProductFirstMedia(ctx context.Context, commodityID int64) (interface{}, error) {
	row := q.db.QueryRowContext(ctx, getProductFirstMedia, commodityID)
	var min interface{}
	err := row.Scan(&min)
	return min, err
}

const getProductInfo = `-- name: GetProductInfo :many
select id, price, name, user_id
from commodity
where is_lend = 1
  and is_free = 0
limit ? offset ?
`

type GetProductInfoParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetProductInfoRow struct {
	ID     int64  `json:"id"`
	Price  string `json:"price"`
	Name   string `json:"name"`
	UserID int64  `json:"user_id"`
}

func (q *Queries) GetProductInfo(ctx context.Context, arg GetProductInfoParams) ([]GetProductInfoRow, error) {
	rows, err := q.db.QueryContext(ctx, getProductInfo, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetProductInfoRow{}
	for rows.Next() {
		var i GetProductInfoRow
		if err := rows.Scan(
			&i.ID,
			&i.Price,
			&i.Name,
			&i.UserID,
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

const getProductLike = `-- name: GetProductLike :one
select id, price, name, user_id,is_free
from commodity
where id = ?
`

type GetProductLikeRow struct {
	ID     int64  `json:"id"`
	Price  string `json:"price"`
	Name   string `json:"name"`
	UserID int64  `json:"user_id"`
	IsFree bool   `json:"is_free"`
}

func (q *Queries) GetProductLike(ctx context.Context, id int64) (GetProductLikeRow, error) {
	row := q.db.QueryRowContext(ctx, getProductLike, id)
	var i GetProductLikeRow
	err := row.Scan(
		&i.ID,
		&i.Price,
		&i.Name,
		&i.UserID,
		&i.IsFree,
	)
	return i, err
}

const getProductMedia = `-- name: GetProductMedia :one
select url
from file
where id = ?
`

func (q *Queries) GetProductMedia(ctx context.Context, id int64) (string, error) {
	row := q.db.QueryRowContext(ctx, getProductMedia, id)
	var url string
	err := row.Scan(&url)
	return url, err
}

const getProductMediaId = `-- name: GetProductMediaId :many
select file_id
from commodity_media
where commodity_id = ?
`

func (q *Queries) GetProductMediaId(ctx context.Context, commodityID int64) ([]int64, error) {
	rows, err := q.db.QueryContext(ctx, getProductMediaId, commodityID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int64{}
	for rows.Next() {
		var file_id int64
		if err := rows.Scan(&file_id); err != nil {
			return nil, err
		}
		items = append(items, file_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserLendProduct = `-- name: GetUserLendProduct :many
select id, name, user_id, price, texts, is_free, is_lend
from commodity
where user_id = ?
  and is_lend = 1
`

func (q *Queries) GetUserLendProduct(ctx context.Context, userID int64) ([]Commodity, error) {
	rows, err := q.db.QueryContext(ctx, getUserLendProduct, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Commodity{}
	for rows.Next() {
		var i Commodity
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.UserID,
			&i.Price,
			&i.Texts,
			&i.IsFree,
			&i.IsLend,
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

const getUserNeedInfo = `-- name: GetUserNeedInfo :many
select id, price, name, user_id
from commodity
where is_lend = 0
  and user_id = ?
`

type GetUserNeedInfoRow struct {
	ID     int64  `json:"id"`
	Price  string `json:"price"`
	Name   string `json:"name"`
	UserID int64  `json:"user_id"`
}

func (q *Queries) GetUserNeedInfo(ctx context.Context, userID int64) ([]GetUserNeedInfoRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserNeedInfo, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetUserNeedInfoRow{}
	for rows.Next() {
		var i GetUserNeedInfoRow
		if err := rows.Scan(
			&i.ID,
			&i.Price,
			&i.Name,
			&i.UserID,
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

const getUserProductInfo = `-- name: GetUserProductInfo :many
select id, price, name, user_id, is_free
from commodity
where is_lend = 1
  and user_id = ?
`

type GetUserProductInfoRow struct {
	ID     int64  `json:"id"`
	Price  string `json:"price"`
	Name   string `json:"name"`
	UserID int64  `json:"user_id"`
	IsFree bool   `json:"is_free"`
}

func (q *Queries) GetUserProductInfo(ctx context.Context, userID int64) ([]GetUserProductInfoRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserProductInfo, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetUserProductInfoRow{}
	for rows.Next() {
		var i GetUserProductInfoRow
		if err := rows.Scan(
			&i.ID,
			&i.Price,
			&i.Name,
			&i.UserID,
			&i.IsFree,
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

const searchLikeText = `-- name: SearchLikeText :many
SELECT
    c.id AS commodity_id,
    c.name AS commodity_name,
    c.price AS commodity_price,
    (SELECT f.url FROM file f INNER JOIN commodity_media cm ON f.id = cm.file_id WHERE cm.commodity_id = c.id LIMIT 1) AS media_url,
    u.name AS username,
    u.avatar,
    c.is_free
FROM
    commodity c
        INNER JOIN
    user u ON c.user_id = u.id
WHERE
    c.name LIKE CONCAT('%',?,'%')
`

type SearchLikeTextRow struct {
	CommodityID    int64  `json:"commodity_id"`
	CommodityName  string `json:"commodity_name"`
	CommodityPrice string `json:"commodity_price"`
	MediaUrl       string `json:"media_url"`
	Username       string `json:"username"`
	Avatar         string `json:"avatar"`
	IsFree         bool   `json:"is_free"`
}

func (q *Queries) SearchLikeText(ctx context.Context, concat interface{}) ([]SearchLikeTextRow, error) {
	rows, err := q.db.QueryContext(ctx, searchLikeText, concat)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SearchLikeTextRow{}
	for rows.Next() {
		var i SearchLikeTextRow
		if err := rows.Scan(
			&i.CommodityID,
			&i.CommodityName,
			&i.CommodityPrice,
			&i.MediaUrl,
			&i.Username,
			&i.Avatar,
			&i.IsFree,
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

const updateProduct = `-- name: UpdateProduct :exec
update commodity
set name    = ?,
    price   = ?,
    texts   = ?,
    is_free = ?
where id = ?
`

type UpdateProductParams struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	Texts  string `json:"texts"`
	IsFree bool   `json:"is_free"`
	ID     int64  `json:"id"`
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) error {
	_, err := q.db.ExecContext(ctx, updateProduct,
		arg.Name,
		arg.Price,
		arg.Texts,
		arg.IsFree,
		arg.ID,
	)
	return err
}
