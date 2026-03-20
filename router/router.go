package router

import (
	"auth-go/controllers"
	"auth-go/middlewares"
	"auth-go/utils"

	"github.com/go-chi/chi/v5"
)

type Router interface {
	Register(r chi.Router)
}

func SetupRouter(UserRouter Router) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middlewares.RequestLogger)
	router.Use(middlewares.RateLimitMiddleware)
	router.Get("/ping", controllers.PingHandler)
	router.HandleFunc("/fakestoreservice/*", utils.ProxyToService("https://fakestoreapi.com", "/fakestoreservice"))
	UserRouter.Register(router)
	return router
}
