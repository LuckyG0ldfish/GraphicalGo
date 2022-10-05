package outputs

import (
	g "github.com/AllenDang/giu"
	"github.com/LuckyG0ldfish/GraphicalGo/context"
	"github.com/LuckyG0ldfish/GraphicalGo/elements"
)

func decideDrawOrder(c *g.Canvas) {
	pro := context.GetPro()
	dstatus := pro.DragInProgress
	for _, v := range pro.Level1 {
		if dstatus {
			if v.GetLevel() == 1 && v != pro.Can.Dragged {
				recDraw(v, c)
			}
		} else {
			if v.GetLevel() == 1 {
				recDraw(v, c)
			}
		}
		
	}
	if pro.DragInProgress {
		recDraw(pro.Can.Dragged, c)
	}
}

func recDraw(d elements.Element, c *g.Canvas) {
	d.Draw(c)
	for _, v := range d.GetSubelements() {
		recDraw(v, c)
	}
}
