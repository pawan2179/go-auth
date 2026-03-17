package router

import (
	"auth-go/controllers"
	"auth-go/middlewares"

	"github.com/go-chi/chi/v5"
)

type Router interface {
	Register(r chi.Router)
}

func SetupRouter(UserRouter Router) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middlewares.RequestLogger)
	router.Get("/ping", controllers.PingHandler)
	UserRouter.Register(router)
	return router
}
