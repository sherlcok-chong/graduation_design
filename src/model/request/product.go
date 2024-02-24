package request

import "mime/multipart"

type Product struct {
	Name     string                  `form:"name" binding:"required"`
	Price    string                  `form:"price" binding:"required"`
	Describe string                  `form:"describe" binding:"required"`
	IsFree   bool                    `form:"is_free" binding:"required"`
	Tags     []int64                 `form:"tags" binding:"required"`
	IsLend   int                     `form:"is_lend" binding:"required"`
	Media    []*multipart.FileHeader `form:"media" binding:"required"`
}

type ProductInfo struct {
	Offset int32 `form:"offset" json:"offset"`
}

type ProductDetails struct {
	ID int64 `form:"id" json:"id" binding:"required"`
}

type ProductID struct {
	ID int64 `form:"id" binding:"required"`
}

type UpdateProduct struct {
	ID       int64                   `form:"id" json:"id" binding:"required"`
	Name     string                  `form:"name" binding:"required" json:"name,omitempty"`
	Price    string                  `form:"price" binding:"required" json:"price,omitempty"`
	Describe string                  `form:"describe" binding:"required" json:"describe,omitempty"`
	IsFree   bool                    `form:"is_free" binding:"required" json:"isFree,omitempty"`
	Tags     []int64                 `form:"tags" binding:"required" json:"tags,omitempty"`
	Media    []*multipart.FileHeader `form:"media" binding:"required" json:"media,omitempty"`
}
type AddComment struct {
	ProductID int64                   `form:"product_id" binding:"required"`
	Comment   string                  `form:"comment" binding:"required"`
	Media     []*multipart.FileHeader `form:"media" binding:"required"`
}

type DeleteComment struct {
	ID int64 `json:"id" binding:"required"`
}

type GetProductComment struct {
	ID int64 `form:"id" binding:"required"`
}
