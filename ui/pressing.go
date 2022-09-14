package ui

import "github.com/LuckyG0ldfish/GraphicalGo/elements"

// fix cursor in object
func isPressable(D elements.Pressable) bool {
	posX, posY := CursorPos()

	if posX > D.GetXLeft() && posY > D.GetYTop() && posX < D.GetXRight()&& posY < D.GetYBot(){ 
		return true
	}
	return false
}