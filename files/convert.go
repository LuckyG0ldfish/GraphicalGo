package files

import (
	"os"
)

func (file File) FilesToString(packageName string, path string) {
	files := "package " + packageName + "\n \n"

	files = files + file.imports + "\n \n"

	for _, s := range file.objects {
		files = files + s.ObjectsToString() + "\n \n"
	}

	for _, s := range file.functions {
		files = files + s + "\n \n"
	}

	os.Create()
}