package main

import (
	"fmt"

	"github.com/AllenDang/giu"

	"github.com/LuckyG0ldfish/GraphicalGo/ui"
)

func main() {

	wnd := giu.NewMasterWindow("Splitter", 800, 600, 0)
	ui.Build(wnd)
	test := "TestName"
	createProject(test)

}

func createProject(name string) {
	fmt.Println(name)


}





