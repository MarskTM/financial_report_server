package internal

import (
	"net/http"

	"github.com/MarskTM/financial_report_server/infrastructure/model"
	internalMiddel "github.com/MarskTM/financial_report_server/services/gateway/internal/middelwares"
	"github.com/MarskTM/financial_report_server/services/gateway/internal/rpc"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
)

func Router(gatewayModel model.GatewayModel) http.Handler {
	r := chi.NewRouter()

	// Sử dụng middleware cho router
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Use this to allow specific origin hosts
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},

		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	controller := rpc.NewGatewayInterface(gatewayModel)

	r.Route("/api/v1", func(router chi.Router) {
		// Ping the API server
		router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("pong"))
		})

		// Public routes
		router.Post("/login", controller.Login)
		router.Post("/logout", func(w http.ResponseWriter, r *http.Request) {})
		router.Post("/refresh", func(w http.ResponseWriter, r *http.Request) {})
		router.Post("/users/register", controller.Register)
		router.Post("/users/forgot-password", func(w http.ResponseWriter, r *http.Request) {})

		// Protected routes with JWT token authentication
		router.Group(func(protectRouter chi.Router) {
			protectRouter.Use(jwtauth.Authenticator)
			protectRouter.Use(internalMiddel.Authenticator)

			protectRouter.Route("/users", func(userRouter chi.Router) {
				userRouter.Put("/reset-password", func(w http.ResponseWriter, r *http.Request) {})
				userRouter.Put("/change-password", func(w http.ResponseWriter, r *http.Request) {})
			})

			protectRouter.Route("/basic-query", func(accessRouter chi.Router) {
				accessRouter.Post("/", controller.BasicQuery)
				accessRouter.Delete("/", controller.BasicQuery)
			})

			protectRouter.Route("/advance-filter", func(accessRouter chi.Router) {
				accessRouter.Post("/", controller.AdvancedFilter)
			})

			protectRouter.Route("/document", func(accessRouter chi.Router) {
				accessRouter.Post("/", func(w http.ResponseWriter, r *http.Request) {})
			})
		})

		router.Group(func(protectedRoute chi.Router) {
			fs := http.StripPrefix("/api/v1/public", http.FileServer(http.Dir("../../../cdn/public")))
			router.Get("/pnk_intern_storage/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fs.ServeHTTP(w, r)
			}))
		})
	})

	return r
}
