// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO user (name, password, email)
VALUES (?, ?, ?)
`

type CreateUserParams struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser, arg.Name, arg.Password, arg.Email)
	return err
}

const existEmail = `-- name: ExistEmail :one
select exists(select 1 from user where email = ?)
`

func (q *Queries) ExistEmail(ctx context.Context, email string) (bool, error) {
	row := q.db.QueryRowContext(ctx, existEmail, email)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const existsUserByID = `-- name: ExistsUserByID :one
select exists(select 1 from user where id = ?)
`

func (q *Queries) ExistsUserByID(ctx context.Context, id int64) (bool, error) {
	row := q.db.QueryRowContext(ctx, existsUserByID, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
select id, name, email, password, avatar, sign, gender, birthday
from user
where email = ?
limit 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Avatar,
		&i.Sign,
		&i.Gender,
		&i.Birthday,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
select id, name, email, password, avatar, sign, gender, birthday
from user
where name = ?
limit 1
`

func (q *Queries) GetUserByUsername(ctx context.Context, name string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, name)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Avatar,
		&i.Sign,
		&i.Gender,
		&i.Birthday,
	)
	return i, err
}

const getUserInfoById = `-- name: GetUserInfoById :one
select id, name, email, password, avatar, sign, gender, birthday
from user
where id = ?
`

func (q *Queries) GetUserInfoById(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserInfoById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Avatar,
		&i.Sign,
		&i.Gender,
		&i.Birthday,
	)
	return i, err
}

const updateUserAvatar = `-- name: UpdateUserAvatar :exec
update user
set avatar = ?
where id = ?
`

type UpdateUserAvatarParams struct {
	Avatar string `json:"avatar"`
	ID     int64  `json:"id"`
}

func (q *Queries) UpdateUserAvatar(ctx context.Context, arg UpdateUserAvatarParams) error {
	_, err := q.db.ExecContext(ctx, updateUserAvatar, arg.Avatar, arg.ID)
	return err
}

const updateUserInfo = `-- name: UpdateUserInfo :exec
update user
set name     = ?,
    sign     = ?,
    gender   = ?,
    birthday = ?
where id = ?
`

type UpdateUserInfoParams struct {
	Name     string `json:"name"`
	Sign     string `json:"sign"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
	ID       int64  `json:"id"`
}

func (q *Queries) UpdateUserInfo(ctx context.Context, arg UpdateUserInfoParams) error {
	_, err := q.db.ExecContext(ctx, updateUserInfo,
		arg.Name,
		arg.Sign,
		arg.Gender,
		arg.Birthday,
		arg.ID,
	)
	return err
}
