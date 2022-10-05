package subelements

import "github.com/LuckyG0ldfish/GraphicalGo/elements"

const VariableType = 4

type Variable struct {
	Name   string
	Typ    string
	parent elements.Element
}

func CreateVariables() (variable string) {

	return variable
}

func (va *Variable) removeVariable(e []*Variable) []*Variable {
	// empty or last removing 
	if len(e) < 2 {
		return make([]*Variable, 0)
	}

	ret := make([]*Variable, len(e)-1)
	tempI := 0 
	for _, v := range e {
		if v != va {
			ret[tempI] = v
			tempI ++ 
		}
	}

	return ret
}