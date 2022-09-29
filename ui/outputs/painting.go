package outputs

// // #include "imguiWrapper.h"
// import "C"
import (
	"image"
	"image/color"

	g "github.com/AllenDang/giu"
	// "github.com/AllenDang/imgui-go"
	// "github.com/LuckyG0ldfish/GraphicalGo/elements"
	"github.com/LuckyG0ldfish/GraphicalGo/context"
)

func DrawUI() {
	Project := context.GetPro()
	// g.SingleWindowWithMenuBar().Layout(
	// 	g.MenuBar().Layout(
	// 		g.Button("New Object").OnClick(func() { Project.Can.Dragables = append(Project.Can.Dragables, subelements.CreateObject("test", 160)) }),	
	// 	), 
	g.SingleWindow().Layout(
		// g.Column(
			g.Custom(func() {
				canvas := g.GetCanvas()
				decideDrawOrder(canvas)

				// AddOverView(canvas)

				// for _, v := range Project.Can.Dragables {
				// 	AddDragable(canvas, v)
				// }

				for _, v := range Project.Obj.Pressables {
					AddButton(canvas, v)
				}
			}),
		// ),
		
	)
}

func AddSeperation(c *g.Canvas, XCord int) {
	Project := context.GetPro()
	pos := g.GetCursorScreenPos()
	c.AddLine(pos.Add(image.Pt(XCord, 0)), pos.Add(image.Pt(XCord, Project.Win.YHeight)), color.White, 1)
}


