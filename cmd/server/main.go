package main

import "fmt"

type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting up Our App")
	return nil
}

func main() {
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error Starting up")
		fmt.Println(err)
	}

}
