package rpc

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/MarskTM/financial_report_server/env"
	"github.com/MarskTM/financial_report_server/infrastructure/model"
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

}

func (c *gatewayController) Logout(w http.ResponseWriter, r *http.Request) {

}

// --------------------------------------------------------------------------------------------------------------------------------
func (c *gatewayController) BasicQuery(w http.ResponseWriter, r *http.Request) {
	var payload BasicQueryPayload
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
	var payload AdvanceFilterPayload
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
type TokenDetail struct {
	Username     string
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	AtExpires    int64
	RtExpires    int64
}

func SaveHttpCookie(fullDomain string, tokenDetail *TokenDetail, w http.ResponseWriter) error {
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
