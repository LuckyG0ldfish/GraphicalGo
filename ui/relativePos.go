package ui

import (
	ele "github.com/LuckyG0ldfish/GraphicalGo/elements"
)

func UpdateDragPos(D ele.Dragable) {
	relX, relY, working := getRelativePos(D)
	if !working {
		return
	}
	for {
		posX, posY := CursorPos()
		updateRelativePos(D, relX, relY, posX, posY)
		if !LeftPressed() {
			return 
		}
	}
}

func getRelativePos(D ele.Dragable) (relX int, relY int, working bool) {
	if !isDragable(D) {
		return 0, 0, false
	}
	posX, posY := CursorPos()
	relX = posX - D.GetXLeft()
	relY = posY - D.GetYTop()
	return relX, relY, true
}

func updateRelativePos(D ele.Dragable, relX int, relY int, curX int, curY int) {
	D.SetXLeft(curX - relX)
	D.SetYTop(curY - relY)
}

func isDragable(D ele.Dragable) bool {
	posX, posY := CursorPos()

	if posX > D.GetXLeft() && posX < D.GetXDist() && posY > D.GetYTop() && posY < D.GetYDist() {
		return true
	}
	return false
}
