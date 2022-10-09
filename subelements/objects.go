package subelements

import (
	"image"
	"image/color"

	g "github.com/AllenDang/giu"
	"github.com/LuckyG0ldfish/GraphicalGo/context"
	"github.com/LuckyG0ldfish/GraphicalGo/elements"
)

const ObjectType = 3
const ObjectBaseWidth = 100
const ObjectBaseHeight = 100

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
	level := 1
	id := context.GetNextID()
	yTop := 100

	return CreateAndInitObject(name, x, yTop, id, level)
}

func CreateAndInitObject(name string, x, y, id, lvl int) *Object {
	pro := context.GetPro()
	var object Object

	object.Name = name
	object.level = lvl
	object.id = id

	object.xLeft = x
	object.yTop = y

	object.xRight = x + ObjectBaseWidth
	object.yBot = y + ObjectBaseHeight

	object.xRelative = 0
	object.yRelative = 0

	pro.Can.Dragables = append(pro.Can.Dragables, &object)
	if lvl == 1 {
		pro.Level1 = append(pro.Level1, &object)
	}
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

func (ob *Object) Adding(e elements.Element) {}

func (ob *Object) Removing(e elements.Element) {
	if e.GetType() == VariableType {
		// folder, er := e.(*Folder)
		// if er {
		// 	removeFolder(ob.folders, folder)
		// }
	}
	context.RecursiveLevelChange(1, e) // base level
	e.SetParent(nil)
	context.GetPro().Level1 = append(context.GetPro().Level1, e)
	context.NotifyOfSizeChange(ob)
}

func (ob *Object) removeObject(e []*Object) []*Object {
	// empty or last removing
	if len(e) < 2 {
		return make([]*Object, 0)
	}

	ret := make([]*Object, len(e)-1)
	tempI := 0
	for _, v := range e {
		if v != ob {
			ret[tempI] = v
			tempI++
		}
	}

	return ret
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

func (ob *Object) SetParent(par elements.Element) {
	ob.parent = par
}

func (ob *Object) SetAsParent(child elements.Element) {
	child.SetParent(ob)
}

func (ob *Object) GetBaseHeight() int {
	return ObjectBaseHeight
}

func (ob *Object) GetBaseWidth() int {
	return ObjectBaseWidth
}
