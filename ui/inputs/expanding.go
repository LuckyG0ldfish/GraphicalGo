package inputs

import "github.com/LuckyG0ldfish/GraphicalGo/elements"

func UpdateExpanding(E elements.Expandable) {
	working := getRelativePosExp(E)
	if !working {
		return
	}
	for {
		updateExpandableSize(E)
		if LeftReleased() {
			E.Expand()
			return
		}
	}
}

func updateExpandableSize(E elements.Expandable) {
	curX, curY := CursorPos()

	E.SetXRight(curX + E.GetRelativeX())
	E.SetYBot(curY + E.GetRelativeY())
}

func getRelativePosExp(E elements.Expandable) (working bool) {
	if !isExpandable(E) {
		return false
	}
	posX, posY := CursorPos()
	relX := E.GetXRight() - posX 
	relY := E.GetYBot() - posY

	E.SetRelativeX(relX)
	E.SetRelativeY(relY)
	return true
}

// fix cursor in object
func isExpandable(E elements.Expandable) bool {
	posX, posY := CursorPos()

	if posX > E.GetXRight()-7 && posY > E.GetYBot()-7 && posX < E.GetXRight()&& posY < E.GetYBot(){ 
		return true
	}
	return false
}