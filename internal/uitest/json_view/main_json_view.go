package main

import (
	"os"

	"github.com/aurc/loggo/internal/loggo"
)

func main() {
	app := loggo.NewApp("")
	view := loggo.NewJsonView(app, true, nil, nil)

	b, err := os.ReadFile("internal/testdata/test1.json")
	if err != nil {
		panic(err)
	}
	view.SetJson(b)

	app.Run(view)
}
