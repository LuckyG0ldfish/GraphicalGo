package elements

type Addable interface {
	Adding(Dragable)

	GetName() string

	GetXLeft() int
	GetYTop() int
	
	GetXRight() int
	GetYBot() int

	GetLevel() int
	GetAddingState() bool 
	SetAddingState(bool)  
}