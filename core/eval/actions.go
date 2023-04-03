package eval

// Action struct is for function or class calls

type Action struct {
	RequestedID       string
	RequiredVariables map[int]Variable
}
