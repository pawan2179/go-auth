package main

import (
	"AuthInGo/app"
	"fmt"
)

func main() {
	fmt.Println("Starting application")

	config := app.NewConfig(":3001")
	app := app.NewApplication(config)
	app.Run()
}
