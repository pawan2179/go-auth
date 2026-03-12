package main

import (
	"AuthInGo/app"
	"fmt"
)

func main() {
	fmt.Println("Starting application")

	config := app.Config{
		Addr: ":3001",
	}

	app := app.Application{
		Config: config,
	}

	app.Run()
}
