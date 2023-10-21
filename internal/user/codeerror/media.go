package codeerror

import (
	"github.com/khuong02/backend/pkg/helper"
	"net/http"
)

var ErrorMediaMessage = map[int]helper.ErrorMessage{
	ErrMaxSizeMediaUploadCode: {
		"Vượt quá số lượng upload cho phép",
		"Maximum media upload",
	},
	ErrInsertMediaToDBCode: {
		"Thêm phương tiện thất bại",
		"Insert media fail",
	},
}

func ErrMaxSizeMediaUpload() helper.Err {
	return helper.Err{
		Raw:          nil,
		HTTPCode:     http.StatusBadRequest,
		ErrorCode:    ErrMaxSizeMediaUploadCode,
		ErrorMessage: ErrorMediaMessage[ErrMaxSizeMediaUploadCode],
	}
}

func ErrInsertMediaToDB(err error) helper.Err {
	return helper.Err{
		Raw:          err,
		HTTPCode:     http.StatusInternalServerError,
		ErrorCode:    ErrInsertMediaToDBCode,
		ErrorMessage: ErrorMediaMessage[ErrInsertMediaToDBCode],
	}
}
