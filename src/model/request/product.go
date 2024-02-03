package request

import "mime/multipart"

type Product struct {
	Name     string                  `form:"name" binding:"required"`
	Price    string                  `form:"price" binding:"required"`
	Describe string                  `form:"describe" binding:"required"`
	IsFree   bool                    `form:"is_free" binding:"required"`
	Tags     string                  `form:"tags" binding:"required"`
	IsLend   bool                    `form:"is_lend" binding:"required"`
	Media    []*multipart.FileHeader `form:"media" binding:"required"`
}
