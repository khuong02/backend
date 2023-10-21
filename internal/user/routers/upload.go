package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/khuong02/backend/internal/user/codeerror"
	"github.com/khuong02/backend/internal/user/dtos"
	"github.com/khuong02/backend/internal/user/payload"
	"github.com/khuong02/backend/pkg/helper"
	"net/http"
)

// Upload media upload
//
//	@Summary		Upload media upload
//	@Description	media upload
//	@Tags			Media
//	@Accept			json
//	@Produce		json
//
//	@Security		BearerAuth
//
// Upload represents parameters for the media endpoint
//
//	@Param			file	formData	file	false	"Media"
//	@Success		200		{object}	helper.ResponseSuccess
//	@Failure		400		{object}	helper.ResponseErr
//	@Failure		500		{object}	helper.ResponseErr
//	@Router			/upload [post] .
func (r *Route) Upload(c *gin.Context) {
	var (
		req  payload.MediaUpload
		resp *dtos.MediaUploadResp
		_    helper.ResponseErr
	)

	file, err := c.FormFile("file")
	if err != nil {
		r.logger.Error(err.Error())

		helper.Error(c, codeerror.ErrBadRequest(err))

		return
	}
	req.Media = file

	resp, err = r.Media.UploadMedia(c, req)
	if err != nil {
		helper.Error(c, err.(helper.Err))

		return
	}

	helper.Success(c, http.StatusCreated, "successfully", resp)
}
