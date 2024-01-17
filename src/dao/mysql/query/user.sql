-- name: ExistEmail :one
select exists(select 1 from user where email = ?);

-- name: CreateUser :exec
INSERT INTO user (name, password, email)
VALUES (?, ?, ?);

-- name: GetUserByUsername :one
select *
from user
where name = ?
limit 1;