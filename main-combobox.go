package main

import (
	"github.com/andlabs/ui"
)

func main() {
	err := ui.Main(func() {
		box := ui.NewVerticalBox()
		cb := ui.NewCombobox()
		cb.Append("Line 1")
		cb.Append("Line 2")
		cb.Append("Line 3")
		cb.SetSelected(0)
		box.Append(cb, false)
		button := ui.NewButton("Greet")
		box.Append(button, false)
		greeting := ui.NewLabel("")
		box.Append(greeting, false)
		window := ui.NewWindow("Hello Combobox", 350, 100, false)
		window.SetMargined(true)
		window.SetChild(box)
		cb.OnSelected(func(*ui.Combobox) {
			// Does not seem t be any way to get Values from Combobox
			//	greeting.SetText("Selected, " + *cb.Selected() + "!")
		})
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}
