package eval

// Action struct is for function or class calls

type Action struct {
	Name              string
	RequestedID       string
	RequiredVariables map[int]Variable
}
