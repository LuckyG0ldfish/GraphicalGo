package conversion

import (
	"os"
	"strconv"

	"github.com/LuckyG0ldfish/GraphicalGo/context"
	"github.com/LuckyG0ldfish/GraphicalGo/elements"
)


func StructsToForm() {
	pro := context.GetPro()
	output := ""
	for _, v := range pro.Level1 {
		output = output + recStructToForm(v)
	}
	
	f, err := os.Create(pro.Name + ".txt") // "conversion/" + 
	
	if err != nil {
		return 
	}

	// f.Close()

	f.WriteString(output)
}

func recStructToForm(e elements.Element) string {
	output := "\n"
	output = output + elementToString(e)
	for _, v := range e.GetSubelements() {
		output = output + recStructToForm(v)
	}
	return output
}

func elementToString(e elements.Element) string {
	output := ""
	output = output + "L " + strconv.Itoa(e.GetLevel()) + " \n"
	output = output + "I " + strconv.Itoa(e.GetID()) + " \n"
	output = output + "T " + strconv.Itoa(e.GetType()) + " \n"
	output = output + "N " + e.GetName() + " \n"
	output = output + "X " + strconv.Itoa(e.GetXLeft()) + " \n"
	output = output + "Y " + strconv.Itoa(e.GetYTop()) + " \n"
	return output
}