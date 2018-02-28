package main

import (
	"fmt"
	"github.com/andlabs/ui"
)

func main() {
	err := ui.Main(func() {
		box := ui.NewVerticalBox()
		cb := ui.NewCombobox()
		lineItem := [3]string{"Line 1", "Line 2", "Line 3"}
		for i := 0; i < len(lineItem); i++ {
			cb.Append(lineItem[i])
		}
		box.Append(cb, false)
		button := ui.NewButton("Greet")
		box.Append(button, false)
		selected := ui.NewLabel("")
		box.Append(selected, false)
		horizRule := ui.NewHorizontalSeparator()
		box.Append(horizRule, false)
		greeting := ui.NewLabel("")
		box.Append(greeting, false)
		window := ui.NewWindow("Hello Combobox", 350, 100, false)
		window.SetMargined(true)
		window.SetChild(box)
		cb.OnSelected(func(*ui.Combobox) {
			showSelection(selected, cb.Selected())
		})
		button.OnClicked(func(*ui.Button) {
			greeting.SetText("Hello, " + lineItem[cb.Selected()] + "!")
		})
		cb.SetSelected(0)
		showSelection(selected, 0)
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

func showSelection(selected *ui.Label, i int) {
	selected.SetText(fmt.Sprintf("Selected [Zero-based]: %v", i))
}
