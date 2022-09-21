package outputs

import (
	g "github.com/AllenDang/giu"
	"github.com/LuckyG0ldfish/GraphicalGo/elements"
	"github.com/LuckyG0ldfish/GraphicalGo/elements/subelements"
)

func decideDrawOrder(c *g.Canvas) {
	pro := subelements.GetPro()

	for _, v := range pro.Folders {
		if v.GetLevel() == 1 {
			recDraw(v, c)
		}
	}

	for _, v := range pro.Files {
		if v.GetLevel() == 1 {
			recDraw(v, c)
		}
	}


}

func recDraw(d elements.Drawable, c *g.Canvas) {
	// var draw elements.Drawable
	// var e bool
	// switch d.GetType() {
	// case 1: // ID Folder 
	// 	draw, e = d.(*subelements.Folder)
	// case 2: // ID File 
	// 	draw, e = d.(*subelements.File)
	// case 3: // ID Object 
	// 	draw, e = d.(*subelements.Object)
	// }
	
	// if e {
		d.Draw(c)
		for _, v := range d.GetSubelements() {
			recDraw(v, c)
		}
	// }
}