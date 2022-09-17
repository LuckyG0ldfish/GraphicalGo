package inputs

import (
	"github.com/LuckyG0ldfish/GraphicalGo/elements"
	// "github.com/LuckyG0ldfish/GraphicalGo/elements/subelements"
)

func UpdateDragPos(D elements.Dragable) {
	working := getRelativePos(D)
	if !working {
		return
	}
	// pro := subelements.GetPro()
	// pro.DragInProgress = true 
	for {
		updateRelativePos(D)
		test, add := checkOnAdding(D)
		if test {
			if !add.GetAddingState() {
				add.Adding()
			}
		}
		if LeftReleased() {
			// pro.DragInProgress = false 
			add.ReverseAdding()
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
	// fmt.Println("")
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
