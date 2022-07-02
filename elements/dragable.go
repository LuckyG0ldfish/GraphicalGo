package elements

// "image"
// "image/color"

// sub "github.com/LuckyG0ldfish/GraphicalGo/elements/subelements"

// "github.com/AllenDang/imgui-go"

type Dragable interface {
	GetID()

	GetXLeft() int
	GetXDist() int
	GetYTop() int
	GetYDist() int

	SetXLeft(int)
	SetXDist(int)
	SetYTop(int)
	SetYDist(int)
}
