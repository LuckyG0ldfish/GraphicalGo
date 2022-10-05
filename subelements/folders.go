package subelements

import (
	"image"
	"image/color"
	// "strconv"

	g "github.com/AllenDang/giu"
	"github.com/LuckyG0ldfish/GraphicalGo/context"
	"github.com/LuckyG0ldfish/GraphicalGo/elements"
)

const FolderType = 1
const FolderBaseWidth = 110
const FolderBaseHeight = 40 


type Folder struct {
	name        string
	level       int
	path        string
	files       []*File
	folders     []*Folder
	addingState bool
	updatedSize bool
	parent      elements.Element

	xLeft int
	yTop  int

	xRight int
	yBot   int

	xRelative int
	yRelative int

	id int
}

func CreateFolders(name string, x int) *Folder {
	pro := context.GetPro()
	var folder Folder
	folder.name = name
	folder.level = 1
	folder.updatedSize = false
	folder.id = context.GetNextID()

	folder.xLeft = x
	folder.yTop = 100

	folder.xRight = x + FolderBaseWidth
	folder.yBot = folder.yTop + FolderBaseHeight

	folder.xRelative = 0
	folder.yRelative = 0

	pro.Can.Dragables = append(pro.Can.Dragables, &folder)
	pro.Can.Expandables = append(pro.Can.Expandables, &folder)
	pro.Can.Add = append(pro.Can.Add, &folder)
	// pro.Folders = append(pro.Folders, &folder)
	pro.Level1 = append(pro.Level1, &folder)
	return &folder
}

func (fol *Folder) Draw(c *g.Canvas) {
	pos := g.GetCursorScreenPos()

	mix := fol.GetXLeft()
	miy := fol.GetYTop()

	max := fol.GetXRight()
	may := fol.GetYBot()

	// name := strconv.Itoa(fol.GetID()) + " " + strconv.Itoa(fol.GetLevel())
	c.AddRectFilled(pos.Add(image.Pt(mix, miy)), pos.Add(image.Pt(mix+110, miy+20)), color.White, 0, 5)
	c.AddRect(pos.Add(image.Pt(mix, miy)), pos.Add(image.Pt(mix+110, miy+20)), color.Black, 0, 0, 1)
	c.AddRectFilled(pos.Add(image.Pt(mix, miy+20)), pos.Add(image.Pt(max, may)), context.Gray, 0, 5)
	c.AddRect(pos.Add(image.Pt(mix, miy+20)), pos.Add(image.Pt(max, may)), color.Black, 0, 0, 1)
	// c.AddLine(pos.Add(image.Pt(mix+3, miy+20)), pos.Add(image.Pt(mix+107, miy+20)), color.Black, 1)
	c.AddText(pos.Add(image.Pt(mix+3, miy+3)), color.Black, fol.GetName())

	// expand
	c.AddRectFilled(pos.Add(image.Pt(max-6, may-6)), pos.Add(image.Pt(max-1, may-1)), color.Black, 0, 5)
}

func (fol *Folder) GetSubelements() []elements.Element {
	r := make([]elements.Element, 0)
	for _, v := range fol.folders {
		r = append(r, v)
	}
	for _, v := range fol.files {
		r = append(r, v)
	}
	return r
}

// func (fol *Folder) GetSubelementsAddable() []elements.Addable {
// 	r := make([]elements.Element, 0)
// 	for _, v := range fol.folders {
// 		r = append(r, v)
// 	}
// 	for _, v := range fol.files {
// 		r = append(r, v)
// 	}
// 	return r
// }

func (fol *Folder) Adding(e elements.Element) {
	pro := context.GetPro()
	fol.addingState = false

	switch e.GetType() {
	case FolderType: 
		fol2, e := e.(*Folder)
		if e {
			context.RecursiveLevelChange(fol.level + 1, fol2) 
			fol2.parent = fol
			fol.folders = append(fol.folders, fol2)
		}
	case FileType: 
		fil, e := e.(*File)
		if e {
			context.RecursiveLevelChange(fol.level + 1, fil)
			fil.parent = fol
			fol.files = append(fol.files, fil)
		}
	}
	pro.Level1 = context.RemoveElement(pro.Level1, e)
	context.NotifyOfSizeChange(fol)
}

func (fol *Folder) Removing(e elements.Element) {
	
	if e.GetType() == FolderType {
		folder, er := e.(*Folder)
		if er {
			fol.folders = folder.removeFolder(fol.folders)
		}
	} else if e.GetType() == FileType {
		file, er := e.(*File)
		if er {
			fol.files = file.removeFile(fol.files)
		}
	}
	context.RecursiveLevelChange(1, e) // base level 
	e.SetParent(nil)
	context.GetPro().Level1 = append(context.GetPro().Level1, e)
	context.NotifyOfSizeChange(fol)
}

func (fol *Folder) removeFolder(e []*Folder) []*Folder {
	// empty or last removing 
	if len(e) < 2 {
		return make([]*Folder, 0)
	}

	ret := make([]*Folder, len(e)-1)
	tempI := 0 
	for _, v := range e {
		if v != fol {
			ret[tempI] = v
			tempI ++ 
		}
	}

	return ret
}

func (fol *Folder) Expand() {
	context.NotifyOfSizeChange(fol.parent)
}

func (fol *Folder) GetXLeft() int {
	return fol.xLeft
}

func (fol *Folder) GetXRight() int {
	return fol.xRight
}

func (fol *Folder) GetYTop() int {
	return fol.yTop
}

func (fol *Folder) GetYBot() int {
	return fol.yBot
}

func (fol *Folder) SetXLeft(x int) {
	fol.xLeft = x
}

func (fol *Folder) SetXRight(x int) {
	fol.xRight = x
}

func (fol *Folder) SetYTop(y int) {
	fol.yTop = y
}

func (fol *Folder) SetYBot(y int) {
	fol.yBot = y
}

func (fol *Folder) GetRelativeX() int {
	return fol.xRelative
}

func (fol *Folder) GetRelativeY() int {
	return fol.yRelative
}

func (fol *Folder) SetRelativeX(i int) {
	fol.xRelative = i
}

func (fol *Folder) SetRelativeY(i int) {
	fol.yRelative = i
}

func (fol *Folder) GetName() string {
	return fol.name
}

func (fol *Folder) GetID() int {
	return fol.id
}

func (fol *Folder) GetType() int {
	return FolderType // Type for Folder
}

func (fol *Folder) GetLevel() int {
	return fol.level
}

func (fol *Folder) SetLevel(i int) {
	fol.level = i
}

func (fol *Folder) GetAddingState() bool {
	return fol.addingState
}

func (fol *Folder) SetAddingState(a bool) {
	fol.addingState = a
}

func (fol *Folder) GetParent() elements.Element {
	return fol.parent
}

func (fol *Folder) SetParent(par elements.Element)  {
	fol.parent = par
}

func (fol *Folder) SetAsParent(child elements.Element)  {
	child.SetParent(fol)
}

func (fol *Folder) GetBaseHeight() int {
	return FolderBaseHeight
}

func (fol *Folder) GetBaseWidth() int {
	return FolderBaseWidth
}