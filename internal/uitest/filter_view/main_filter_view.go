package main

import (
	"github.com/egor3f/loggo/internal/color"
	"github.com/egor3f/loggo/internal/loggo"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := loggo.NewApp("internal/config-sample/gcp.yaml")
	main := tview.NewFlex().SetDirection(tview.FlexRow)
	main.AddItem(loggo.NewFilterView(app, nil), 4, 1, true).
		AddItem(loggo.NewHorizontalSeparator(color.FieldStyle, loggo.LineHThick, "test", tcell.ColorYellow), 1, 2, false)
	app.Run(main)
}
