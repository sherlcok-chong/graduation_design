package reply

type UpdateUserAvatar struct {
	UserID int64  `json:"user_id"`
	Url    string `json:"url"`
}
