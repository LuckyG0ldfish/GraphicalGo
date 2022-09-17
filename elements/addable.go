package elements

type Addable interface {
	Adding()

	GetName() string

	GetXLeft() int
	GetYTop() int
	
	GetXRight() int
	GetYBot() int

	GetLevel() int
	GetAddingState() bool 
	SetAddingState(bool)  
}