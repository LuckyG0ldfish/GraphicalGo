package outputs

import (
	"image"
	"image/color"

	g "github.com/AllenDang/giu"

	// "github.com/AllenDang/imgui-go"
	"github.com/LuckyG0ldfish/GraphicalGo/elements"
	// "github.com/LuckyG0ldfish/GraphicalGo/elements/subelements"
)

func AddDragable(c *g.Canvas, drag elements.Element) {
	typeID := drag.GetType()
	switch typeID {
	case 1: // ID Folder
		AddFolder(c, drag)
	case 2: // ID File
		AddFile(c, drag)
	case 3: // ID Object
		AddObject(c, drag)
	}
}

func AddFolder(c *g.Canvas, drag elements.Element) {
	pos := g.GetCursorScreenPos()

	mix := drag.GetXLeft()
	miy := drag.GetYTop()

	max := drag.GetXRight()
	may := drag.GetYBot()

	c.AddRectFilled(pos.Add(image.Pt(mix, miy)), pos.Add(image.Pt(mix+110, miy+20)), color.White, 0, 5)
	c.AddRectFilled(pos.Add(image.Pt(mix, miy+20)), pos.Add(image.Pt(max, may)), color.White, 0, 5)
	c.AddLine(pos.Add(image.Pt(mix+3, miy+20)), pos.Add(image.Pt(mix+107, miy+20)), color.Black, 1)
	c.AddText(pos.Add(image.Pt(mix+3, miy+3)), color.Black, drag.GetName())

	// expand
	c.AddRectFilled(pos.Add(image.Pt(max-6, may-6)), pos.Add(image.Pt(max-1, may-1)), color.Black, 0, 5)
}

func AddFile(c *g.Canvas, drag elements.Element) {

}

func AddObject(c *g.Canvas, drag elements.Element) {
	pos := g.GetCursorScreenPos()

	mix := drag.GetXLeft()
	miy := drag.GetYTop()

	max := drag.GetXRight()
	may := drag.GetYBot()

	c.AddRectFilled(pos.Add(image.Pt(mix, miy)), pos.Add(image.Pt(max, may)), color.White, 0, 5)
	c.AddLine(pos.Add(image.Pt(mix+3, miy+20)), pos.Add(image.Pt(max-3, miy+20)), color.Black, 1)
	c.AddText(pos.Add(image.Pt(mix+3, miy+3)), color.Black, drag.GetName())
}
