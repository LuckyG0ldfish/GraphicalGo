package conversion

import (
	"os"

	sub "github.com/LuckyG0ldfish/GraphicalGo/elements/subelements"
)

func VariablesToString(vari *sub.Variable) (variables string) {
	variables = vari.Name + " " + vari.Typ
	return
}

func ObjectsToString(obj *sub.Object) (objects string) {
	objects = "type " + obj.Name + " struct {\n"
	for _, s := range obj.Variables {
		objects = objects + VariablesToString(s) + "\n"
	}
	objects = objects + "}"
	return objects
}

func FilesToString(file sub.File, packageName string, path string) {
	files := "package " + packageName + "\n \n"

	files = files + file.Imports + "\n \n"

	for _, s := range file.Objects {
		files = files + ObjectsToString(s) + "\n \n"
	}

	for _, s := range file.Functions {
		files = files + s + "\n \n"
	}

	f, err := os.Create(path + "/" + file.Name + ".go")

	if err != nil {
		return 
	}

	// f.Close()

	_, err = f.WriteString(files)
	
	if err != nil {
		return 
	}
}

func FoldersToString(folder sub.Folder) (folders string) {
	
	return folders
}