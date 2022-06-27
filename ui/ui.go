package ui

import (
	"fmt"

	g "github.com/AllenDang/giu"
)



func WindowManager() {
	items = make([]string, 100)
	for i := range items {
		items[i] = fmt.Sprintf("Item %d", i)
	}

	w := g.NewMasterWindow("GraphicalGo", 1000, 800, 0)
	w.Run(gloop)
}