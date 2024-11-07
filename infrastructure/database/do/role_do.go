package do

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID          int32   `json:"id" gorm:"primaryKey"`
	Code        string  `json:"code" gorm:"unique"`
	Type        string  `json:"type"`
	Description *string `json:"description"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
