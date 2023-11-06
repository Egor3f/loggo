package main

import "github.com/egor3f/loggo/internal/loggo"

func main() {
	app := loggo.NewApp("")
	view := loggo.NewSplashScreen(app)
	app.Run(view)
}
