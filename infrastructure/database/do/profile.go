package do

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	ID         int32  `json:"id"`
	UserID     int32  `json:"user_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Alias      string `json:"alias"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Birthdate  string `json:"birthdate"`
	OtherLinks string `json:"other_links"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
