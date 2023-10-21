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
	ErrNotExistUserCode: {
		"tài khoản không tồn tại",
		"doesn't exist account",
	},
	ErrLoginFailedCode: {
		"đăng nhập thất bại",
		"login fail",
	},
	ErrWrongPasswordCode: {
		"sai mật khẩu",
		"wrong password",
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

func ErrNotExistUser(err error) helper.Err {
	return helper.Err{
		Raw:          err,
		HTTPCode:     http.StatusNotFound,
		ErrorCode:    ErrNotExistUserCode,
		ErrorMessage: errorRegisterMessage[ErrNotExistUserCode],
	}
}

func ErrLoginFailed(err error) helper.Err {
	return helper.Err{
		Raw:          err,
		HTTPCode:     http.StatusInternalServerError,
		ErrorCode:    ErrLoginFailedCode,
		ErrorMessage: errorRegisterMessage[ErrLoginFailedCode],
	}
}

func ErrWrongPassword(err error) helper.Err {
	return helper.Err{
		Raw:          err,
		HTTPCode:     http.StatusBadRequest,
		ErrorCode:    ErrWrongPasswordCode,
		ErrorMessage: errorRegisterMessage[ErrWrongPasswordCode],
	}
}
