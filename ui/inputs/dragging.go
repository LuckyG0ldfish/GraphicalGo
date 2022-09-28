package inputs

import (
	"fmt"

	"github.com/LuckyG0ldfish/GraphicalGo/elements"
	"github.com/LuckyG0ldfish/GraphicalGo/elements/subelements"
	// "github.com/LuckyG0ldfish/GraphicalGo/elements/subelements"
)

func UpdateDragPos(D elements.Element) {
	updateRelPosRec(D)
	pare := D.GetParent()
	if pare != nil {
		pare.Removing(D)
		subelements.RecursiveLevelChange(1, D) // base level 
		D.SetParent(nil)
	}
	for {
		updatePosRec(D)
		test, add := checkOnAdding(D)
		if test {
			if !add.GetAddingState() {
				add.SetAddingState(true)
				fmt.Println("add1")
			}
		} else {
			reverseAllAddingStates()
			// fmt.Println("add2")
		}
		
		if LeftReleased() {
			// pro.DragInProgress = false
			if test {
				add.Adding(D)
				fmt.Println("add3")
			} else {

			}
			return
		}
	}
}

func updateRelPosRec(D elements.Element) {
	working := updateRelativePos(D)
	if !working {
		fmt.Println("not working")
		return
	}
	for _, v := range D.GetSubelements() {
		updateRelPosRec(v)
	}
}

func updateRelativePos(D elements.Element) (working bool) {
	// if !isDragable(D) {
	// 	return false
	// }
	posX, posY := CursorPos()
	relX := posX - D.GetXLeft()
	relY := posY - D.GetYTop()

	// fmt.Println("relX: " + strconv.Itoa(relX) + " relY: " + strconv.Itoa(relY))
	D.SetRelativeX(relX)
	D.SetRelativeY(relY)
	// fmt.Println("rel pos updated")
	return true
}

func updatePosRec(D elements.Element) {
	curX, curY := CursorPos()

	mix := D.GetXLeft()
	miy := D.GetYTop()

	max := D.GetXRight()
	may := D.GetYBot()

	D.SetXLeft(curX - D.GetRelativeX())
	D.SetYTop(curY - D.GetRelativeY())
	// fmt.Println("")
	D.SetXRight(D.GetXLeft() + (max - mix))
	D.SetYBot(D.GetYTop() + (may - miy))

	for _, v := range D.GetSubelements() {
		updatePosRec(v)
	}
}

func updatePos(D elements.Element) {
	curX, curY := CursorPos()

	mix := D.GetXLeft()
	miy := D.GetYTop()

	max := D.GetXRight()
	may := D.GetYBot()

	D.SetXLeft(curX - D.GetRelativeX())
	D.SetYTop(curY - D.GetRelativeY())
	// fmt.Println("")
	D.SetXRight(D.GetXLeft() + (max - mix))
	D.SetYBot(D.GetYTop() + (may - miy))
}

// fix cursor in object
func isDragable(D elements.Element) bool {
	posX, posY := CursorPos()

	if posX > D.GetXLeft() && posY > D.GetYTop() && posX < D.GetXRight() && posY < D.GetYBot() {
		return true
	}
	return false
}
