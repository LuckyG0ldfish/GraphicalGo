package elements

type Expandable interface {
	Expand()

	// GetXLeft() int
	// GetYTop() int
	// SetXLeft(int)
	// SetYTop(int)

	GetXRight() int
	GetYBot() int
	SetXRight(int)
	SetYBot(int)

	GetRelativeX() int 
	GetRelativeY() int
	SetRelativeX(int) 
	SetRelativeY(int) 
}