package loggo

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const LineHThick = '\u2501'
const LineHThin = '\u2500'

func NewHorizontalSeparator(lineStyle tcell.Style, lineRune rune, text string, textColor tcell.Color) *tview.Box {
	return tview.NewBox().
		SetDrawFunc(func(screen tcell.Screen, x int, y int, width int, height int) (int, int, int, int) {
			// Draw a horizontal line across the middle of the box.
			centerY := y + height/2
			for cx := x; cx < x+width; cx++ {

				screen.SetContent(cx, centerY, lineRune, nil, lineStyle)
			}

			// Write some text along the horizontal line.
			if len(text) > 0 {
				tview.Print(screen, text, x+1, centerY, width-2, tview.AlignCenter, textColor)
			}

			// Space for other content.
			return x + 1, centerY + 1, width - 2, height - (centerY + 1 - y)
		})
}
