package eval

type Class struct {
	FunctionList map[int]Function
	Variables    map[int]Variable
}
