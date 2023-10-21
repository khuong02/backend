package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	UserName string    `json:"user_name"`
	Password string    `json:"password"`

	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at"`
}

func (User) GetTableName() string {
	return "users"
}
