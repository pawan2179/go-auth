package app

import (
	dbConfig "auth-go/config/db"
	config "auth-go/config/env"
	"auth-go/controllers"
	repo "auth-go/db/repositories"
	"auth-go/router"
	services "auth-go/service"
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

	db, err := dbConfig.SetupDB()

	if err != nil {
		println("Error setting DB: ", err)
		return err
	}

	ur := repo.NewUserRepository(db)
	us := services.NewUserService(ur)
	uc := controllers.NewUserController(us)
	uRouter := router.NewUserRouter(uc)

	rpr := repo.NewRolePermissionRepository(db)

	rr := repo.NewRoleRepository(db)
	rs := services.NewRoleService(rr, rpr)
	rc := controllers.NewRoleController(rs)
	rRouter := router.NewRoleRouter(rc)

	server := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      router.SetupRouter(uRouter, rRouter),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Starting server on: ", app.Config.Addr)
	return server.ListenAndServe()
}
