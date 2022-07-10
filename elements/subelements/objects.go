package subelements

type Object struct {
	Name      string
	Variables []Variable

	xLeft 	int 
	yTop	int 


	xRelative int 
	yRelative int 

	id 		int
}

func CreateObject(name string) (*Object) {
	var object Object
	object.xLeft = 100
	object.yTop = 100
	object.Name = name
	object.xRelative = 100
	object.yRelative = 100 
	return &object
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

func (ob Object) GetID() int{
	return ob.id
}
