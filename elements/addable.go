package elements

type Addable interface {
	Adding()
	ReverseAdding()

	GetXLeft() int
	GetYTop() int
	
	GetXRight() int
	GetYBot() int

	GetLevel() int
	GetAddingState() bool 
}