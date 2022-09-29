package subelements

import (
	"fmt"
	"image"
	"image/color"
	"strconv"

	g "github.com/AllenDang/giu"
	"github.com/LuckyG0ldfish/GraphicalGo/context"
	"github.com/LuckyG0ldfish/GraphicalGo/elements"
)

const FileType = 2
const FileBaseWidth = 110
const FileBaseHeight = 40

type File struct {
	name        string
	level       int
	imports     string
	objects     []*Object
	functions   []string
	id          int
	addingState bool
	parent      elements.Element

	objBut *Button
	varBut *Button

	xLeft int
	yTop  int

	xRight int
	yBot   int

	xRelative int
	yRelative int
}

func (fil *File) Adding(e elements.Element) {
	
	switch e.GetType() {
	case ObjectType: 
		obj, er := e.(*Object)
		if er {
			context.RecursiveLevelChange(fil.level + 1, obj)
			obj.parent = fil
			fil.objects = append(fil.objects, obj)
			context.NotifyOfSizeChange(obj.parent)
			fmt.Println("fileadding")
		} else {
			fmt.Println("failed to convert")
		}
	case VariableType: 
	}
	context.GetPro().Level1 = append(context.GetPro().Level1, e)
	fil.addingState = false
	
	fmt.Println("file adding done")
}

func (fil *File) Removing(e elements.Element) {
	if e.GetType() == ObjectType {
		obj, er := e.(*Object)
		if er {
			fil.objects = obj.removeObject(fil.objects)
		}
	} // else if e.GetType() == FileType {
	// 	file, er := e.(*File)
	// 	if er {
	// 		removeVariable(fil.variables, file)
	// 	}
	// }
	context.RecursiveLevelChange(1, e) // base level 
	e.SetParent(nil)
	context.GetPro().Level1 = append(context.GetPro().Level1, e)
	context.NotifyOfSizeChange(fil)
}

func (fil *File) Draw(c *g.Canvas) {
	pos := g.GetCursorScreenPos()

	mix := fil.GetXLeft()
	miy := fil.GetYTop()

	max := fil.GetXRight()
	may := fil.GetYBot()

	
	name := fil.GetName() + " " + strconv.Itoa(fil.GetLevel())
	c.AddTriangleFilled(pos.Add(image.Pt(mix+90, miy+1)), pos.Add(image.Pt(mix+90, miy+20)), pos.Add(image.Pt(mix+109, miy+20)), color.White)
	c.AddTriangle(pos.Add(image.Pt(mix+90, miy+1)), pos.Add(image.Pt(mix+90, miy+20)), pos.Add(image.Pt(mix+109, miy+20)), color.Black, 1)
	c.AddRectFilled(pos.Add(image.Pt(mix, miy)), pos.Add(image.Pt(mix+90, miy+20)), color.White, 0, 5)
	c.AddRect(pos.Add(image.Pt(mix, miy)), pos.Add(image.Pt(mix+90, miy+20)), color.Black, 0, 0, 1)
	c.AddRectFilled(pos.Add(image.Pt(mix, miy+20)), pos.Add(image.Pt(max, may)), color.White, 0, 5)
	c.AddRect(pos.Add(image.Pt(mix, miy+20)), pos.Add(image.Pt(max, may)), color.Black, 0, 0, 1)
	// c.AddLine(pos.Add(image.Pt(mix+3, miy+20)), pos.Add(image.Pt(mix+106, miy+20)), color.Black, 1)
	c.AddText(pos.Add(image.Pt(mix+3, miy+3)), color.Black, name)

	// expand
	c.AddRectFilled(pos.Add(image.Pt(max-6, may-6)), pos.Add(image.Pt(max-1, may-1)), color.Black, 0, 5)
}

func (fil *File) GetSubelements() []elements.Element {
	r := make([]elements.Element, 0)
	for _, v := range fil.objects {
		r = append(r, v)
	}
	// add Variable
	// for _, v := range fil.files {
	// 	r = append(r, v)
	// }
	return r
}

func CreateFiles(name string, x int) *File {
	pro := context.GetPro()
	var fil File
	fil.name = name
	fil.level = 1
	fil.id = context.GetNextID()

	fil.xLeft = x
	fil.yTop = 100

	fil.xRight = x + FileBaseWidth
	fil.yBot = fil.yTop + FileBaseHeight

	fil.xRelative = 0
	fil.yRelative = 0

	// fil.objBut = CreateButton(" ", 0, fil.Expand)
	// fil.varBut = CreateButton(" ", 0, fil.Expand)
	pro.Can.Dragables = append(pro.Can.Dragables, &fil)
	pro.Can.Expandables = append(pro.Can.Expandables, &fil)
	pro.Can.Add = append(pro.Can.Add, &fil)
	pro.Level1 = append(pro.Level1, &fil)
	// pro.Files = append(pro.Files, &fil)
	return &fil
}

func (fil *File) removeFile(e []*File) []*File {
	// empty or last removing 
	if len(e) < 2 {
		return make([]*File, 0)
	}

	ret := make([]*File, len(e)-1)
	tempI := 0 
	for _, v := range e {
		if v != fil {
			ret[tempI] = v
			tempI ++ 
		}
	}

	return ret
}

func (fil *File) GetXLeft() int {
	return fil.xLeft
}

func (fil *File) GetXRight() int {
	return fil.xRight
}

func (fil *File) GetYTop() int {
	return fil.yTop
}

func (fil *File) GetYBot() int {
	return fil.yBot
}

func (fil *File) SetXLeft(x int) {
	fil.xLeft = x
}

func (fil *File) SetXRight(x int) {
	fil.xRight = x
}

func (fil *File) SetYTop(y int) {
	fil.yTop = y
}

func (fil *File) SetYBot(y int) {
	fil.yBot = y
}

func (fil *File) GetID() int {
	return fil.id
}

func (fil *File) GetType() int {
	return FileType // Type of File
}

func (fil *File) GetLevel() int {
	return fil.level
}

func (fil *File) SetLevel(i int) {
	fil.level = i
}

func (fil *File) GetName() string {
	return fil.name
}

func (fil *File) GetRelativeX() int {
	return fil.xRelative
}

func (fil *File) GetRelativeY() int {
	return fil.yRelative
}

func (fil *File) SetRelativeX(i int) {
	fil.xRelative = i
}

func (fil *File) SetRelativeY(i int) {
	fil.yRelative = i
}

func (fil *File) GetAddingState() bool {
	return fil.addingState
}

func (fil *File) SetAddingState(a bool) {
	fil.addingState = a
}

func (fil *File) GetParent() elements.Element {
	return fil.parent
}

func (fil *File) SetParent(par elements.Element)  {
	fil.parent = par 
}

func (fil *File) SetAsParent(child elements.Element)  {
	child.SetParent(fil)
}

func (fil *File) Expand() {
	context.NotifyOfSizeChange(fil.parent)
}

func (fil *File) GetBaseHeight() int {
	return FileBaseHeight
}

func (fil *File) GetBaseWidth() int {
	return FileBaseWidth
}