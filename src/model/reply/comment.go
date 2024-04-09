package reply

type Comment struct {
	ID       int64    `json:"id"`
	UserID   int64    `json:"user_id"`
	Username string   `json:"username"`
	Avatar   string   `json:"avatar"`
	Text     string   `json:"text"`
	Media    []string `json:"media"`
}
