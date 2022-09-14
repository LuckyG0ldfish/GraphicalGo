package subelements

import "github.com/LuckyG0ldfish/GraphicalGo/elements"

var pro Project
var Dist int = 40 

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
	XLeft int 
	XRight int 
	Pressables []elements.Pressable
}

type ObjectSelect struct {
	XLeft int 
	XRight int 
	Pressables []elements.Pressable
}

func NewProject(Name string, Width int, Height int) *Project{
	pro.Name = Name
	pro.Win = newWindow(Width, Height) 
	pro.Can.XLeft = 100 
	pro.Can.XRight = Width-200
	trialSetup()
	return &pro
}

// only for testing
func trialSetup() {
	pro.Can.Dragables = make([]elements.Dragable, 3)
	pro.Obj.Pressables = make([]elements.Pressable, 2)
	pro.Can.Dragables[0] = CreateObject("Test :Object", 100)
	pro.Can.Dragables[1] = CreateObject("Test :Object2", 300)
	pro.Can.Dragables[2] = CreateObject("Test :Object3", 500)
	pro.Obj.Pressables[0] = CreateButton("Create Object", 0, func(){
		pro.Can.Dragables = append(pro.Can.Dragables, CreateObject("test", 200)) 
	})
	pro.Obj.Pressables[1] = CreateButton("Create Object 2", 1, func(){
		pro.Can.Dragables = append(pro.Can.Dragables, CreateObject("test", 200)) 
	})
}

func newWindow(Width int, Height int) Window {
	var Win Window  
	Win.XWidth = Width
	Win.YHeight = Height
	return Win 
}

func GetPro() *Project {
	return &pro
}