package chat

import "time"

type MsgSend struct {
	ID        int64     `json:"id"`
	FClientID int64     `json:"f_client_id"`
	FUserID   int64     `json:"f_user_id"`
	TUserID   int64     `json:"t_user_id"`
	MsgType   int64     `json:"msg_type"`
	Text      string    `json:"text"`
	IsRead    bool      `json:"is_read"`
	CreateAt  time.Time `json:"create_at"`
}

type NotReadMsg struct {
	UserID   int64     `json:"user_id,omitempty"`
	UserName string    `json:"user_name,omitempty"`
	Avatar   string    `json:"avatar,omitempty"`
	Msg      []MsgSend `json:"msg,omitempty"`
}
