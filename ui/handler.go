package ui

import (
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
	Dragable       elements.Dragable
)

func handleUI() {
	if LeftPressed() && isDragable(Dragable){
		go UpdateDragPos(Dragable)
	}
	g.SingleWindow().Layout(
		g.Custom(func() {
			canvas := g.GetCanvas()

			AddObject(canvas, Dragable)
			
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

func UpdateDragPos(D elements.Dragable) {
	working := getRelativePos(D)
	if !working {
		return
	}
	for {
		updateRelativePos(D)
		if LeftReleased() {
			return
		}
	}
}

func getRelativePos(D elements.Dragable) (working bool) {
	if !isDragable(D) {
		return false
	}
	posX, posY := CursorPos()
	relX := posX - D.GetXLeft()
	relY := posY - D.GetYTop()

	// fmt.Println("relX: " + strconv.Itoa(relX) + " relY: " + strconv.Itoa(relY))
	D.SetRelativeX(relX)
	D.SetRelativeY(relY)
	// fmt.Println("rel pos updated")
	return true
}

func updateRelativePos(D elements.Dragable) {
	curX, curY := CursorPos()

	mix := D.GetXLeft()
	miy := D.GetYTop()

	max := D.GetXRight()
	may := D.GetYBot()

	D.SetXLeft(curX - D.GetRelativeX())
	D.SetYTop(curY - D.GetRelativeY())

	D.SetXRight(D.GetXLeft()+ (max-mix))
	D.SetYBot(D.GetYTop()+ (may-miy))
}

// fix cursor in object
func isDragable(D elements.Dragable) bool {
	posX, posY := CursorPos()

	if posX > D.GetXLeft() && posY > D.GetYTop() && posX < D.GetXRight()&& posY < D.GetYBot(){ 
		return true
	}
	return false
}
