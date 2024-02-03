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

-- name: GetUserByEmail :one
select *
from user
where email = ?
limit 1;

-- name: UpdateUserInfo :exec
update user
set name     = ?,
    sign     = ?,
    gender   = ?,
    birthday = ?
where id = ?;

-- name: UpdateUserAvatar :exec
update user
set avatar = ?
where id = ?;

-- name: ExistsUserByID :one
select exists(select 1 from user where id = ?);

-- name: GetUserInfoById :one
select *
from user
where id = ?;