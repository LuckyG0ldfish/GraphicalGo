package subelements

import "github.com/LuckyG0ldfish/GraphicalGo/elements"

type Project struct {
	Name string
	Win Window
	Can Canvas
	Over OverView
	Obj ObjectSelect
}

type Window struct {
	XWidth int 
	YHeight int 
}

type Canvas struct {
	XLeft int 
	XRight int 
	Dragables []elements.Dragable
	Draw []elements.Drawable
}

type OverView struct {
	
}

type ObjectSelect struct {

}

func NewProject(Name string, Width int, Height int) *Project{
	var Pro Project
	Pro.Name = Name
	Pro.Win = newWindow(Width, Height) 
	Pro.Can.Dragables = make([]elements.Dragable, 3)
	Pro.Can.XLeft = 0 
	Pro.Can.XRight = Width
	Pro.Can.Dragables[0] = CreateObject("Test :Object", 100)
	Pro.Can.Dragables[1] = CreateObject("Test :Object2", 300)
	Pro.Can.Dragables[2] = CreateObject("Test :Object3", 500)
	return &Pro
}

func newWindow(Width int, Height int) Window {
	var Win Window  
	Win.XWidth = Width
	Win.YHeight = Height
	return Win 
}