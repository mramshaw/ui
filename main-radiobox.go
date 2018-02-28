package main

import (
	"github.com/andlabs/ui"
)

func main() {
	err := ui.Main(func() {
		box := ui.NewVerticalBox()
		rb := ui.NewRadioButtons()
		rb.Append("Line 1")
		rb.Append("Line 2")
		rb.Append("Line 3")
		box.Append(rb, false)
		button := ui.NewButton("Greet")
		box.Append(button, false)
		greeting := ui.NewLabel("")
		box.Append(greeting, false)
		window := ui.NewWindow("Hello RadioButtons", 350, 100, false)
		window.SetMargined(true)
		window.SetChild(box)
		button.OnClicked(func(*ui.Button) {
			// Does not seem to be any way to get Values from a Radiobox
			//	greeting.SetText("Hello, " + rb.Text() + "!")
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
