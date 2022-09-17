package inputs

import (
	"github.com/LuckyG0ldfish/GraphicalGo/elements"
	"github.com/LuckyG0ldfish/GraphicalGo/elements/subelements"
)

func checkOnAdding(e elements.Dragable) (bool, elements.Addable) {
	pro := subelements.GetPro()
	typ := e.GetType()
	var temp elements.Addable
	check := false 
	if typ == 1 || typ == 2 { // Folder or file 
		for _, v := range pro.Can.Add {
			if CanAdd(v.GetXLeft(), v.GetXRight(), v.GetYTop(), v.GetYBot(), e.GetXLeft(), e.GetXRight(), e.GetYTop(), e.GetYBot()) {
				if temp == nil || temp.GetLevel() < v.GetLevel() {
					check = true 
					temp = v
				}
			}
		}
	}
	return check , temp 
}

// Corner in field
func CanAdd(xLeftAddable, xRightAddable, yTopAddable, yBotAddable, xLeftToAdd, xRightToAdd, yTopToAdd, yBotToAdd int) bool {
	if inField(xLeftAddable, xRightAddable, yTopAddable, yBotAddable, xLeftToAdd, yTopToAdd) || 
		inField(xLeftAddable, xRightAddable, yTopAddable, yBotAddable, xRightToAdd, yBotToAdd) || 
		inField(xLeftAddable, xRightAddable, yTopAddable, yBotAddable, xRightToAdd, yTopToAdd) ||
		inField(xLeftAddable, xRightAddable, yTopAddable, yBotAddable, xLeftToAdd, yBotToAdd) {
		return true 
	}
	return false
}

func inField(xLeftAddable, xRightAddable, yTopAddable, yBotAddable, x, y int) bool {
	if xLeftAddable < x && xRightAddable > x && yTopAddable < y && yBotAddable > y {
		return true 
	}
	return false 
}