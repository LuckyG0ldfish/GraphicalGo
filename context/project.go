package context

import (
	"github.com/LuckyG0ldfish/GraphicalGo/elements"
)

var pro Project
var Dist int = 40

type Project struct {
	Name           string
	DragInProgress bool
	nextID 	int 
	Win     Window
	Can     Canvas
	Over    OverView
	Obj     ObjectSelect
	Level1 	[]elements.Element
	SaveLoaded bool 
}

type Window struct {
	XWidth  int
	YHeight int
}

type Canvas struct {
	XLeft       int
	XRight      int
	Dragged 	elements.Element
	Dragables   []elements.Element
	Expandables []elements.Expandable
	Draw        []elements.Element
	Add         []elements.Addable
}

type OverView struct {
	XLeft      int
	XRight     int
	Pressables []OverViewElement
}

type OverViewElement struct {
	Name 	string 
	Level 	int 
}

type ObjectSelect struct {
	XLeft      int
	XRight     int
	Pressables []elements.Pressable
}

func NewProject(Name string, Width int, Height int) *Project {
	pro.Name = Name
	pro.Win = newWindow(Width, Height)
	pro.Can.XLeft = 200
	pro.Can.XRight = Width - 200
	pro.DragInProgress = false
	pro.nextID = 1
	pro.Over.XLeft = 0 
	pro.Over.XRight = pro.Can.XLeft
	pro.SaveLoaded = false 
	return &pro
}

func NewOverViewElement(name string, level int) {
	var ov OverViewElement
	ov.Level = level
	ov.Name = name
	pro.Over.Pressables = append(pro.Over.Pressables, ov)
}

func newWindow(Width int, Height int) Window {
	var Win Window
	Win.XWidth = Width
	Win.YHeight = Height
	return Win
}

func (obj *ObjectSelect) AdjustButtonPositions() {
	for _, v := range obj.Pressables {
		v.SetXLeftObjectSelect(obj.XLeft)
		v.SetXRightObjectSelect(obj.XRight)
	}
} 

func GetPro() *Project {
	return &pro
}

func GetNextID() int {
	i := pro.nextID
	pro.nextID++
	return i 
}

func (pro *Project) SetNextID(i int) {
	pro.nextID = i 
}