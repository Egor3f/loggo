package loggo

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/egor3f/loggo/internal/filter"

	"github.com/egor3f/loggo/internal/reader"

	"github.com/egor3f/loggo/internal/color"
	"github.com/egor3f/loggo/internal/config"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type LogView struct {
	tview.Flex
	app                *LoggoApp
	chanReader         reader.Reader
	table              *tview.Table
	jsonView           *JsonView
	data               *LogData
	templateView       *TemplateView
	layout             *tview.Flex
	config             *config.Config
	keyMap             map[string]*config.Key
	navMenu            *tview.Flex
	mainMenu           *tview.Flex
	filterView         *FilterView
	linesView          *tview.TextView
	followingView      *tview.TextView
	logFullScreen      bool
	templateFullScreen bool
	inSlice            []map[string]interface{}
	finSlice           []map[string]interface{}
	filterChannel      chan *filter.Expression
	filterLock         sync.RWMutex
	globalCount        int64
	isFollowing        bool
	hideFilter         bool
	rebufferFilter     bool
	selectionEnabled   bool
	mouseSel           *tview.TextView
}

func NewLogReader(app *LoggoApp, reader reader.Reader) *LogView {
	lv := &LogView{
		Flex:          *tview.NewFlex(),
		app:           app,
		config:        app.Config(),
		chanReader:    reader,
		filterChannel: make(chan *filter.Expression, 1),
		filterLock:    sync.RWMutex{},
		hideFilter:    true,
		isFollowing:   true,
	}
	lv.makeUIComponents()
	lv.makeLayouts()
	reader.ErrorNotifier(func(err error) {
		go func() {
			time.Sleep(time.Second)
			lv.app.Draw()
		}()
		lv.app.ShowPrefabModal(fmt.Sprintf("An error occurred with the input stream: %v "+
			"\nYou can continue browsing the buffered logs or close the app.", err), 50, 20,
			func(event *tcell.EventKey) *tcell.EventKey {
				switch event.Rune() {
				case 'q', 'Q':
					lv.app.Stop()
					return nil
				case 'c', 'C':
					lv.app.DismissModal(lv.table)
					return nil
				}
				return event
			},
			tview.NewButton("[darkred::bu]Q[-::-]uit").SetSelectedFunc(func() {
				lv.app.Stop()
			}),
			tview.NewButton("[darkred::bu]C[-::-]ancel").SetSelectedFunc(func() {
				lv.app.DismissModal(lv.table)
			}))
	})

	go func() {

	}()

	lv.read()
	lv.filter()
	lv.filterChannel <- nil

	go func() {
		lv.app.ShowModal(NewSplashScreen(lv.app), 71, 16, tcell.ColorBlack, nil)
		lv.app.Draw()
		time.Sleep(2 * time.Second)
		lv.app.DismissModal(lv.table)
		lv.app.Draw()

		time.Sleep(10 * time.Millisecond)
		lv.isFollowing = true
		lv.app.SetFocus(lv.table)
	}()
	return lv
}

func (l *LogView) makeUIComponents() {
	l.templateView = NewTemplateView(l.app, false, func() {
		// Toggle full screen func
		l.templateFullScreen = !l.templateFullScreen
		l.makeLayoutsWithTemplateView()
	}, l.makeLayouts)
	l.templateView.SetBorder(true).SetTitle("Template Editor")
	l.data = &LogData{
		logView: l,
	}
	selection := func(row, column int) {
		if row > 0 && row-1 < len(l.finSlice) {
			l.jsonView = NewJsonView(l.app, false,
				func() {
					// Toggle full screen func
					l.logFullScreen = !l.logFullScreen
					l.makeLayoutsWithJsonView()
				}, l.makeLayouts)
			l.jsonView.SetBorder(true).SetTitle("Log Entry")
			var b []byte
			if _, ok := l.finSlice[row-1][config.ParseErr]; ok {
				b = []byte(fmt.Sprintf(`%v`, l.finSlice[row-1][config.TextPayload]))
			} else {
				b, _ = json.Marshal(l.finSlice[row-1])
			}
			l.jsonView.SetJson(b)
			l.makeLayoutsWithJsonView()
			l.updateBottomBarMenu()
		} else {
			l.makeLayouts()
		}
	}
	l.table = tview.NewTable().
		SetSelectable(true, false).
		SetFixed(1, 1).
		SetSeparator(tview.Borders.Vertical).
		SetContent(l.data)
	l.table.
		SetFocusFunc(func() {
			if l.isJsonViewShown() {
				l.updateBottomBarMenu()
			}
		})
	l.table.SetSelectedFunc(selection).
		SetBackgroundColor(color.ColorBackgroundField)
	l.table.SetSelectionChangedFunc(func(row, column int) {
		// stop scrolling!
		if l.isFollowing {
			l.isFollowing = false

			go func() {
				r, c := l.table.GetOffset()
				l.updateLineView()
				l.table.SetOffset(r, c)
				l.table.Select(r, c)
				go l.app.Draw()
			}()
		} else {
			r, c := l.table.GetOffset()
			l.updateLineView()
			l.table.SetOffset(r, c)
			if l.isJsonViewShown() {
				go func() {
					selection(row, column)
					l.app.Draw()
				}()
			}
		}
	})

	l.keyEvents()

	l.linesView = tview.NewTextView().SetDynamicColors(true).SetTextAlign(tview.AlignRight)
	l.followingView = tview.NewTextView().
		SetRegions(true).
		SetDynamicColors(true)
	l.followingView.SetFocusFunc(func() {
		go l.toggledFollowing()
	})
	l.followingView.SetBlurFunc(func() {
		l.followingView.Highlight("")
	})
	l.populateMenu()
	l.updateLineView()

	l.filterView = NewFilterView(l.app, func(expression *filter.Expression) {
		l.rebufferFilter = true
		l.filterChannel <- expression
		go func() {
			time.Sleep(200 * time.Millisecond)
			l.app.Draw()
		}()
	})
}

func (l *LogView) toggleFilter() {
	if l.isJsonViewShown() || l.isTemplateViewShown() {
		l.hideFilter = false
	} else {
		l.hideFilter = !l.hideFilter
	}
	l.makeLayouts()
	if !l.hideFilter {
		go l.app.SetFocus(l.filterView.expressionField)
	}
}

func (l *LogView) makeLayouts() {
	mainContent := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(l.table, 0, 2, true).
		AddItem(l.navMenu, 26, 1, false)

	l.Flex.Clear().SetDirection(tview.FlexRow)
	if !l.hideFilter {
		l.Flex.AddItem(l.filterView, 4, 2, false).
			AddItem(NewHorizontalSeparator(color.FieldStyle, LineHThick, "", 0), 1, 2, false)
	}
	l.Flex.AddItem(mainContent, 0, 2, false).
		//AddItem(l.navMenu, 1, 1, false).
		//AddItem(l.mainMenu, 1, 1, false).
		SetBackgroundColor(color.ColorBackgroundField)
	l.app.SetFocus(l.table)
}

func (l *LogView) showAbout() {
	l.app.ShowModal(NewSplashScreen(l.app), 71, 16, tcell.ColorBlack, nil)
	l.app.Draw()
	time.Sleep(4 * time.Second)
	l.app.DismissModal(l.table)
	l.app.Draw()
}

func (l *LogView) isTemplateViewShown() bool {
	return l.Flex.GetItemCount() > 0 && l.Flex.GetItem(0) == l.templateView ||
		l.Flex.GetItemCount() > 1 && l.Flex.GetItem(1) == l.templateView
}

func (l *LogView) isJsonViewShown() bool {
	return l.Flex.GetItemCount() > 0 && l.Flex.GetItem(0) == l.jsonView ||
		l.Flex.GetItemCount() > 1 && l.Flex.GetItem(1) == l.jsonView
}

func (l *LogView) toggledFollowing() {
	l.isFollowing = !l.isFollowing
	l.updateLineView()
	go l.app.Draw()
}

func (l *LogView) makeLayoutsWithJsonView() {
	l.Flex.Clear().SetDirection(tview.FlexRow)
	if !l.logFullScreen {
		l.Flex.AddItem(l.table, 0, 1, false)
	}
	l.Flex.
		AddItem(l.jsonView, 0, 2, false).
		AddItem(l.mainMenu, 1, 1, false)

	l.jsonView.textView.SetFocusFunc(func() {
		go func() {
			time.Sleep(10 * time.Millisecond)
			l.updateBottomBarMenu()
			l.app.Draw()
		}()
	})
	l.app.SetFocus(l.table)
}

func (l *LogView) makeLayoutsWithTemplateView() {
	l.isFollowing = false
	l.Flex.Clear().SetDirection(tview.FlexRow)
	if !l.templateFullScreen {
		l.Flex.AddItem(l.table, 0, 1, false)
	}
	l.templateView.config = l.config
	l.Flex.
		AddItem(l.templateView, 0, 2, false).
		AddItem(l.mainMenu, 1, 1, false)

	l.app.SetFocus(l.templateView.table)
}
