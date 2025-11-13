package dto

import (
	"mime/multipart"
)

type UploadFileDto struct {
	File   *multipart.FileHeader `form:"file" binding:"required"`
	Folder string                `form:"folder" binding:"required"`

	SessionID string `form:"sessionId"`
}
