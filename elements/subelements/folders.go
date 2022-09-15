package subelements

type Folder struct {
	name string 
	level int
	path string
	files []File

	xLeft 	int 
	yTop	int 

	xRight 	int 
	yBot	int 

	xRelative int 
	yRelative int 

	id 		int
}

func CreateFolders(name string, x int) (*Folder) {
	var folder Folder
	folder.name = name
	folder.level = 1
	
	folder.xLeft = x
	folder.yTop = 100
	
	folder.xRight = x+300
	folder.yBot = 300

	folder.xRelative = 0
	folder.yRelative = 0 

	pro.Can.Dragables = append(pro.Can.Dragables, &folder) 
	pro.Can.Expandables = append(pro.Can.Expandables, &folder)
	return &folder
}

func (fol *Folder) Expand() {}

func (fol *Folder) GetXLeft() int{
	return fol.xLeft
}

func (fol *Folder) GetXRight() int{
	return fol.xRight
}

func (fol *Folder) GetYTop() int{
	return fol.yTop
}

func (fol *Folder) GetYBot() int{
	return fol.yBot
}

func (fol *Folder) SetXLeft(x int) {
	fol.xLeft = x
}

func (fol *Folder) SetXRight(x int) {
	fol.xRight = x
}

func (fol *Folder) SetYTop(y int) {
	fol.yTop = y
}

func (fol *Folder) SetYBot(y int) {
	fol.yBot = y
}

func (fol *Folder) GetRelativeX() int{
	return fol.xRelative
}

func (fol *Folder) GetRelativeY() int{
	return fol.yRelative
}

func (fol *Folder) SetRelativeX(i int) {
	fol.xRelative = i 

}

func (fol *Folder) SetRelativeY(i int) {
	fol.yRelative = i 
}

func (fol *Folder) GetName() string{
	return fol.name
}

func (fol *Folder) GetID() int{
	return fol.id
}

func (fol *Folder) GetType() int{
	return 1	// Type for Folder
}

func (fol *Folder) GetLevel() int{
	return fol.level
}