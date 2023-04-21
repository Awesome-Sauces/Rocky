package eval

// Action struct is for function or class calls

type Action struct {
	Name              string
	RequestedID       string
	RequiredVariables map[int]*Variable
}

// Construct an Action Struct, mainly to be used in Classes and Functions but may be used normally to store tasks to run later

/*

We want to structure the Action to be convertable
To a JSON format easily. First we need to decide
what can be done inside an Action.

So maybe we write a universal eval, that takes in a map of tokens and runs the map of tokens

*/
