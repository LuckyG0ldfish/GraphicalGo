package outputs

import (
	"image"
	"image/color"

	g "github.com/AllenDang/giu"

	// "github.com/AllenDang/imgui-go"
	"github.com/LuckyG0ldfish/GraphicalGo/context"
	"github.com/LuckyG0ldfish/GraphicalGo/elements"
)

func AddButton(c *g.Canvas, pres elements.Pressable) {
	pos := g.GetCursorScreenPos()
	
	mix := pres.GetXLeft()
	miy := pres.GetYTop()

	max := pres.GetXRight()
	may := pres.GetYBot()

	c.AddRectFilled(pos.Add(image.Pt(mix, miy)), pos.Add(image.Pt(max, may)), color.White, 0, 5)
	c.AddText(pos.Add(image.Pt(mix+3, miy+3)), color.Black, pres.GetName())
}

func AddObjectSelect(c *g.Canvas) {
	Project := context.GetPro()
	pos := g.GetCursorScreenPos()
	c.AddRectFilled(pos.Add(image.Pt(Project.Win.XWidth, -20)), pos.Add(image.Pt(Project.Can.XRight, Project.Win.YHeight)), color.White, 0, 5)
}