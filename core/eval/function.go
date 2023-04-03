package eval

type Function struct {
	Name         string
	ActionList   map[int]Action
	VariableList map[int]Variable
}
