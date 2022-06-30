package subelements

type Object struct {
	Name      string
	Variables []Variable

	xLeft 	int 
	xRight	int 
	yLeft	int 
	yRight	int

	id 		int
}

func CreateObject() (object string) {
	
	return object
}

func (ob *Object) GetXLeft() int{
	return ob.xLeft
}

func (ob *Object) GetXRight() int{
	return ob.xRight
}

func (ob *Object) GetYLeft() int{
	return ob.yLeft
}

func (ob *Object) GetYRight() int{
	return ob.yRight
}

func (ob *Object) SetXLeft(x int) {
	ob.xLeft = x
}

func (ob *Object) SetXRight(x int) {
	ob.xRight = x
}

func (ob *Object) SetYLeft(y int) {
	ob.yLeft = y
}

func (ob *Object) SetYRight(y int) {
	ob.yRight = y
}

func (ob *Object) GetID() int{
	return ob.id
}
