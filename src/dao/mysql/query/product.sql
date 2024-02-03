-- name: CreateProduct :exec
insert into commodity (user_id, price, texts, is_free, is_lend)
values (?, ?, ?, ?, ?);

-- name: DeleteProduct :exec
delete
from commodity
where id = ?;

-- name: GetUserLendProduct :many
select *
from commodity
where user_id = ?
  and is_lend = 1;

-- name: CreateNewMediaProduct :exec
INSERT INTO commodity_media (commodity_id, file_id)
VALUES (?, ?);

-- name: GetProductMediaId :many
select file_id
from commodity_media
where commodity_id = ?;

-- name: GetProductMedia :one
select url
from file
where id = ?;

-- name: GetLastProductID :one
SELECT LAST_INSERT_ID();