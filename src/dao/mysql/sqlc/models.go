// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import ()

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Sign     string `json:"sign"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
}
