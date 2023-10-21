package payload

import "mime/multipart"

type MediaUpload struct {
	Media *multipart.FileHeader `form:"media"`
}
