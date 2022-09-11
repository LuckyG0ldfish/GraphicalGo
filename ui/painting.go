package ui

import (
	"image"
	"image/color"

	g "github.com/AllenDang/giu"
	// "github.com/AllenDang/imgui-go"
	"github.com/LuckyG0ldfish/GraphicalGo/elements"
	"github.com/LuckyG0ldfish/GraphicalGo/elements/subelements"
)

func drawUI() {
	g.SingleWindowWithMenuBar().Layout(
		g.MenuBar().Layout(
			g.Button("New Object").OnClick(func() { Project.Can.Dragables = append(Project.Can.Dragables, subelements.CreateObject("test", 160)) }),	
		), 
	// g.SingleWindow().Layout(
		g.Column(
			g.Custom(func() {
				canvas := g.GetCanvas()
				AddOverView(canvas)
				AddObjectSelect(canvas)

				for _, v := range Project.Can.Dragables {
					AddObject(canvas, v)
				}
			}),
		),
		
	)
}

func AddObject(c *g.Canvas, drag elements.Dragable) {
	pos := g.GetCursorScreenPos()

	mix := drag.GetXLeft()
	miy := drag.GetYTop()

	max := drag.GetXRight()
	may := drag.GetYBot()

	c.AddRectFilled(pos.Add(image.Pt(mix, miy)), pos.Add(image.Pt(max, may)), color.White, 0, 5)
	c.AddLine(pos.Add(image.Pt(mix+3, miy+20)), pos.Add(image.Pt(max-3, miy+20)), color.Black, 1)
	c.AddText(pos.Add(image.Pt(mix+3, miy+3)), color.Black, drag.GetName())
}

func AddSeperation(c *g.Canvas, XCord int) {
	pos := g.GetCursorScreenPos()
	c.AddLine(pos.Add(image.Pt(XCord, 0)), pos.Add(image.Pt(XCord, Project.Win.YHeight)), color.White, 1)
}

func AddOverView(c *g.Canvas) {
	pos := g.GetCursorScreenPos()
	c.AddRectFilled(pos.Add(image.Pt(0, 0)), pos.Add(image.Pt(Project.Can.XLeft, Project.Win.YHeight)), color.White, 0, 5)
}

func AddObjectSelect(c *g.Canvas) {
	pos := g.GetCursorScreenPos()
	c.AddRectFilled(pos.Add(image.Pt(Project.Win.XWidth, 0)), pos.Add(image.Pt(Project.Can.XRight, Project.Win.YHeight)), color.White, 0, 5)
}