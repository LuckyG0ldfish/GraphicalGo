package ui

import (
	"fmt"
	"image"
	"image/color"
	"strconv"

	g "github.com/AllenDang/giu"
	// "github.com/AllenDang/imgui-go"
	"github.com/LuckyG0ldfish/GraphicalGo/elements"
	// imgui "github.com/AllenDang/imgui-go"
)

var (
	// testVar g.CustomWidget
	dragInProgress bool
	Dragable       elements.Dragable
)

func handleUI() {
	if LeftPressed() && !dragInProgress {
		fmt.Println("pressed")
		dragInProgress = true
		getRelativePos(Dragable)
	}
	if dragInProgress {
		fmt.Println("drag in progress")
		updateRelativePos(Dragable)
		g.Update()
	}
	if LeftReleased() {
		fmt.Println("release")
		updateRelativePos(Dragable)
		dragInProgress = false
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

	c.AddRectFilled(pos.Add(image.Pt(mix, miy)), pos.Add(image.Pt(mix+100, miy+100)), color.White, 0, 5)
	c.AddLine(pos.Add(image.Pt(mix+3, miy+20)), pos.Add(image.Pt(mix+97, miy+120)), color.Black, 1)
	c.AddText(pos.Add(image.Pt(mix+3, miy+3)), color.Black, "TestObject")
}

func AdjustObjectPosition(c *g.Canvas, drag elements.Dragable) {
	// x, y := CursorPos()
	// mix := drag.GetXLeft()
	// miy := drag.GetYTop()

	// rx := drag.GetRelativeX()
	// ry := drag.GetRelativeY()
}

func UpdateDragPos(D elements.Dragable) {
	working := getRelativePos(D)
	if !working {
		return
	}
	for {
		updateRelativePos(D)
		if !LeftPressed() {
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

	fmt.Println("relX: " + strconv.Itoa(relX) + " relY: " + strconv.Itoa(relY))
	D.SetRelativeX(relX)
	D.SetRelativeY(relY)
	fmt.Println("rel pos updated")
	return true
}

func updateRelativePos(D elements.Dragable) {
	curX, curY := CursorPos()

	fmt.Println("X: " + strconv.Itoa(curX))
	fmt.Println("Y: " + strconv.Itoa(curY))

	D.SetXLeft(curX - D.GetRelativeX())
	D.SetYTop(curY - D.GetRelativeY())

	fmt.Println("DX: " + strconv.Itoa(D.GetXLeft()))
	fmt.Println("DY: " + strconv.Itoa(D.GetYTop()))
}

// fix cursor in object
func isDragable(D elements.Dragable) bool {
	posX, posY := CursorPos()

	if posX > D.GetXLeft() && posY > D.GetYTop() { //posX < D.GetXDist()&& posY < D.GetYDist()
		return true
	}
	return false
}
