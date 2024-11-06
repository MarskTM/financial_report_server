package rpc

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/MarskTM/financial_report_server/env"
	"github.com/MarskTM/financial_report_server/infrastructure/model"
	"github.com/MarskTM/financial_report_server/infrastructure/proto/pb"
	"github.com/MarskTM/financial_report_server/utils"
	"github.com/go-chi/render"
)

type gatewayController struct {
	GateModel model.GatewayModel
}

type GatewayController interface {
	// access api
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)

	// CRUD...
	BasicQuery(w http.ResponseWriter, r *http.Request)
	AdvancedFilter(w http.ResponseWriter, r *http.Request)
}

func (c *gatewayController) Login(w http.ResponseWriter, r *http.Request) {
	var response utils.Response

	// TODO: Create new JWT token for user login the first time.
	var payload model.LoginPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := c.GateModel.BizClient.Authenticate(ctx, &pb.Credentials{
		Username: payload.Username,
		Password: payload.Password,
	})
	if err != nil {
		utils.InternalServerErrorResponse(w, r, err)
		return
	}

	tokenDetail, err := utils.CreateToken(res.Session, res.UserId, res.Usernames, res.Roles, c.GateModel.EncodeAuth)
	if err != nil {
		utils.InternalServerErrorResponse(w, r, err)
		return
	}

	fullDomain := r.Header.Get("Origin")
	errCookie := SaveHttpCookie(fullDomain, tokenDetail, w)
	if errCookie != nil {
		utils.InternalServerErrorResponse(w, r, err)
		return
	}

	response = utils.Response{
		Data: &model.LoginResponse{
			ID:           uint(res.UserId),
			Role:         res.Roles,
			Username:     res.Usernames,
			AccessToken:  tokenDetail.AccessToken,
			RefreshToken: tokenDetail.RefreshToken,
		},
		Success: true,
		Message: "Authenticated",
	}
	render.JSON(w, r, response)
}

func (c *gatewayController) Logout(w http.ResponseWriter, r *http.Request) {

}

func (c *gatewayController) Refresh(w http.ResponseWriter, r *http.Request) {
	// validate request body
	authorization := r.Header.Get("Authorization")
	if authorization != "" {
	}

	authorizationBearer := strings.Split(authorization, " ")[1]
	accessToken := strings.Split(authorizationBearer, ";")[0]
	accessClaims, errDecodeToken := utils.GetAndDecodeToken(accessToken, c.GateModel.DecodeAuth)
	if errDecodeToken != nil {
		utils.UnauthorizedResponse(w, r, errDecodeToken)
		return
	}

	refreshToken := strings.Split(authorizationBearer, ";")[1]
	refreshClaims, errDecodeToken := utils.GetAndDecodeToken(refreshToken, c.GateModel.DecodeAuth)
	if errDecodeToken != nil {
		utils.UnauthorizedResponse(w, r, errDecodeToken)
		return
	}

	accessUuid := accessClaims["access_uuid"].(string)
	expiresAt := accessClaims["expires_at"].(time.Time)

	refreshUuid := refreshClaims["refresh_uuid"].(string)
	session := refreshClaims["session"].(int32)
	userId := refreshClaims["user_id"].(int32)
	username := refreshClaims["username"].(string)
	roles := refreshClaims["role"].([]string)

	// Handle the SSO authentication
	if accessUuid != "" && userId != 0 && len(roles) > 0 && time.Now().Before(expiresAt) {

		response := utils.Response{
			Data:    nil,
			Success: true,
			Message: "Loging successfully!",
		}
		render.JSON(w, r, response)
		return
	} else if refreshUuid != "" && userId != 0 && len(roles) > 0 && time.Now().Before(expiresAt) {

		newTokenDetails, err := utils.CreateToken(session, userId, username, roles, c.GateModel.EncodeAuth)
		if err != nil {
			utils.InternalServerErrorResponse(w, r, err)
			return
		}

		fullDomain := r.Header.Get("Origin")
		errCookie := SaveHttpCookie(fullDomain, newTokenDetails, w)
		if errCookie != nil {
			utils.InternalServerErrorResponse(w, r, err)
			return
		}

		response := utils.Response{
			Data: &model.LoginResponse{
				ID:           uint(userId),
				Role:         roles,
				Username:     username,
				AccessToken:  newTokenDetails.AccessToken,
				RefreshToken: newTokenDetails.RefreshToken,
			},
			Success: true,
			Message: "Refresh successfully!",
		}
		render.JSON(w, r, response)
		return
	}
}

// --------------------------------------------------------------------------------------------------------------------------------
func (c *gatewayController) BasicQuery(w http.ResponseWriter, r *http.Request) {
	var payload model.BasicQueryPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	// >>>> Handle rpc here!!!!!!!!!!!!!!!!!
	// temp, err := c.BasicQueryService.Upsert(payload)
	// if err != nil {
	// 	utils.InternalServerErrorResponse(w, r, err)
	// 	return
	// }

	res := utils.Response{
		Data:    "",
		Success: true,
		Message: "Upsert success",
	}
	render.JSON(w, r, res)
	return
	// Implementation for basic query functionality
}

func (c *gatewayController) AdvancedFilter(w http.ResponseWriter, r *http.Request) {
	var res utils.Response
	var payload model.AdvanceFilterPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	// >>>> Handle rpc here!!!!!!!!!!!!!!!!!
	// temp, err := c.AdvanceFilterService.Filter(payload)
	// if err != nil {
	// 	utils.InternalServerErrorResponse(w, r, err)
	// 	return
	// }

	res = utils.Response{
		Data:    "",
		Success: true,
		Message: "Get " + payload.ModelType + " success",
	}
	render.JSON(w, r, res)
	return
}

func NewGatewayInterface(model model.GatewayModel) GatewayController {
	return &gatewayController{
		GateModel: model,
	}
}

// -----------------------------------------Utils func-------------------------------------------------------
// TokenDetail details for token authentication
func SaveHttpCookie(fullDomain string, tokenDetail *model.TokenDetail, w http.ResponseWriter) error {
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
		Expires:  time.Now().Add(time.Hour * time.Duration(env.ExtendHour)),
	}

	cookie_refresh := http.Cookie{
		Name:     "RefreshToken",
		Domain:   domain.Hostname(),
		Path:     "/",
		Value:    tokenDetail.RefreshToken,
		HttpOnly: false,
		Secure:   false,
		Expires:  time.Now().Add(time.Hour * time.Duration(env.ExtendRefreshHour)),
	}

	http.SetCookie(w, &cookie_access)
	http.SetCookie(w, &cookie_refresh)
	return nil
}
