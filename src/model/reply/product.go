package reply

type Product struct {
	ID          int64    `json:"id"`
	Name        string   `json:"name"`
	UserID      int64    `json:"user_id"`
	Price       string   `json:"price"`
	Texts       string   `json:"texts"`
	IsFree      bool     `json:"is_free"`
	IsLend      bool     `json:"is_lend"`
	Media       []string `json:"media"`
	Tags        []Tags   `json:"tags"`
	DisableDate []string `json:"disable_date"`
	IsLike      bool     `json:"is_like"`
}

type ProductInfo struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Price    string `json:"price"`
	Media    string `json:"media"`
	UserName string `json:"user_name"`
	Avatar   string `json:"avatar"`
	IsFree   bool   `json:"is_free"`
}

type Tags struct {
	ID  int64  `json:"id"`
	Tag string `json:"tag"`
}
