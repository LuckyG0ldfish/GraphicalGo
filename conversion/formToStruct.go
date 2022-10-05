package conversion

import (
	"os"

	"github.com/LuckyG0ldfish/GraphicalGo/context"
)

func FormToStruct() (bool, context.Project) {
	var pro context.Project

	s, e := os.ReadFile("project.txt")
	if e == nil {
		return false, pro 
	}
	if len(s) == 0 {
		return false, pro 
	}

	

	return true, pro
}

