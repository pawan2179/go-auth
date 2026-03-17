package main

import (
	app "auth-go/app"
	config "auth-go/config/env"
	"fmt"
)

func main() {
	config.Load()
	fmt.Println("Starting application")

	config := app.NewConfig()
	app := app.NewApplication(config)
	app.Run()
}
