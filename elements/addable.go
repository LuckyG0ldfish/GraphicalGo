package elements

type Addable interface {
	Adding(Element)
	Removing(Element)
	GetType() int
	
	SetAsParent(Element)
	// GetName() string
	GetID() int 
	// GetSubelements() []Addable
	GetXLeft() int
	GetYTop() int

	GetXRight() int
	GetYBot() int

	GetLevel() int
	GetAddingState() bool
	SetAddingState(bool)
}
