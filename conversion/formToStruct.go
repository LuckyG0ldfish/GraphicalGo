package conversion

import (
	"fmt"
	"os"
	"strconv"

	"github.com/LuckyG0ldfish/GraphicalGo/context"
)

type element struct {
	L int 
	I int 
	T int 
	N string 
	X int
	Y int
}

func FormToStruct() (bool, context.Project) {
	var pro context.Project

	s, e := os.ReadFile("GraphicalGo.txt")
	if e != nil {
		fmt.Println("E1")
		return false, pro 
	}
	if len(s) == 0 {
		fmt.Println("E2")
		return false, pro 
	}
	fmt.Println("parsing")
	// ele := parseArguments(s, true)
	return true, pro
}

// 10 = new line 
// 32 = space 
func parseArguments(bytes []byte, integer bool) []element {
	temp := make([]byte, 0)
	last := false  
	var ele element
	eles := make([]element, 0)

	for _, b := range bytes {
		if b == 32 {
		} else if last && b == 10 {
			eles = append(eles, ele)
			last = false
		} else if b == 10 {
			parseArgument(temp, &ele)
			temp = make([]byte, 0)
			last = true
		} else {
			temp = append(temp, b)
		}
	}

	return eles 
}

func parseArgument(bytes []byte, e *element) {
	if len(bytes) == 0  {
		return 
	}

	flag := string(bytes[0])
	value := ""

	for i, b := range bytes {
		
		if i > 1 {
			value = value + string(b)
		} 
	}

	switch flag {
	case "L":
		i, er := strconv.Atoi(value)
		if er != nil {
			e.L = i 
		}
	case "I":
		i, er := strconv.Atoi(value)
		if er != nil {
			e.I = i 
		}
	case "T":
		i, er := strconv.Atoi(value)
		if er != nil {
			e.T = i 
		}
	case "N":
		e.N = value
	case "X":
		i, er := strconv.Atoi(value)
		if er != nil {
			e.X = i 
		}
	case "Y":
		i, er := strconv.Atoi(value)
		if er != nil {
			e.Y = i 
		}
	}
}
