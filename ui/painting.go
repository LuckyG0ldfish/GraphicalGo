package ui

import (
	"image"
	"image/color"

	g "github.com/AllenDang/giu"
	"github.com/LuckyG0ldfish/GraphicalGo/elements"
)



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