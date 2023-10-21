package payload

import (
	"github.com/khuong02/backend/internal/user/codeerror"
	"strings"
)

type Login struct {
	UserName string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func (req Login) Validate() error {
	if strings.TrimSpace(req.UserName) == "" || strings.TrimSpace(req.Password) == "" {
		return codeerror.ErrInvalidParams("username or password")
	}

	return nil
}
