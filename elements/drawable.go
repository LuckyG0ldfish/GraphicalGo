package elements

import (
	g "github.com/AllenDang/giu"
)

type Drawable interface {
	GetType() int
	Draw(*g.Canvas)
	GetSubelements() []Drawable
}
