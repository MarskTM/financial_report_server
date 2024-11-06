package model

import "time"

type AccessCache struct {
	Uuid      int32     `json:"uuid"`
	UserID    int32     `json:"user_id"`
	Role      []string  `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
