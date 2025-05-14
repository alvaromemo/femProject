package main

import "github.com/alvaromemo/femProject/internal/app"

func main() {
	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	app.Logger.Println("we are running out app")
}
