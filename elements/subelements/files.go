package subelements

import (
	"image"
	"image/color"

	g "github.com/AllenDang/giu"

	"github.com/LuckyG0ldfish/GraphicalGo/elements"
	
) 

type File struct {
	Name      string
	level     int
	Imports   string
	Objects   []*Object
	Functions []string
	id        int

	xLeft int
	yTop  int

	xRight int
	yBot   int

	xRelative int
	yRelative int
}

func CreateFiles() (files string) {

	return files
}

func (fil *File) GetXLeft() int {
	return fil.xLeft
}

func (fil *File) GetXRight() int {
	return fil.xRight
}

func (fil *File) GetYTop() int {
	return fil.yTop
}

func (fil *File) GetYBot() int {
	return fil.yBot
}

func (fil *File) SetXLeft(x int) {
	fil.xLeft = x
}

func (fil *File) SetXRight(x int) {
	fil.xRight = x
}

func (fil *File) SetYTop(y int) {
	fil.yTop = y
}

func (fil *File) SetYBot(y int) {
	fil.yBot = y
}

func (fil *File) GetID() int {
	return fil.id
}

func (fil *File) GetType() int {
	return 2 // Type of File
}

func (fil *File) GetLevel() int {
	return fil.level
}

func (fil *File) GetName() string {
	return fil.Name
}

func (fil *File) GetRelativeX() int {
	return fil.xRelative
}

func (fil *File) GetRelativeY() int {
	return fil.yRelative
}

func (fil *File) SetRelativeX(i int) {
	fil.xRelative = i
}

func (fil *File) SetRelativeY(i int) {
	fil.yRelative = i
}

func (fil *File) Draw(c *g.Canvas) {
	pos := g.GetCursorScreenPos()

	mix := fil.GetXLeft()
	miy := fil.GetYTop()

	max := fil.GetXRight()
	may := fil.GetYBot()

	c.AddRectFilled(pos.Add(image.Pt(mix, miy)), pos.Add(image.Pt(mix+110, miy+20)), color.White, 0, 5)
	c.AddRectFilled(pos.Add(image.Pt(mix, miy+20)), pos.Add(image.Pt(max, may)), color.White, 0, 5)
	c.AddLine(pos.Add(image.Pt(mix+3, miy+20)), pos.Add(image.Pt(mix+107, miy+20)), color.Black, 1)
	c.AddText(pos.Add(image.Pt(mix+3, miy+3)), color.Black, fil.GetName())

	// expand
	c.AddRectFilled(pos.Add(image.Pt(max-6, may-6)), pos.Add(image.Pt(max-1, may-1)), color.Black, 0, 5)
}

func (fil *File) GetSubelements() []elements.Drawable {
	r := make([]elements.Drawable, 0)
	for _, v := range fil.Objects {
		r = append(r, v)
	}
	// add Variable
	// for _, v := range fil.files {
	// 	r = append(r, v)
	// }
	return r
}
