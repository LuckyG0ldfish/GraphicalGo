package subelements

type Object struct {
	Name      string
	Variables []Variable

	xLeft 	int 
	yTop	int 

	xRight 	int 
	yBot	int 

	xRelative int 
	yRelative int 

	id 		int
}

func CreateObject(name string) (*Object) {
	var object Object
	object.Name = name
	object.xLeft = 100
	object.yTop = 100

	object.xRight = 200
	object.yBot = 200

	object.xRelative = 0
	object.yRelative = 0 
	return &object
}

func (ob *Object) GetName() string{
	return ob.Name
}

func (ob *Object) GetRelativeX() int{
	return ob.xRelative
}

func (ob *Object) GetRelativeY() int{
	return ob.yRelative
}

func (ob *Object) SetRelativeX(i int) {
	ob.xRelative = i 

}

func (ob *Object) SetRelativeY(i int) {
	ob.yRelative = i 
}

func (ob *Object) GetXLeft() int{
	return ob.xLeft
}

func (ob *Object) GetYTop() int{
	return ob.yTop
}

func (ob *Object) SetXLeft(x int) {
	ob.xLeft = x
}

func (ob *Object) SetYTop(y int) {
	ob.yTop = y
}

func (ob *Object) GetXRight() int{
	return ob.xRight
}

func (ob *Object) GetYBot() int{
	return ob.yBot
}

func (ob *Object) SetXRight(x int) {
	ob.xRight = x
}

func (ob *Object) SetYBot(y int) {
	ob.yBot = y
}

func (ob Object) GetID() int{
	return ob.id
}
