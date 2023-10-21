package helper

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	HTTPCode int `json:"http_code,omitempty"`
	Code     int `json:"code,omitempty"`
	Message  any `json:"message,omitempty"`
}

type ResponseSuccess struct {
	Response
	Data any `json:"data,omitempty"`
}

type ResponseErr struct {
	Response
	Details []any `json:"details"`
}

func Success(c *gin.Context, httpCode int, message any, data interface{}) {
	c.JSON(httpCode, ResponseSuccess{
		Response: Response{
			HTTPCode: httpCode,
			Message:  message,
		},
		Data: data,
	})
}

func Error(c *gin.Context, err Err) {
	c.AbortWithStatusJSON(err.HTTPCode, ResponseErr{
		Response: Response{
			HTTPCode: err.HTTPCode,
			Code:     err.ErrorCode,
			Message:  err.Error(),
		},
		Details: []any{err},
	})
}
