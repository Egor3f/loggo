package loggo

import (
	"github.com/aurc/loggo/internal/config"
	"github.com/aurc/loggo/internal/reader"
	"github.com/aurc/loggo/internal/util"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var BuildVersion string

type LoggoApp struct {
	appScaffold
	chanReader reader.Reader
	logView    *LogView
}

type Loggo interface {
	Draw()
	SetInputCapture(cap func(event *tcell.EventKey) *tcell.EventKey)
	Stop()
	SetFocus(primitive tview.Primitive)
	ShowPopMessage(text string, waitSecs int64, resetFocusTo tview.Primitive)
	ShowPrefabModal(text string, width, height int, capture inputCapture, buttons ...*tview.Button)
	ShowModal(p tview.Primitive, width, height int, bgColor tcell.Color, capture inputCapture)
	DismissModal(resetFocusTo tview.Primitive)
	Config() *config.Config
	StackView(p tview.Primitive)
	PopView()
}

func NewLoggoApp(reader reader.Reader, configFile string) *LoggoApp {
	app := NewApp(configFile)
	lapp := &LoggoApp{
		appScaffold: *app,
		chanReader:  reader,
	}

	lapp.logView = NewLogReader(lapp, reader)

	lapp.pages = tview.NewPages().
		AddPage("background", lapp.logView, true, true)

	return lapp
}

func (a *LoggoApp) Run() {
	if err := a.app.
		SetRoot(a.pages, true).
		EnableMouse(true).
		Run(); err != nil {
		util.Log().Error(err)
		panic(err)
	}
}
