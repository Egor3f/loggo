package loggo

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (l *LogView) keyEvents() {
	l.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if l.app.inputCapture != nil {
			return l.app.inputCapture(event)
		}
		switch event.Key() {
		case tcell.KeyCtrlN:
			l.toggleSelectionMouse()
			return nil
		case tcell.KeyCtrlA:
			go func() {
				l.showAbout()
			}()
			return nil
		case tcell.KeyCtrlT:
			l.makeLayoutsWithTemplateView()
			return nil
		case tcell.KeyCtrlSpace:
			l.toggledFollowing()
			return nil
		case tcell.KeyTAB:
			if l.isJsonViewShown() {
				if l.jsonView.textView.HasFocus() {
					l.app.SetFocus(l.table)
					go func() {
						time.Sleep(time.Millisecond)
						l.updateBottomBarMenu()
					}()
				} else {
					l.app.SetFocus(l.jsonView.textView)
					go func() {
						time.Sleep(time.Millisecond)
						l.updateBottomBarMenu()
					}()
				}
				return nil
			}
			return event
		}
		prim := l.app.app.GetFocus()
		if _, ok := prim.(*tview.InputField); ok {
			return event
		}
		switch event.Rune() {
		case ':':
			l.toggleFilter()
			return nil
		}
		if prim == l.table && l.isJsonViewShown() {
			switch event.Rune() {
			case 'f', '`', 's', 'r', 'g', 'G', 'w', 'x':
				return l.jsonView.textView.GetInputCapture()(event)
			}
		}

		return event
	})
}
