package variables

func (vari Variable) VariablesToString() (variables string) {
	variables = vari.name + " " + vari.typ
	return
}
