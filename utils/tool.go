package utils

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/MarskTM/financial_report_server/env"
	"github.com/go-chi/jwtauth"
	"github.com/golang/glog"
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

func SaveHttpCookie(fullDomain string, tokenDetail *env.TokenDetail, w http.ResponseWriter) error {
	domain, err := url.Parse(fullDomain)
	if err != nil {
		return err
	}

	cookie_access := http.Cookie{
		Name:     "AccessToken",
		Domain:   domain.Hostname(),
		Path:     "/",
		Value:    tokenDetail.AccessToken,
		HttpOnly: false,
		Secure:   false,
		Expires:  time.Now().Add(time.Hour * time.Duration(env.AccessTokenTime)),
	}

	cookie_refresh := http.Cookie{
		Name:     "RefreshToken",
		Domain:   domain.Hostname(),
		Path:     "/",
		Value:    tokenDetail.RefreshToken,
		HttpOnly: false,
		Secure:   false,
		Expires:  time.Now().Add(time.Hour * time.Duration(env.RefreshTokenTime)),
	}

	http.SetCookie(w, &cookie_access)
	http.SetCookie(w, &cookie_refresh)
	return nil
}

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
