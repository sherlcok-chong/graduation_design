-- name: CreateProduct :exec
insert into commodity (user_id, price, texts, is_free, is_lend, name)
values (?, ?, ?, ?, ?, ?);

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

-- name: GetProductInfo :many
select id, price, name, user_id
from commodity
where is_lend = 1
  and is_free = 1
limit 10 offset ?;

-- name: GetProductFirstMedia :one
select MIN(file_id)
from commodity_media
where commodity_id = ?;

-- name: GetProductByID :one
select *
from commodity
where id = ?;

-- name: GetUserProductInfo :many
select id, price, name, user_id, is_free
from commodity
where is_lend = 1
  and user_id = ?;

-- name: GetUserNeedInfo :many
select id, price, name, user_id
from commodity
where is_lend = 0
  and user_id = ?;

-- name: DeleteFileMedia :exec
delete
from commodity_media
where commodity_id = ?;

-- name: UpdateProduct :exec
update commodity
set name    = ?,
    price   = ?,
    texts   = ?,
    is_free = ?
where id = ?;