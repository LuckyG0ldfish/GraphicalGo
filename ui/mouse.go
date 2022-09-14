package ui

import (
	g "github.com/AllenDang/giu"
	// "github.com/hajimehoshi/ebiten"
	// imgui "github.com/AllenDang/imgui-go"
)

var offset int = 7 

func CursorPos() (x int, y int) {	
	// x, y = ebiten.CursorPosition()
	point := g.GetMousePos()
	x = point.X - offset
	y = point.Y - offset
	return 
}

func LeftPressed() bool {
	// O is left button according to imgui doc
	return g.IsMouseClicked(0)
}

func LeftReleased() bool {
	// O is left button according to imgui doc
	return g.IsMouseReleased(0)
}

func RightPressed() bool {
	// 1 is right button according to imgui doc
	return g.IsMouseClicked(1)
}

func RightReleased() bool {
	// 1 is right button according to imgui doc
	return g.IsMouseReleased(1)
}

func MiddlePressed() bool {
	return g.IsMouseClicked(2)
}