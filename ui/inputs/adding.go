package inputs

import (
	"github.com/LuckyG0ldfish/GraphicalGo/context"
	"github.com/LuckyG0ldfish/GraphicalGo/elements"
	"github.com/LuckyG0ldfish/GraphicalGo/subelements"
)

func reverseAllAddingStates() {
	pro := context.GetPro()
	for _, v := range pro.Can.Add {
		v.SetAddingState(false)
	}
}

func checkOnAdding(e elements.Element) (bool, elements.Addable) {
	pro := context.GetPro()
	// typ := e.GetType()
	var temp elements.Addable
	check := false
	id := 0
	if e.GetParent() != nil {
		id = e.GetParent().GetID()
	}
	for _, v := range pro.Can.Add {
		if v.GetID() == e.GetID() || id == v.GetID() {
			continue
		} else if canAddType(e.GetType(), v.GetType()) && canAdd(v.GetXLeft(), v.GetXRight(), v.GetYTop(), v.GetYBot(), e.GetXLeft(), e.GetXRight(), e.GetYTop(), e.GetYBot()) {
			if temp == nil || temp.GetLevel() < v.GetLevel() {
				check = true
				temp = v
			}
		}
	
	}
	return check, temp
}

// Corner in field
func canAdd(xLeftAddable, xRightAddable, yTopAddable, yBotAddable, xLeftToAdd, xRightToAdd, yTopToAdd, yBotToAdd int) bool {
	if inField(xLeftAddable, xRightAddable, yTopAddable, yBotAddable, xLeftToAdd, yTopToAdd) ||
		inField(xLeftAddable, xRightAddable, yTopAddable, yBotAddable, xRightToAdd, yBotToAdd) ||
		inField(xLeftAddable, xRightAddable, yTopAddable, yBotAddable, xRightToAdd, yTopToAdd) ||
		inField(xLeftAddable, xRightAddable, yTopAddable, yBotAddable, xLeftToAdd, yBotToAdd) {
		return true
	}
	return false
}

func canAddType(dragType, addType int) bool {
	switch dragType {
		case subelements.FolderType: 
			if addType == subelements.FolderType {
				return true 
			}
		case subelements.FileType: 
			if addType == subelements.FolderType {
				return true 
			}
		case subelements.ObjectType: 
			if addType == subelements.FileType {
				return true
			}
		case subelements.VariableType: 
			if addType == subelements.ObjectType || addType == subelements.FileType{
				return true 
			}
	} 
	return false 
}

func inField(xLeftAddable, xRightAddable, yTopAddable, yBotAddable, x, y int) bool {
	if xLeftAddable < x && xRightAddable > x && yTopAddable < y && yBotAddable > y {
		return true
	}
	return false
}
