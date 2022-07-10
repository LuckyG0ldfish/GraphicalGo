package elements

// "image"
// "image/color"

// sub "github.com/LuckyG0ldfish/GraphicalGo/elements/subelements"

// "github.com/AllenDang/imgui-go"

type Dragable interface {
	GetID() int

	GetXLeft() int
	GetYTop() int


	GetRelativeX() int 
	GetRelativeY() int

	SetXLeft(int)
	SetYTop(int)


	SetRelativeX(int) 
	SetRelativeY(int) 
}
