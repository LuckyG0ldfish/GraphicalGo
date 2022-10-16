package conversion

import (
	"fmt"
	"os"
	"strconv"

	"github.com/LuckyG0ldfish/GraphicalGo/context"
	"github.com/LuckyG0ldfish/GraphicalGo/elements"
	"github.com/LuckyG0ldfish/GraphicalGo/subelements"
)

type structElement struct {
	L int 
	I int 
	T int 
	N string 
	X int
	Y int
}

func FormToStruct() bool {
	pro := context.GetPro()
	if pro.SaveLoaded {
		return false 
	}
	
	s, e := os.ReadFile("GraphicalGo.txt")
	if e != nil {
		fmt.Println("File Corrupted / Not existing")
		return false
	}
	if len(s) == 0 {
		fmt.Println("Empty Save")
		return false
	}
	check, id := addElementsToProject(s)
	if check {
		pro.SetNextID(id)
		pro.SaveLoaded = true 
		return true 
	}
	return false 
}

func addElementsToProject(s []byte) (bool, int) {
	check, eles := parseArguments(s, true)
	if !check {
		return false, 0
	}
	tempID := 0 
	var prev elements.Element
	for _, v := range eles {	
		if v.I > tempID {
			tempID = v.I
		}	
		ele := elementToSubelement(v)
		if v.L != 1 && prev != nil {
			if prev.GetLevel() < v.L{
				prev.Adding(ele)
			} else if prev.GetLevel() > v.L {
				temp := prev.GetParent()
				for temp.GetLevel() >= v.L {
					temp = temp.GetParent()
					if temp == nil {
						return false, 0
					}	
				}
				temp.Adding(ele)
			} else {
				prev.GetParent().Adding(ele)
			}
		}
		prev = ele
	}
	return true, tempID + 1
}

func elementToSubelement(e structElement) elements.Element{
	var ele elements.Element

	
	switch e.T {
	case subelements.FolderType:
		ele = subelements.CreateAndInitFolders(e.N, e.X, e.Y, e.I, e.L)
	case subelements.FileType: 
		ele = subelements.CreateAndInitFiles(e.N, e.X, e.Y, e.I, e.L)
	case subelements.ObjectType:
		ele = subelements.CreateAndInitObject(e.N, e.X, e.Y, e.I, e.L)
	// case subelements.VariableType: 	TODO
	}
	
	return ele
}

// 10 = new line 
// 32 = space 
func parseArguments(bytes []byte, integer bool) (bool, []structElement) {
	temp := make([]byte, 0)
	last := false  
	var ele structElement
	eles := make([]structElement, 0)

	for _, b := range bytes {
		if b == 32 {
		} else if last && b == 10 {
			last = false
			if checkStructEle(ele) {
				eles = append(eles, ele)
			} else {
				fmt.Println("unvollständig")
			}
			ele = createStructureElement()
		} else if b == 10 {
			if !parseArgument(temp, &ele) {
				return false, eles 
			}
			temp = make([]byte, 0)
			last = true
		} else {
			temp = append(temp, b)
			last = false 
		}
	}
	eles = append(eles, ele)
	return true, eles 
}

func createStructureElement() structElement {
	var ele structElement
	ele.I = 0 
	ele.L = 0 
	ele.N = "" 
	ele.T = 0
	ele.X = 0
	ele.Y = 0  
	return ele 
}

func checkStructEle(e structElement) bool {
	if e.I == 0 || e.L == 0 || e.N == "" || e.T == 0 || e.X == 0 || e.Y == 0 {
		return false 
	}
	return true 
}

func parseArgument(bytes []byte, e *structElement) bool {
	if len(bytes) == 0  {
		return true  
	}

	flag := string(bytes[0])
	value := ""

	for i, b := range bytes {
		if i > 0 {
			value = value + string(b)
		} 
	}
	switch flag {
	case "L":
		i, er := strconv.Atoi(value)
		if er == nil {
			e.L = i 
		} else {
			return false 
		}
	case "I":
		i, er := strconv.Atoi(value)
		if er == nil {
			e.I = i 
		} else {
			return false 
		}
	case "T":
		i, er := strconv.Atoi(value)
		if er == nil {
			e.T = i 
		} else {
			return false 
		}
	case "N":
		e.N = value
	case "X":
		i, er := strconv.Atoi(value)
		if er == nil {
			e.X = i 
		} else {
			return false 
		}
	case "Y":
		i, er := strconv.Atoi(value)
		if er == nil {
			e.Y = i 
		} else {
			return false 
		}
	}
	return true 
}


