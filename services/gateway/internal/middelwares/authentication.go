package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/MarskTM/financial_report_server/infrastructure/cache"
	"github.com/MarskTM/financial_report_server/infrastructure/encrypt"
	"github.com/MarskTM/financial_report_server/utils"
	"github.com/golang/glog"
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

		accessData, err := cache.FetchAuth(accessUuid)
		if err != nil {
			utils.UnauthorizedResponse(w, r, fmt.Errorf("Unauthorized"))
			return
		}

		if accessData.Uuid == 0 || accessData.UserID == 0 {
			glog.V(1).Info("(-) Unauthorized access data is empty")
			utils.UnauthorizedResponse(w, r, fmt.Errorf("Unauthorized"))
			return
		}

		next.ServeHTTP(w, r)
	})
}
