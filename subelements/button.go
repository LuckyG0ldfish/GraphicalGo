package subelements

import (
	"image"
	"image/color"

	g "github.com/AllenDang/giu"
	"github.com/LuckyG0ldfish/GraphicalGo/context"
	"github.com/LuckyG0ldfish/GraphicalGo/elements"
)

const ButtonXLeftOffset int = 5
const ButtonXRightOffset int = 5

type Button struct {
	name     string
	pos      int
	function func()
	level    int

	xRight int
	yBot   int

	xLeft int
	yTop  int
}

func CreateButton(name string, pos int, fun func()) *Button {
	pro := context.GetPro()
	var but Button
	but.name = name
	but.pos = pos
	but.function = fun
	but.level = 1

	yVar := 10 + pos*context.Dist
	mix := pro.Can.XRight
	max := pro.Win.XWidth

	but.xLeft = mix + ButtonXLeftOffset
	but.xRight = max - ButtonXRightOffset

	but.yTop = yVar
	but.yBot = yVar + 20

	pro.Obj.Pressables = append(pro.Obj.Pressables, &but)

	return &but
}

func (but *Button) Draw(c *g.Canvas) {
	pos := g.GetCursorScreenPos()

	mix := but.GetXLeft()
	miy := but.GetYTop()

	max := but.GetXRight()
	may := but.GetYBot()

	c.AddRectFilled(pos.Add(image.Pt(mix, miy)), pos.Add(image.Pt(max, may)), color.White, 0, 5)
	c.AddText(pos.Add(image.Pt(mix+3, miy+3)), color.Black, but.GetName())
}

func (but *Button) GetSubelements() []elements.Element {
	r := make([]elements.Element, 0)
	// add Variable
	// for _, v := range fil.files {
	// 	r = append(r, v)
	// }
	return r
}

func (but *Button) Press() {
	but.function()
}

func (but *Button) GetPos() int {
	return but.pos
}

func (but *Button) GetName() string {
	return but.name
}

// func (but *Button) SetName(name string){
// 	but.name = name
// }

// func (but *Button) SetFunc(fun func) {
// 	but.func = fun
// }

// func ObjectButton(Pro Project) {
// 	Pro.Can.Dragables = append(Pro.Can.Dragables, CreateObject("test", 200))
// }

func (but *Button) GetXRight() int {
	return but.xRight
}

func (but *Button) GetYBot() int {
	return but.yBot
}

func (but *Button) GetXLeft() int {
	return but.xLeft
}

func (but *Button) GetYTop() int {
	return but.yTop
}

func (but *Button) SetXLeft(i int) {
	but.xLeft = i
}

func (but *Button) SetXRight(i int) {
	but.xRight = i
}

func (but *Button) SetYTop(i int) {
	but.yTop = i
}

func (but *Button) SetYBot(i int) {
	but.yBot = i
}

func (but *Button) SetXLeftObjectSelect(objXLeft int) {
	but.xLeft = objXLeft + ButtonXLeftOffset
}

func (but *Button) SetXRightObjectSelect(objXRight int) {
	but.xRight = objXRight - ButtonXRightOffset
}