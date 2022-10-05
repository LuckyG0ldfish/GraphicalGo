package inputs

import (
	"github.com/LuckyG0ldfish/GraphicalGo/elements"
	"github.com/LuckyG0ldfish/GraphicalGo/context"
	// "github.com/LuckyG0ldfish/GraphicalGo/elements/subelements"
)

func UpdateDragPos(D elements.Element) {
	updateRelPosRec(D)
	pro := context.GetPro()
	pro.Can.Dragged = D
	
	pare := D.GetParent()
	if pare != nil {
		pare.Removing(D)
	}
	pro.DragInProgress = true 
	for {
		updatePosRec(D)
		test, add := checkOnAdding(D)
		if test {
			if !add.GetAddingState() {
				add.SetAddingState(true)
			}
		// } else {
		// 	// reverseAllAddingStates()
		}
		
		if LeftReleased() {
			// pro.DragInProgress = false
			if test {
				add.Adding(D)
			} 
			pro.DragInProgress = false 
			return
		}
	}
}

func updateRelPosRec(D elements.Element) {
	working := updateRelativePos(D)
	if !working {
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

	D.SetRelativeX(relX)
	D.SetRelativeY(relY)
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
