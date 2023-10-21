package codeerror

import (
	"github.com/khuong02/backend/pkg/helper"
	"net/http"
)

var errorRegisterMessage = map[int]helper.ErrorMessage{
	ErrRegisterFailCode: {
		"đăng ký thất bại",
		"register fail",
	},
}

func ErrRegisterFail(err error) helper.Err {
	return helper.Err{
		Raw:          err,
		HTTPCode:     http.StatusInternalServerError,
		ErrorCode:    ErrRegisterFailCode,
		ErrorMessage: errorRegisterMessage[ErrInvalidParamsCode],
	}
}
