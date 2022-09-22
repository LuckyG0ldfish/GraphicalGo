package elements

type Addable interface {
	Adding(Element)

	// GetName() string
	GetID() int 

	GetXLeft() int
	GetYTop() int

	GetXRight() int
	GetYBot() int

	GetLevel() int
	GetAddingState() bool
	SetAddingState(bool)
}
