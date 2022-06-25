package conversion

import (
	"os"

	"github.com/LuckyG0ldfish/GraphicalGo/elements/files"
	"github.com/LuckyG0ldfish/GraphicalGo/elements/folders"
	"github.com/LuckyG0ldfish/GraphicalGo/elements/objects"
	"github.com/LuckyG0ldfish/GraphicalGo/elements/variables"
)

func VariablesToString(vari variables.Variable) (variables string) {
	variables = vari.Name + " " + vari.Typ
	return
}

func ObjectsToString(obj objects.Object) (objects string) {
	objects = "type " + obj.Name + " struct {\n"
	for _, s := range obj.Variables {
		objects = objects + VariablesToString(s) + "\n"
	}
	objects = objects + "}"
	return objects
}

func FilesToString(file files.File, packageName string, path string) {
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

func FoldersToString(folder folders.Folder) (folders string) {
	
	return folders
}