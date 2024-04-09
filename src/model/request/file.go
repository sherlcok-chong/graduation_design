package request

import "mime/multipart"

type UpdateUserAvatar struct {
	Avatar *multipart.FileHeader `form:"avatar" binding:"required"`
}

type UploadFile struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}
