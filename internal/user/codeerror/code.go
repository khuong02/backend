package codeerror

import (
	"fmt"
	"github.com/khuong02/backend/pkg/helper"
	"net/http"
)

const (
	ErrInvalidParamsCode = 20001
	ErrBadRequestCode    = 20002

	// auth
	ErrRegisterFailCode  = 20101
	ErrNotExistUserCode  = 20102
	ErrLoginFailedCode   = 20103
	ErrWrongPasswordCode = 20104

	// media
	ErrMaxSizeMediaUploadCode = 20201
	ErrInsertMediaToDBCode    = 20202
)

var errorMessage = map[int]helper.ErrorMessage{
	ErrInvalidParamsCode: {
		"Thiếu tham số",
		"Invalid Param",
	},
	ErrBadRequestCode: {
		"Yêu cầu không hợp lệ",
		"Bad request",
	},
}

func ErrInvalidParams(field string) helper.Err {
	return helper.Err{
		Raw:       nil,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: ErrInvalidParamsCode,
		ErrorMessage: helper.ErrorMessage{
			VI: fmt.Sprintf("%v %v", errorMessage[ErrInvalidParamsCode].VI, field),
			EN: fmt.Sprintf("%v %v", errorMessage[ErrInvalidParamsCode].EN, field),
		},
	}
}

func ErrBadRequest(err error) helper.Err {
	return helper.Err{
		Raw:          err,
		HTTPCode:     http.StatusBadRequest,
		ErrorCode:    ErrBadRequestCode,
		ErrorMessage: errorMessage[ErrBadRequestCode],
	}
}
