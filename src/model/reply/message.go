package reply

// MsgInfo 完整的消息详情
type MsgInfo struct {
	ID         int64  `json:"id"`          // 消息ID
	MsgType    string `json:"msg_type"`    // 消息类型 [text,file]
	MsgContent string `json:"msg_content"` // 消息内容 文件则为url，文本则为文本内容，由拓展信息进行补充
	FileID     int64  `json:"file_id"`     // 文件ID 当消息类型为file时>0
	FromID     int64  `json:"form_id"`     // 发送者ID
	ToID       int64  `json:"to_id"`       // 收取者ID
}
