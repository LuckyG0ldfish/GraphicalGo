package ui

import (
	// "fmt"
	// "fmt"
	"image"
	"image/color"

	// "strconv"

	g "github.com/AllenDang/giu"
	// "github.com/AllenDang/imgui-go"
	"github.com/LuckyG0ldfish/GraphicalGo/elements"
	// imgui "github.com/AllenDang/imgui-go"
)

var (
	// testVar g.CustomWidget
	// dragInProgress bool
	// Dragable       elements.Dragable
	Dragables 		[]elements.Dragable
)

func handleUI() {
	if LeftPressed() {
		for _, v := range Dragables {
			if isDragable(v){
				go UpdateDragPos(v)
			}
		}
	}
	
	g.SingleWindow().Layout(
		g.Custom(func() {
			canvas := g.GetCanvas()	
			
			for _, v := range Dragables {
				AddObject(canvas, v)
			}
		}),
	)
}

func handleClickToCanvas() {}

func handleClickToOverView() {}

func handleClickToObjectSelect() {}

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

