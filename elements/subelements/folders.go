package subelements

import (
	"fmt"
	"image"
	"image/color"

	g "github.com/AllenDang/giu"

	"github.com/LuckyG0ldfish/GraphicalGo/elements"
) 

type Folder struct {
	name string 
	level int
	path string
	files []*File
	folders []*Folder
	addingState	bool 
	updatedSize bool 
	parent elements.Drawable

	xLeft 	int 
	yTop	int 

	xRight 	int 
	yBot	int 

	xRelative int 
	yRelative int 

	id 		int
}

func CreateFolders(name string, x int) (*Folder) {
	var folder Folder
	folder.name = name
	folder.level = 1
	folder.updatedSize = false 
	
	folder.xLeft = x
	folder.yTop = 100
	
	folder.xRight = x+300
	folder.yBot = 300

	folder.xRelative = 0
	folder.yRelative = 0 

	pro.Can.Dragables = append(pro.Can.Dragables, &folder) 
	pro.Can.Expandables = append(pro.Can.Expandables, &folder)
	pro.Can.Add = append(pro.Can.Add, &folder)
	pro.Folders = append(pro.Folders, &folder)
	return &folder
}

func (fol *Folder) Draw(c *g.Canvas) {
	pos := g.GetCursorScreenPos()

	mix := fol.GetXLeft()
	miy := fol.GetYTop()

	max := fol.GetXRight()
	may := fol.GetYBot()

	c.AddRectFilled(pos.Add(image.Pt(mix, miy)), pos.Add(image.Pt(mix+110, miy+20)), color.White, 0, 5)
	c.AddRectFilled(pos.Add(image.Pt(mix, miy+20)), pos.Add(image.Pt(max, may)), color.White, 0, 5)
	c.AddLine(pos.Add(image.Pt(mix+3, miy+20)), pos.Add(image.Pt(mix+107, miy+20)), color.Black, 1)
	c.AddText(pos.Add(image.Pt(mix+3, miy+3)), color.Black, fol.GetName())

	// expand
	c.AddRectFilled(pos.Add(image.Pt(max-6, may-6)), pos.Add(image.Pt(max-1, may-1)), color.Black, 0, 5)
}

func (fol *Folder) GetSubelements() []elements.Drawable {
	r := make([]elements.Drawable, 0)
	for _, v := range fol.folders {
		r = append(r, v)
	}
	for _, v := range fol.files {
		r = append(r, v)
	}
	return r
}

func (fol *Folder) Adding(e elements.Dragable) {
	fol.addingState = false 
	fmt.Println("file adding done")

	switch e.GetType() {
		case 1: // Type Folder 
			fol2, e := e.(*Folder)
			if e {
				fol2.level = fol.level + 1
				fol2.parent = fol 
				fol.folders = append(fol.folders, fol2)
			}
		case 2: // Type File 
			fil, e := e.(*File)
			if e {
				fil.level = fol.level + 1
				fil.parent = fol
				fol.files = append(fol.files, fil)
			}
	}
}

func (fol *Folder) Expand() {}

func (fol *Folder) GetXLeft() int{
	return fol.xLeft
}

func (fol *Folder) GetXRight() int{
	return fol.xRight
}

func (fol *Folder) GetYTop() int{
	return fol.yTop
}

func (fol *Folder) GetYBot() int{
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

func (fol *Folder) GetRelativeX() int{
	return fol.xRelative
}

func (fol *Folder) GetRelativeY() int{
	return fol.yRelative
}

func (fol *Folder) SetRelativeX(i int) {
	fol.xRelative = i 
}

func (fol *Folder) SetRelativeY(i int) {
	fol.yRelative = i 
}

func (fol *Folder) GetName() string{
	return fol.name
}

func (fol *Folder) GetID() int{
	return fol.id
}

func (fol *Folder) GetType() int{
	return 1	// Type for Folder
}

func (fol *Folder) GetLevel() int{
	return fol.level
}

func (fol *Folder) GetAddingState() bool {
	return fol.addingState
}

func (fol *Folder) SetAddingState(a bool) {
	fol.addingState = a
}

func (fol *Folder) GetParent() elements.Drawable {
	return fol.parent
}
