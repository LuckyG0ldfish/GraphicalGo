package subelements

import (
	"fmt"
	"image"
	"image/color"
	"strconv"

	g "github.com/AllenDang/giu"

	"github.com/LuckyG0ldfish/GraphicalGo/elements"
)

type File struct {
	name        string
	level       int
	Imports     string
	Objects     []*Object
	Functions   []string
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
	case 3: // Type Object
		obj, er := e.(*Object)
		if er {
			obj.level = fil.level + 1
			obj.parent = fil
			fil.Objects = append(fil.Objects, obj)
			NotifyOfSizeChange(obj.parent)
		}
	case 4: // Type Variable
	}

	fil.addingState = false
	fmt.Println("file adding done")
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
	for _, v := range fil.Objects {
		r = append(r, v)
	}
	// add Variable
	// for _, v := range fil.files {
	// 	r = append(r, v)
	// }
	return r
}

func CreateFiles(name string, x int) *File {
	var fil File
	fil.name = name
	fil.level = 1
	fil.id = GetNextID()

	fil.xLeft = x
	fil.yTop = 100

	fil.xRight = x + 110
	fil.yBot = 140

	fil.xRelative = 0
	fil.yRelative = 0

	// fil.objBut = CreateButton(" ", 0, fil.Expand)
	// fil.varBut = CreateButton(" ", 0, fil.Expand)
	pro.Can.Dragables = append(pro.Can.Dragables, &fil)
	pro.Can.Expandables = append(pro.Can.Expandables, &fil)
	pro.Can.Add = append(pro.Can.Add, &fil)
	pro.Files = append(pro.Files, &fil)
	return &fil
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
	return 2 // Type of File
}

func (fil *File) GetLevel() int {
	return fil.level
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

func (fil *File) Expand() {
	NotifyOfSizeChange(fil.parent)
}
