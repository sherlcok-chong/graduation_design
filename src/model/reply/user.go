package reply

import (
	"GraduationDesign/src/model/common"
)

// UserInfo 用户信息
type UserInfo struct {
	ID    int64  `json:"id"`    // user id
	Email string `json:"email"` // 邮箱
}

type Register struct {
	UserInfo  UserInfo     `json:"user_info"`  // 用户信息
	UserToken common.Token `json:"user_token"` // 用户令牌
}

type Login struct {
	UserInfo  UserInfo     `json:"user_info"`  // 用户信息
	UserToken common.Token `json:"user_token"` // 用户令牌
}

type UserData struct {
	UserInfo UserInfo `json:"user_info"`
	Name     string   `json:"name"`     // 昵称
	Avatar   string   `json:"avatar"`   // 头像
	Sign     string   `json:"sign"`     // 签名
	Gender   string   `json:"gender"`   // 性别
	Birthday string   `json:"birthday"` // 生日
	Address  string   `json:"address"`
}
