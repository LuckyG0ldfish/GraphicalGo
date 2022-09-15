package ui

import (
	"time"

	g "github.com/AllenDang/giu"
	"github.com/LuckyG0ldfish/GraphicalGo/elements/subelements"
	"github.com/LuckyG0ldfish/GraphicalGo/ui/inputs"
	"github.com/LuckyG0ldfish/GraphicalGo/ui/outputs"
)

type Handler interface {
	handleUI()
}

func WindowManager() {
	Project := subelements.NewProject("GraphicalGo", 1500, 800)
	w := g.NewMasterWindow(Project.Name, Project.Win.XWidth, Project.Win.YHeight, 0)
	go inputs.HandleGeneralMousePosition()
	go updateFrameSize(w)
	w.Run(outputs.DrawUI)
}

func updateFrameSize(w *g.MasterWindow) {
	for {
		wid, hei := w.GetSize()
		Project := subelements.GetPro()
		Project.Win.XWidth = wid
		Project.Win.YHeight = hei
		Project.Can.XRight = wid - 200 
		time.Sleep(10 * time.Millisecond)
	}
}