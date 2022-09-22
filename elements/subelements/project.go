package subelements

import "github.com/LuckyG0ldfish/GraphicalGo/elements"

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
	Folders []*Folder
	Files   []*File
}

type Window struct {
	XWidth  int
	YHeight int
}

type Canvas struct {
	XLeft       int
	XRight      int
	Dragables   []elements.Element
	Expandables []elements.Expandable
	Draw        []elements.Element
	Add         []elements.Addable
}

type OverView struct {
	XLeft      int
	XRight     int
	Pressables []elements.Pressable
}

type ObjectSelect struct {
	XLeft      int
	XRight     int
	Pressables []elements.Pressable
}

func NewProject(Name string, Width int, Height int) *Project {
	pro.Name = Name
	pro.Win = newWindow(Width, Height)
	pro.Can.XLeft = 100
	pro.Can.XRight = Width - 200
	pro.DragInProgress = false
	pro.nextID = 1
	trialSetup()
	return &pro
}

// only for testing
func trialSetup() {
	// pro.Can.Dragables = make([]elements.Element, 0)
	// pro.Can.Expandables = make([]elements.Expandable, 0)
	// pro.Obj.Pressables = make([]elements.Pressable, 0)

	// CreateObject("Test1 :Object", 100)
	// CreateObject("Test2 :Object", 300)
	// CreateObject("Test3 :Object", 500)
	CreateFolders("Test1 :Package", 700)
	CreateFiles("Test : File", 400)

	CreateButton("Create Package", 0, func() {
		CreateFolders("addedPackage", 700)
	})
	CreateButton("Create File", 1, func() {
		CreateFiles("addedFile", 400)
	})
	// CreateButton("Create Object", 2, func() {
	// 	CreateObject("addedObject", 200)
	// })
}

func newWindow(Width int, Height int) Window {
	var Win Window
	Win.XWidth = Width
	Win.YHeight = Height
	return Win
}

func (obj *ObjectSelect) AdjustButtonPositions() {
	newButtonXLeft := obj.XLeft + ButtonXLeftOffset
	newButtonXRight := obj.XRight - ButtonXRightOffset
	for _, v := range obj.Pressables {
		v.SetXLeft(newButtonXLeft)
		v.SetXRight(newButtonXRight)
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