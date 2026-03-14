package app

import (
	config "AuthInGo/config/env"
	"AuthInGo/controllers"
	db "AuthInGo/db/repositories"
	"AuthInGo/router"
	services "AuthInGo/service"
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Addr string
}

type Application struct {
	Config Config
}

func NewConfig() Config {

	port := config.GetString("PORT", ":8080")
	return Config{
		Addr: port,
	}
}

func NewApplication(config Config) *Application {
	return &Application{
		Config: config,
	}
}

func (app *Application) Run() error {
	ur := db.NewUserRepository()
	us := services.NewUserService(ur)
	uc := controllers.NewUserController(us)
	uRouter := router.NewUserRouter(uc)

	server := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      router.SetupRouter(uRouter),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Starting server on: ", app.Config.Addr)
	return server.ListenAndServe()
}
