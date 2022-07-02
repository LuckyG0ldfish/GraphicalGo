package ui

import (
	imgui "github.com/AllenDang/imgui-go"
)

func CursorPos() (x int, y int) {
	x = int(imgui.CursorPosX())
	y = int(imgui.CursorPosY())
	return 
}

func LeftPressed() bool {
	// O is left button according to imgui doc
	return imgui.IsMouseClicked(0)
}
