package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/MarskTM/financial_report_server/infrastructure/encrypt"
	redisV2 "github.com/MarskTM/financial_report_server/infrastructure/resdis"
	"github.com/MarskTM/financial_report_server/utils"
)

func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if authorization == "" {
			utils.BadRequestResponse(w, r, fmt.Errorf("Authorization header not found"))
			return
		}
		authorizationBearer := strings.Split(authorization, " ")[1]
		accessToken := strings.Split(authorizationBearer, ";")[0]
		accessClaims, errDecodeToken := utils.GetAndDecodeToken(accessToken, encrypt.GetDecodeAuth())
		if errDecodeToken != nil {
			utils.UnauthorizedResponse(w, r, errDecodeToken)
			return
		}

		// Check authentication
		accessUuid, ok := accessClaims["access_uuid"].(string)
		if !ok {
			utils.UnauthorizedResponse(w, r, fmt.Errorf("can't parse access uuid from token"))
			return
		}

		if index, isExist := redisV2.FetchAuth(accessUuid); isExist != nil || index == 0 {
			utils.UnauthorizedResponse(w, r, fmt.Errorf("Unauthorized"))
			return
		}
		next.ServeHTTP(w, r)
	})
}


