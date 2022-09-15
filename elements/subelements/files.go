package subelements


type File struct{
	Name string
	level int
	Imports string
	Objects []Object
	Functions []string

	xLeft 	int 
	xRight	int 
	yLeft	int 
	yRight	int

	id 		int
}
func CreateFiles() (files string) {

	return files
}

func (fil *File) GetXLeft() int{
	return fil.xLeft
}

func (fil *File) GetXRight() int{
	return fil.xRight
}

func (fil *File) GetYLeft() int{
	return fil.yLeft
}

func (fil *File) GetYRight() int{
	return fil.yRight
}

func (fil *File) SetXLeft(x int) {
	fil.xLeft = x
}

func (fil *File) SetXRight(x int) {
	fil.xRight = x
}

func (fil *File) SetYLeft(y int) {
	fil.yLeft = y
}

func (fil *File) SetYRight(y int) {
	fil.yRight = y
}

func (fil *File) GetID() int{
	return fil.id
}

func (fil *File) GetType() int{
	return 2 	// Type of File 
}

func (fil *File) GetLevel() int{
	return fil.level 
}