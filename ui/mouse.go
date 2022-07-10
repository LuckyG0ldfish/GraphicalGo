package ui

import (
	imgui "github.com/AllenDang/imgui-go"
)

func CursorPos() (x int, y int) {
	vec := imgui.CursorPos()
	
	x = int(vec.X)
	y = int(vec.Y)
	return 
}

func LeftPressed() bool {
	// O is left button according to imgui doc
	return imgui.IsMouseClicked(0)
}

func LeftReleased() bool {
	// O is left button according to imgui doc
	return imgui.IsMouseReleased(0)
}

func RightPressed() bool {
	// 1 is right button according to imgui doc
	return imgui.IsMouseClicked(1)
}

func RightReleased() bool {
	// 1 is right button according to imgui doc
	return imgui.IsMouseReleased(1)
}
