package objects

func (obj Object) ObjectsToString() (objects string) {
	objects = "type " + obj.name + " struct {\n"
	for _, s := range obj.variables {
		objects = objects + s.VariablesToString() + "\n"
	}
	objects = objects + "}"
	return objects
}