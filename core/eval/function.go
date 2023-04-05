package eval

import (
	"fmt"
	"strings"
)

// Function registrar
var functions = make(map[string]*Function)

// Register New Function
func RegisterFunction(function *Function) {
	functions[function.Name] = function
}

// Get Function
func GetFunction(name string) *Function {
	return functions[name]
}

// Register Default Rocky Function
func DefaultRegistery() {
	print := NewFunction("print",
		make(map[string]*Variable), make(map[string]*Action), NewArgs("input-STRING"),
		func() {
			fmt.Println("Hello World")
		})

	RegisterFunction(print)
}

type Function struct {
	Name         string
	Args         map[string]VTYPE
	VariableList map[string]*Variable
	ActionList   map[string]*Action
	Executable   func()
}

func (function *Function) Execute() {
	// To be made
	function.Executable()
}

// Structure the arg like this: name-VTYPE
func NewArgs(args ...string) map[string]VTYPE {
	ArgList := make(map[string]VTYPE)

	for _, arg := range args {
		array := strings.Split(arg, "-")

		ArgList[array[0]] = VTypeFromString(array[1])
	}

	return ArgList
}

func NewFunction(Name string, VariableList map[string]*Variable,
	ActionList map[string]*Action, ArgList map[string]VTYPE, function func()) *Function {
	return &Function{Name: Name, VariableList: VariableList, ActionList: ActionList, Args: ArgList, Executable: function}
}

func NewVariableList(vars ...*Variable) *map[string]*Variable {
	VariableList := make(map[string]*Variable)

	for _, variable := range vars {
		VariableList[variable.Name] = variable
	}

	return &VariableList
}

func NewActionList(acts ...*Action) *map[string]*Action {
	ActionList := make(map[string]*Action)

	for _, action := range acts {
		ActionList[action.Name] = action
	}

	return &ActionList
}
