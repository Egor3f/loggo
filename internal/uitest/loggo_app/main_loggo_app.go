package main

import (
	"os"

	"github.com/egor3f/loggo/internal/loggo"
	"github.com/egor3f/loggo/internal/reader"
	"github.com/egor3f/loggo/internal/uitest/helper"
)

func main() {
	inputChan := make(chan string, 1)
	rd := reader.MakeReader("", inputChan)
	oldStdIn := os.Stdin
	defer func() {
		os.Stdin = oldStdIn
	}()
	r, w, _ := os.Pipe()
	os.Stdin = r
	go func() {
		helper.JsonGenerator(w)
	}()

	_ = rd.StreamInto()
	app := loggo.NewLoggoApp(rd, "")
	app.Run()
}
