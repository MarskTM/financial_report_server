package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/MarskTM/financial_report_server/env"
	"github.com/MarskTM/financial_report_server/infrastructure/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"github.com/golang/glog"
	"github.com/twinj/uuid"
)

func LoadConfig(model interface{}) (data interface{}, err error) {
	switch model.(type) {
	case *env.GatewayConfig,
		*env.DocumentConfig,
		*env.AuthenConfig:

		if _, err := toml.DecodeFile("./config.toml", &model); err != nil {
			fmt.Println("Error loading config file:", err)
			return nil, err
		}

		glog.V(1).Infof("load configuration for gateway successfully: %+v", model)
		return model, nil
	default:
		return nil, fmt.Errorf("unsupported model type: %T", model)
	}
}

// ------------------------------------------------------------------------------------------------
func GetAndDecodeToken(token string, decodeAuth *jwtauth.JWTAuth) (map[string]interface{}, error) {
	if token == "" {
		return nil, nil
	}
	decodedToken, err := decodeAuth.Decode(token)
	if err != nil {
		return nil, err
	}
	claims, err := decodedToken.AsMap(context.Background())
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func CreateToken(session int32, userId int32, username string, roles []string, encodeAuth *jwtauth.JWTAuth) (*model.TokenDetail, error) {
	var err error

	// Create token details
	tokenDetail := &model.TokenDetail{}

	tokenDetail.Username = username
	tokenDetail.AccessUUID = PatternGet(uint(userId)) + uuid.NewV4().String()
	tokenDetail.RefreshUUID = PatternGet(uint(userId)) + uuid.NewV4().String()
	tokenDetail.AtExpires = time.Now().Add(time.Hour * time.Duration(env.ExtendHour)).Unix()
	tokenDetail.RtExpires = time.Now().Add(time.Hour * time.Duration(env.ExtendRefreshHour)).Unix()

	// Create Access Token
	atClaims := jwt.MapClaims{}
	atClaims["access_uuid"] = tokenDetail.AccessUUID
	atClaims["username"] = tokenDetail.Username
	atClaims["user_id"] = userId
	atClaims["role"] = roles
	atClaims["session"] = session
	atClaims["exp"] = tokenDetail.AtExpires

	_, tokenDetail.AccessToken, err = encodeAuth.Encode(atClaims)
	if err != nil {
		return nil, err
	}

	// Create Resfresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = tokenDetail.RefreshUUID
	rtClaims["username"] = tokenDetail.Username
	rtClaims["user_id"] = userId
	rtClaims["role"] = roles
	rtClaims["session"] = session
	rtClaims["exp"] = tokenDetail.RtExpires
	_, tokenDetail.RefreshToken, err = encodeAuth.Encode(rtClaims)
	if err != nil {
		return nil, err
	}

	return tokenDetail, nil
}
