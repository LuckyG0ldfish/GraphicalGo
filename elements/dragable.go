package elements

// "image"
// "image/color"

// sub "github.com/LuckyG0ldfish/GraphicalGo/elements/subelements"

// "github.com/AllenDang/imgui-go"

type Dragable interface {
	GetID() int
	GetName() string

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
