package subelements


type Folder struct {
	name string 
	path string
	files []File

	xLeft 	int 
	xRight	int 
	yLeft	int 
	yRight	int

	id 		int
}

func CreateFolders() (folders string) {

	return folders
}

func (fol *Folder) GetXLeft() int{
	return fol.xLeft
}

func (fol *Folder) GetXRight() int{
	return fol.xRight
}

func (fol *Folder) GetYLeft() int{
	return fol.yLeft
}

func (fol *Folder) GetYRight() int{
	return fol.yRight
}

func (fol *Folder) SetXLeft(x int) {
	fol.xLeft = x
}

func (fol *Folder) SetXRight(x int) {
	fol.xRight = x
}

func (fol *Folder) SetYLeft(y int) {
	fol.yLeft = y
}

func (fol *Folder) SetYRight(y int) {
	fol.yRight = y
}

func (fol *Folder) GetID() int{
	return fol.id
}