package outputs

import (
	"image"
	"image/color"
	g "github.com/AllenDang/giu"
	// "github.com/AllenDang/imgui-go"
	"github.com/LuckyG0ldfish/GraphicalGo/elements"
	"github.com/LuckyG0ldfish/GraphicalGo/elements/subelements"
)

const OverviewXDist int = 8 
const OverviewYDist int = 10 

func AddOverView(c *g.Canvas) {
	Pro := subelements.GetPro()
	for _, v := range Pro.Over.Pressables {
		addOverViewElement(c, v)
	}
	
}

func addOverViewElement(c *g.Canvas, pres elements.Pressable) {
	pos := g.GetCursorScreenPos()
	c.AddRectFilled(pos.Add(image.Pt(-20, -20)), pos.Add(image.Pt(0, 0)), color.White, 0, 5)
}