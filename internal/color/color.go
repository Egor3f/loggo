package color

import "github.com/gdamore/tcell/v2"

const (
	ColorBackgroundField    = tcell.ColorBlack
	ColorForegroundField    = tcell.ColorWhite
	ColorSelectedBackground = tcell.Color69
	ColorSelectedForeground = tcell.ColorWhite
)

var (
	FieldStyle = tcell.StyleDefault.
			Background(ColorBackgroundField).
			Foreground(ColorForegroundField)
	PlaceholderStyle = tcell.StyleDefault.
				Background(ColorBackgroundField).
				Foreground(tcell.ColorDarkGray)
	SelectStyle = tcell.StyleDefault.
			Background(ColorSelectedBackground).
			Foreground(ColorSelectedForeground)
)

const (
	ClField   = "[#ffaf00::b]"
	ClWhite   = "[#ffffff::-]"
	ClNumeric = "[#00afff]"
	ClString  = "[#6A9F59]"
)
