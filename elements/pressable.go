package elements

type Pressable interface {
	Press()
	GetPos() int 
	GetName() string
	
	GetXLeft() int
	GetYTop() int

	GetXRight() int
	GetYBot() int
}