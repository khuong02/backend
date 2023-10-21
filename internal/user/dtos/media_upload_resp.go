package dtos

type MediaUploadResp struct {
	Link string `json:"link"`
}

func NewMediaUploadResp(link string) *MediaUploadResp {
	return &MediaUploadResp{Link: link}
}
