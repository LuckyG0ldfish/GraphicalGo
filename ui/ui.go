package ui

import (
	"time"

	g "github.com/AllenDang/giu"
	"github.com/LuckyG0ldfish/GraphicalGo/elements/subelements"
)

type Handler interface {
	handleUI()
}

var Project subelements.Project

func WindowManager() {
	Project = *subelements.NewProject("GraphicalGo", 1500, 800)
	w := g.NewMasterWindow(Project.Name, Project.Win.XWidth, Project.Win.YHeight, 0)
	go handleGeneralMousePosition()
	go updateFrameSize(w)
	w.Run(drawUI)
}

func updateFrameSize(w *g.MasterWindow) {
	for {
		wid, hei := w.GetSize()
		Project.Win.XWidth = wid
		Project.Win.YHeight = hei
		Project.Can.XRight = wid - 200 
		time.Sleep(10 * time.Millisecond)
	}
}