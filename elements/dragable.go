package elements

import (
	g "github.com/AllenDang/giu"
)

// "image"
// "image/color"

// sub "github.com/LuckyG0ldfish/GraphicalGo/elements/subelements"

// "github.com/AllenDang/imgui-go"

type Element interface {
	GetID() int // 1 = Folder, 2 = File, 3 = Object
	GetType() int
	GetName() string
	GetLevel() int

	Draw(*g.Canvas)
	GetSubelements() []Element
	GetParent() Element

	GetXLeft() int
	GetYTop() int
	SetXLeft(int)
	SetYTop(int)

	GetXRight() int
	GetYBot() int
	SetXRight(int)
	SetYBot(int)

	GetRelativeX() int
	GetRelativeY() int
	SetRelativeX(int)
	SetRelativeY(int)
}
