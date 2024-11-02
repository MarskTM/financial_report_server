package rpc

import (
	"encoding/json"
	"net/http"

	"github.com/MarskTM/financial_report_server/env"
	"github.com/MarskTM/financial_report_server/utils"
	"github.com/go-chi/render"
	"google.golang.org/grpc"
)

type gatewayController struct {
	grpcConnected map[string]*grpc.ClientConn
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
	var payload env.BasicQueryPayload
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
	var payload env.AdvanceFilterPayload
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

func NewGatewayInterface(grpcConnected map[string]*grpc.ClientConn) GatewayController {
	return &gatewayController{
		grpcConnected: grpcConnected,
	}
}
