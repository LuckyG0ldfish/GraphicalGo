package inputs

import (
	"time"

	"github.com/LuckyG0ldfish/GraphicalGo/elements"
	"github.com/LuckyG0ldfish/GraphicalGo/elements/subelements"
)

var (
	// testVar g.CustomWidget
	// dragInProgress bool
	// Dragable       elements.Dragable
	// Dragables []elements.Dragable
)

func HandleGeneralMousePosition() {
	for {
		for CursorOnOverView() {
			handleClickToOverView()
		}
		for CursorOnObjectSelect() {
			handleClickToObjectSelect()
		}
		for CursorOnCanvas() {
			handleClickToCanvas()
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func CursorOnCanvas() bool {
	Project := subelements.GetPro()
	curX, _ := CursorPos()
	return (curX > Project.Can.XLeft && curX < Project.Can.XRight)
}

func CursorOnOverView() bool {
	Project := subelements.GetPro()
	curX, _ := CursorPos()
	return (curX < Project.Can.XLeft)
}

func CursorOnObjectSelect() bool {
	Project := subelements.GetPro()
	curX, _ := CursorPos()
	return (curX > Project.Can.XRight)
}

func handleClickToCanvas() {
	Project := subelements.GetPro()
	if LeftPressed() {
		for _, v := range Project.Can.Expandables {
			if isExpandable(v) {
				UpdateExpanding(v)
			}
		}
		var drag elements.Dragable
		for _, v := range Project.Can.Dragables {
			if isDragable(v) {
				if drag == nil || drag.GetLevel() < v.GetLevel(){
					drag = v 
				}
			}
		}
		UpdateDragPos(drag)
	}
}

func handleClickToOverView() {

	return
}

func handleClickToObjectSelect() {
	Project := subelements.GetPro()
	if LeftPressed() {
		for _, v := range Project.Obj.Pressables {
			if isPressable(v) {
				for !LeftReleased(){}
				v.Press()
			}
		}
	}
}
