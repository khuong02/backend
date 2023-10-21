package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/khuong02/backend/internal/user/codeerror"
	"github.com/khuong02/backend/internal/user/dtos"
	"github.com/khuong02/backend/internal/user/payload"
	"github.com/khuong02/backend/pkg/helper"
	"net/http"
)

// Register register
//
//	@Summary		Register register
//	@Description	register
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//
// Register represents parameters for the register endpoint
//
//	@Param			todo	body		payload.Register	true	"register screen"
//	@Success		201		{object}	dtos.AuthResponse
//	@Failure		400		{object}	helper.ResponseErr
//	@Failure		500		{object}	helper.ResponseErr
//	@Router			/user/register [post] .
func (r *Route) Register(c *gin.Context) {
	var (
		ctx  = c.Request.Context()
		req  payload.Register
		resp *dtos.AuthResponse
		_    helper.ResponseErr
	)

	if err := c.ShouldBind(&req); err != nil {
		r.logger.Error(err.Error())

		helper.Error(c, codeerror.ErrBadRequest(err))

		return
	}

	resp, err := r.Auth.Register(ctx, req)
	if err != nil {
		helper.Error(c, err.(helper.Err))

		return
	}

	helper.Success(c, http.StatusCreated, "successfully", resp)
}
