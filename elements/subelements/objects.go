package subelements

import (
	"image"
	"image/color"

	g "github.com/AllenDang/giu"

	"github.com/LuckyG0ldfish/GraphicalGo/elements"
)

const ObjectType = 3

type Object struct {
	Name      string
	level     int
	Variables []*Variable
	parent    elements.Element

	xLeft int
	yTop  int

	xRight int
	yBot   int

	xRelative int
	yRelative int

	id int
}

func CreateObject(name string, x int) *Object {
	var object Object
	object.Name = name
	object.level = 1
	object.id = GetNextID()

	object.xLeft = x
	object.yTop = 100

	object.xRight = x + 100
	object.yBot = 200

	object.xRelative = 0
	object.yRelative = 0

	pro.Can.Dragables = append(pro.Can.Dragables, &object)

	return &object
}

func (ob *Object) Draw(c *g.Canvas) {
	pos := g.GetCursorScreenPos()

	mix := ob.GetXLeft()
	miy := ob.GetYTop()

	max := ob.GetXRight()
	may := ob.GetYBot()

	c.AddRectFilled(pos.Add(image.Pt(mix, miy)), pos.Add(image.Pt(max, may)), color.White, 0, 5)
	c.AddRect(pos.Add(image.Pt(mix, miy)), pos.Add(image.Pt(max, may)), color.Black, 0, 0, 1)
	c.AddLine(pos.Add(image.Pt(mix+3, miy+20)), pos.Add(image.Pt(max-3, miy+20)), color.Black, 1)
	c.AddText(pos.Add(image.Pt(mix+3, miy+3)), color.Black, ob.GetName())
}

func (ob *Object) GetSubelements() []elements.Element {
	r := make([]elements.Element, 0)
	// add Variable
	// for _, v := range fil.files {
	// 	r = append(r, v)
	// }
	return r
}

func (ob *Object) Removing(e elements.Element) {
	if e.GetType() == VariableType {
		// folder, er := e.(*Folder)
		// if er {
		// 	removeFolder(ob.folders, folder)
		// }
	} 
	NotifyOfSizeChange(ob)
}

func (ob *Object) GetName() string {
	return ob.Name
}

func (ob *Object) GetRelativeX() int {
	return ob.xRelative
}

func (ob *Object) GetRelativeY() int {
	return ob.yRelative
}

func (ob *Object) SetRelativeX(i int) {
	ob.xRelative = i

}

func (ob *Object) SetRelativeY(i int) {
	ob.yRelative = i
}

func (ob *Object) GetXLeft() int {
	return ob.xLeft
}

func (ob *Object) GetYTop() int {
	return ob.yTop
}

func (ob *Object) SetXLeft(x int) {
	ob.xLeft = x
}

func (ob *Object) SetYTop(y int) {
	ob.yTop = y
}

func (ob *Object) GetXRight() int {
	return ob.xRight
}

func (ob *Object) GetYBot() int {
	return ob.yBot
}

func (ob *Object) SetXRight(x int) {
	ob.xRight = x
}

func (ob *Object) SetYBot(y int) {
	ob.yBot = y
}

func (ob *Object) GetID() int {
	return ob.id
}

func (ob *Object) GetType() int {
	return ObjectType // type of object
}

func (ob *Object) GetLevel() int {
	return ob.level
}

func (ob *Object) SetLevel(i int) {
	ob.level = i
}

func (ob *Object) GetParent() elements.Element {
	return ob.parent
}

func (ob *Object) SetParent(par elements.Element)  {
	ob.parent = par
}

func (ob *Object) SetAsParent(child elements.Element)  {
	child.SetParent(ob)
}
