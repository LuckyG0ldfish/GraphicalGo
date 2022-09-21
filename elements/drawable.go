package elements

import (
	g "github.com/AllenDang/giu"
)

type Drawable interface {
	GetType() int
	Draw(*g.Canvas)
	GetSubelements() []Drawable
	GetParent() Drawable

	GetXLeft() int
	GetYTop() int
	
	GetXRight() int
	GetYBot() int

	SetXLeft(int)
	SetYTop(int)

	SetXRight(int)
	SetYBot(int)
}
