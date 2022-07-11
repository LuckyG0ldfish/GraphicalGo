package ui

import (
	g "github.com/AllenDang/giu"
	"github.com/LuckyG0ldfish/GraphicalGo/elements"
	"github.com/LuckyG0ldfish/GraphicalGo/elements/subelements"
)

// This is an example from AllenDang/giu

type Handler interface {
	handleUI()
}

func WindowManager() {
	Dragables = make([]elements.Dragable, 3)
	Dragables[0] = subelements.CreateObject("Test :Object", 100 ) 
	Dragables[1] = subelements.CreateObject("Test :Object2", 300 ) 
	Dragables[2] = subelements.CreateObject("Test :Object3", 500 ) 

	w := g.NewMasterWindow("GraphicalGo", 1000, 800, 0)
	w.Run(handleUI)
}
