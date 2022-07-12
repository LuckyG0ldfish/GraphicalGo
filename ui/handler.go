package ui

import "time"

var (
	// testVar g.CustomWidget
	// dragInProgress bool
	// Dragable       elements.Dragable
	// Dragables []elements.Dragable
)

func handleGeneralMousePosition() {
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
	curX, _ := CursorPos()
	return (curX > Project.Can.XLeft && curX < Project.Can.XRight)
}

func CursorOnOverView() bool {
	curX, _ := CursorPos()
	return (curX < Project.Can.XLeft)
}

func CursorOnObjectSelect() bool {
	curX, _ := CursorPos()
	return (curX > Project.Can.XRight)
}

func handleClickToCanvas() {
	if LeftPressed() {
		for _, v := range Project.Can.Dragables {
			if isDragable(v) {
				UpdateDragPos(v)
			}
		}
	}
}

func handleClickToOverView() {

	return
}

func handleClickToObjectSelect() {

	return
}
