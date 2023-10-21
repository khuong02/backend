package payload

import (
	"github.com/khuong02/backend/internal/user/codeerror"
	"github.com/khuong02/backend/internal/user/dtos"
	"github.com/khuong02/backend/utils"
	"strings"
)

type Register struct {
	UserName string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func (req Register) Validate() error {
	if strings.TrimSpace(req.UserName) == "" || strings.TrimSpace(req.Password) == "" {
		return codeerror.ErrInvalidParams("username or password")
	}

	return nil
}

func (req Register) ToDTO() dtos.CreateUser {
	pass, _ := utils.Password(req.Password)

	return dtos.CreateUser{
		UserName: req.UserName,
		Password: pass,
	}
}
