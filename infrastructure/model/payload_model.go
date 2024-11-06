package model

import "github.com/lib/pq"

type AdvanceFilterPayload struct {
	ModelType         string         `json:"modelType"`
	IgnoreAssociation bool           `json:"ignoreAssociation"`
	Page              int            `json:"page"`
	PageSize          int            `json:"pageSize"`
	IsPaginateDB      bool           `json:"isPaginateDB"`
	QuerySerch        string         `json:"querySearch"`
	SelectColumn      pq.StringArray `json:"selectColumn"`
}

type BasicQueryPayload struct {
	Contructor string      `json:"contructor"`
	Data       interface{} `json:"data"`
}

type ListModelId struct {
	ID        []uint `gorm:"column:id"`
	ModelType string `json:"modelType"`
}

// AccessDetail access detail only from token
type AccessDetail struct {
	AccessUUID string
	UserID     int
}

// Payload for authentication
type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterPayload struct {
	FullName string `json:"fullName"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Code     string `json:"code"`
	Phone    string `json:"phone"`
	Birthday string `json:"birthday"`
}

type ChangePasswordPayload struct {
	Username    string `json:"username"`
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type ForgotPasswordPayload struct {
	FogortCode  string `json:"forgotCode"`
	NewPassword string `json:"newPassword"`
}

type EmailForgotPayload struct {
	Email string `json:"email"`
}


// TokenDetail details for token authentication
type TokenDetail struct {
	Username     string
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	AtExpires    int64
	RtExpires    int64
}