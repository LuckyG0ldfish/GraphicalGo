package elements

type Pressable interface {
	Press()
	GetPos() int 
	GetName() string
	
	GetXLeft() int
	GetYTop() int
	SetXLeft(int)
	SetYTop(int)

	GetXRight() int
	GetYBot() int
	SetXRight(int)
	SetYBot(int)
}