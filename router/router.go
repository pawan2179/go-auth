package router

import (
	"AuthInGo/controllers"

	"github.com/go-chi/chi/v5"
)

type Router interface {
	Register(r chi.Router)
}

func SetupRouter(UserRouter Router) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/ping", controllers.PingHandler)
	UserRouter.Register(router)
	return router
}
