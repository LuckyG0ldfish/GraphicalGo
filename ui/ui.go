package ui

import (
	g "github.com/AllenDang/giu"
	"github.com/LuckyG0ldfish/GraphicalGo/elements/subelements"
)

// This is an example from AllenDang/giu

type Handler interface {
	handleUI()
}

func WindowManager() {
	Dragable = subelements.CreateObject("Test :Object")

	w := g.NewMasterWindow("GraphicalGo", 1000, 800, 0)
	w.Run(handleUI)
}
