package ui

import (
	"time"

	g "github.com/AllenDang/giu"
	"github.com/LuckyG0ldfish/GraphicalGo/subelements"
	"github.com/LuckyG0ldfish/GraphicalGo/context"
	"github.com/LuckyG0ldfish/GraphicalGo/ui/inputs"
	"github.com/LuckyG0ldfish/GraphicalGo/ui/outputs"
)

type Handler interface {
	handleUI()
}

func WindowManager() {
	Project := context.NewProject("GraphicalGo", 1500, 800)
	w := g.NewMasterWindow(Project.Name, Project.Win.XWidth, Project.Win.YHeight, 0)
	trialSetup() // testing
	go inputs.HandleGeneralMousePosition()
	go updateFrameSize(w)
	w.Run(outputs.DrawUI)
}

func updateFrameSize(w *g.MasterWindow) {
	for {
		wid, hei := w.GetSize()
		Project := context.GetPro()
		
		if wid != Project.Win.XWidth || hei != Project.Win.YHeight {
			Project.Win.XWidth = wid
			Project.Win.YHeight = hei
			Project.Can.XRight = wid - 200 
			Project.Obj.XLeft = wid - 200
			Project.Obj.XRight = wid
			Project.Obj.AdjustButtonPositions()
		}
		time.Sleep(100 * time.Millisecond)
	}
}

// only for testing
func trialSetup() {
	// pro.Can.Dragables = make([]elements.Element, 0)
	// pro.Can.Expandables = make([]elements.Expandable, 0)
	// pro.Obj.Pressables = make([]elements.Pressable, 0)

	subelements.CreateObject("Test1 :Object", 200)
	// CreateObject("Test2 :Object", 300)
	// CreateObject("Test3 :Object", 500)
	subelements.CreateFolders("Test1 :Package", 700)
	subelements.CreateFiles("Test : File", 400)

	subelements.CreateButton("Create Package", 0, func() {
		subelements.CreateFolders("addedPackage", 700)
	})
	subelements.CreateButton("Create File", 1, func() {
		subelements.CreateFiles("addedFile", 400)
	})
	subelements.CreateButton("Create Object", 2, func() {
		subelements.CreateObject("addedObject", 200)
	})
}