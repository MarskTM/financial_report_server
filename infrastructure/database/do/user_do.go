package do

import (
	"time"
)

type User struct {
	ID       int32  `json:"id"`
	Username string `json:"Username"`
	Password string `json:"Password"`

	UserRoles []UserRole `json:"user_roles" gorm:"foreignKey:UserID"`

	CreatedAt time.Time `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updatedAt" swaggerignore:"true"`
	DeletedAt time.Time `json:"-" swaggerignore:"true"`
}

type UserResponse struct {
	ID       int32    `json:"id"`
	Roles     []string `json:"role"`
	Username string   `json:"username"`
	FullName string   `json:"fullname"`
	Profile  *Profile `json:"profile"`
}
