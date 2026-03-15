package main

import (
	app "AuthInGo/app"
	config "AuthInGo/config/env"
	"fmt"
)

func main() {
	config.Load()
	fmt.Println("Starting application")

	config := app.NewConfig()
	app := app.NewApplication(config)
	app.Run()
}
