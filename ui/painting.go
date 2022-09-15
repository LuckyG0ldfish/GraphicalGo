package ui

// // #include "imguiWrapper.h"
// import "C"
import (
	"image"
	"image/color"

	g "github.com/AllenDang/giu"
	// "github.com/AllenDang/imgui-go"
	"github.com/LuckyG0ldfish/GraphicalGo/elements"
	"github.com/LuckyG0ldfish/GraphicalGo/elements/subelements"
)

func drawUI() {
	Project := subelements.GetPro()
	// g.SingleWindowWithMenuBar().Layout(
	// 	g.MenuBar().Layout(
	// 		g.Button("New Object").OnClick(func() { Project.Can.Dragables = append(Project.Can.Dragables, subelements.CreateObject("test", 160)) }),	
	// 	), 
	g.SingleWindow().Layout(
		// g.Column(
			g.Custom(func() {
				canvas := g.GetCanvas()
				AddOverView(canvas)
				// AddObjectSelect(canvas)

				for _, v := range Project.Can.Dragables {
					AddDragable(canvas, v)
				}

				for _, v := range Project.Obj.Pressables {
					AddButton(canvas, v)
				}

				for _, v := range Project.Over.Pressables {
					AddButton(canvas, v)
				}
			}),
		// ),
		
	)
}

func AddDragable(c *g.Canvas, drag elements.Dragable) {
	typeID := drag.GetType()
	switch typeID {
	case 1:	// ID Folder 
		AddFolder(c, drag)
	case 2:	// ID File 
		AddFile(c, drag)
	case 3:	// ID Object 
		AddObject(c, drag)
	} 
}

func AddFolder(c *g.Canvas, drag elements.Dragable) {
	pos := g.GetCursorScreenPos()

	mix := drag.GetXLeft()
	miy := drag.GetYTop()

	max := drag.GetXRight()
	may := drag.GetYBot()

	c.AddRectFilled(pos.Add(image.Pt(mix, miy)), pos.Add(image.Pt(mix+100, miy+20)), color.White, 0, 5)
	c.AddRectFilled(pos.Add(image.Pt(mix, miy+20)), pos.Add(image.Pt(max, may)), color.White, 0, 5)
	c.AddLine(pos.Add(image.Pt(mix+3, miy+20)), pos.Add(image.Pt(max-3, miy+20)), color.Black, 1)
	c.AddText(pos.Add(image.Pt(mix+3, miy+3)), color.Black, drag.GetName())

	// expand
	c.AddRectFilled(pos.Add(image.Pt(max-6, may-6)), pos.Add(image.Pt(max-1, may-1)), color.Black, 0, 5)
}

func AddFile(c *g.Canvas, drag elements.Dragable) {
	
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

func AddButton(c *g.Canvas, pres elements.Pressable) {
	pos := g.GetCursorScreenPos()
	
	mix := pres.GetXLeft()
	miy := pres.GetYTop()

	max := pres.GetXRight()
	may := pres.GetYBot()

	c.AddRectFilled(pos.Add(image.Pt(mix, miy)), pos.Add(image.Pt(max, may)), color.White, 0, 5)
	c.AddText(pos.Add(image.Pt(mix+3, miy+3)), color.Black, pres.GetName())
}

func AddSeperation(c *g.Canvas, XCord int) {
	Project := subelements.GetPro()
	pos := g.GetCursorScreenPos()
	c.AddLine(pos.Add(image.Pt(XCord, 0)), pos.Add(image.Pt(XCord, Project.Win.YHeight)), color.White, 1)
}

func AddOverView(c *g.Canvas) {
	Project := subelements.GetPro()
	pos := g.GetCursorScreenPos()
	c.AddRectFilled(pos.Add(image.Pt(-20, -20)), pos.Add(image.Pt(Project.Can.XLeft, Project.Win.YHeight)), color.White, 0, 5)
}

func AddObjectSelect(c *g.Canvas) {
	Project := subelements.GetPro()
	pos := g.GetCursorScreenPos()
	c.AddRectFilled(pos.Add(image.Pt(Project.Win.XWidth, -20)), pos.Add(image.Pt(Project.Can.XRight, Project.Win.YHeight)), color.White, 0, 5)
}