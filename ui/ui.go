package ui

import (
	g "github.com/AllenDang/giu"
)

// This is an example from AllenDang/giu

type Handler interface {
	handleUI() 
}

func WindowManager() {

	w := g.NewMasterWindow("GraphicalGo", 1000, 800, 0)
	w.Run(handleUI)
}

