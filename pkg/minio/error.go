package minio

import (
	"github.com/khuong02/backend/pkg/helper"
	"net/http"
)

const (
	// upload media
	ErrFileUploadCode      = 10301
	ErrFileUploadMinioCode = 10302
)

var ErrorMediaMessage = map[int]helper.ErrorMessage{
	ErrFileUploadCode: {
		"File bị lỗi",
		"File upload error",
	},
	ErrFileUploadMinioCode: {
		"Upload file lên minio thất bại",
		"Upload file to minio failed",
	},
}

func ErrFileUpload(err error) helper.Err {
	return helper.Err{
		Raw:          err,
		HTTPCode:     http.StatusInternalServerError,
		ErrorCode:    ErrFileUploadCode,
		ErrorMessage: ErrorMediaMessage[ErrFileUploadCode],
	}
}

func ErrFileUploadMinio(err error) helper.Err {
	return helper.Err{
		Raw:          err,
		HTTPCode:     http.StatusInternalServerError,
		ErrorCode:    ErrFileUploadMinioCode,
		ErrorMessage: ErrorMediaMessage[ErrFileUploadMinioCode],
	}
}
