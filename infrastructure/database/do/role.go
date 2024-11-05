package do

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID   int32  `json:"id"`
	Type string `json:"type"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
