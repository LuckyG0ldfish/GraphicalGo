package ui

import (
	g "github.com/AllenDang/giu"
	"github.com/LuckyG0ldfish/GraphicalGo/elements/subelements"
)

// This is an example from AllenDang/giu

type Handler interface {
	handleUI()
}

var Project subelements.Project

func WindowManager() {
	Project = *subelements.NewProject("GraphicalGo", 1000, 800)
	w := g.NewMasterWindow(Project.Name, Project.Win.XWidth, Project.Win.YHeight, 0)
	go handleGeneralMousePosition()
	w.Run(drawUI)
}
