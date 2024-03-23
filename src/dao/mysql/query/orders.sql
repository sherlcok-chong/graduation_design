-- name: CreateOrder :exec
insert into orders (lend_user_id, borrow_user_id, product_id, unit_price,
                    total_price, completion_time, product_status,
                    start_time, end_time)
values (?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetOrderDetail :one
select *
from orders
where id = ?;

-- name: GetUserLendOrder :many
select *
from orders
where borrow_user_id = ?;

-- name: GetUserBorrowOrder :many
select *
from orders
where lend_user_id = ?;

-- name: GetProductNotFreeTime :many
select start_time, end_time
from orders
where product_id = ?;

-- name: DeleteOrder :exec
delete
from orders
where id = ?;

-- name: UpdateOrderExpress :exec
update orders
set express_number = ?,
    product_status = 1
where id = ?;

-- name: EnsureExpress :exec
update orders
set product_status = 2
where id = ?;

-- name: EnsureRec :exec
update orders
set product_status = 3
where id = ?;
