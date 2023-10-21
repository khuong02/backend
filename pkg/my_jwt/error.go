package my_jwt

import (
	"github.com/khuong02/backend/pkg/helper"
	"net/http"
)

const (
	GenerateJWTFailedCode     = 10001
	VerifyJWTUnauthorizedCode = 10002
	VerifyJWTFailedCode       = 10003
	JWTExpiredCode            = 10004
)

var errorMessage = map[int]helper.ErrorMessage{
	GenerateJWTFailedCode: {
		"gen jwt thất bại",
		"generate jwt failed",
	},
	VerifyJWTUnauthorizedCode: {
		"bạn không có quyền",
		"Unauthorized",
	},
	VerifyJWTFailedCode: {
		"xác minh thất bại",
		"verify fail",
	},
	JWTExpiredCode: {
		"hết phiên đăng nhập",
		"expire time",
	},
}

func GenerateJWTFailed(err error) helper.Err {
	return helper.Err{
		Raw:          err,
		HTTPCode:     http.StatusInternalServerError,
		ErrorCode:    GenerateJWTFailedCode,
		ErrorMessage: errorMessage[GenerateJWTFailedCode],
	}
}

func VerifyJWTUnauthorized(err error) helper.Err {
	return helper.Err{
		Raw:          err,
		HTTPCode:     http.StatusUnauthorized,
		ErrorCode:    VerifyJWTUnauthorizedCode,
		ErrorMessage: errorMessage[VerifyJWTUnauthorizedCode],
	}
}

func VerifyJWTFailed(err error) helper.Err {
	return helper.Err{
		Raw:          err,
		HTTPCode:     http.StatusInternalServerError,
		ErrorCode:    VerifyJWTFailedCode,
		ErrorMessage: errorMessage[VerifyJWTFailedCode],
	}
}

func JWTExpired(err error) helper.Err {
	return helper.Err{
		Raw:          err,
		HTTPCode:     http.StatusForbidden,
		ErrorCode:    JWTExpiredCode,
		ErrorMessage: errorMessage[JWTExpiredCode],
	}
}
