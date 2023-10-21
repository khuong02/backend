package dtos

type CreateUser struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
