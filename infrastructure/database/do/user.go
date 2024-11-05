package do

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID       int32  `json:"id"`
	RoleID   int32  `json:"role_id"`
	Username string `json:"Username"`
	Password string `json:"Password"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

