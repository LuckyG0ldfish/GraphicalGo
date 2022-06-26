package main

import (
	// "github.com/AllenDang/giu"

	// "github.com/LuckyG0ldfish/GraphicalGo/ui"#

	app "fyne.io/fyne/v2/app"
	container "fyne.io/fyne/v2/container"
	widget "fyne.io/fyne/v2/widget"
)

func main() {

	// wnd := giu.NewMasterWindow("Splitter", 800, 600, 0)
	// ui.Build(wnd)
	test := "TestName"
	createProject(test)

}

func createProject(name string) {
	a := app.New()
	w := a.NewWindow(name)

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))

	w.ShowAndRun()
	
}





