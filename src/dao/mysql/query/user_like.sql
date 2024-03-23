-- name: LikeProduct :exec
insert into user_like (user_id, product_id)
VALUES (?, ?);

-- name: DisLikeProduct :exec
delete
from user_like
where user_id = ?
  and product_id = ?;

-- name: CheckUserLike :one
select exists(select 1 from user_like where user_id = ? and product_id = ?);

-- name: GetLikeList :many
select product_id
from user_like
where user_id = ?;

-- name: DeleteLike :exec
delete
from user_like
where product_id = ?;