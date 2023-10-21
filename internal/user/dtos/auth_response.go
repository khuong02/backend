package dtos

type AuthResponse struct {
	AccessToken string `json:"access_token"`
}

func NewAuthResponse(accessToken string) *AuthResponse {
	return &AuthResponse{AccessToken: accessToken}
}
