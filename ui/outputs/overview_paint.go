package outputs

import (
	"image"
	"image/color"

	g "github.com/AllenDang/giu"

	// "github.com/AllenDang/imgui-go"
	"github.com/LuckyG0ldfish/GraphicalGo/context"
)

const OverviewXDist int = 10 
const OverviewYDist int = 20 
const OverviewOffset int = 5 

func AddOverView(c *g.Canvas) {
	Pro := context.GetPro()
	xLeft := Pro.Over.XLeft
	xRight := Pro.Over.XRight
	for i, v := range Pro.Over.Pressables {
		addOverViewElement(c, i, xLeft, xRight, v)
	}
}

func addOverViewElement(c *g.Canvas, i, xleft, xright int, ove context.OverViewElement) {
	pos := g.GetCursorScreenPos()
	
	frameXLeft := xleft + OverviewOffset + (ove.Level)*OverviewXDist
	frameXRight := xright - OverviewOffset
	frameYTop := OverviewOffset + i*OverviewYDist 
	frameYBot := frameYTop + 18 

	c.AddRectFilled(pos.Add(image.Pt(frameXLeft, frameYTop)), pos.Add(image.Pt(frameXRight, frameYBot)), color.White, 0, 5)
	c.AddRect(pos.Add(image.Pt(frameXLeft, frameYTop)), pos.Add(image.Pt(frameXRight, frameYBot)), color.Black, 0, 0, 1)
	c.AddText(pos.Add(image.Pt(frameXLeft+3, frameYTop+3)), color.Black, ove.Name)
}