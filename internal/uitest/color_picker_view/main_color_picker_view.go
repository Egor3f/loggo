package main

import "github.com/egor3f/loggo/internal/loggo"

func main() {
	app := loggo.NewApp("")
	view := loggo.NewColorPickerView(app, "Select Color",
		func(c string) {
		}, func() {
			app.Stop()
		}, func() {
			app.Stop()
		})
	app.Run(view)
}
