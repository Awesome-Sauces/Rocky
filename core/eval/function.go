package eval

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/novalagung/golpal"
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
		func(function *Function) {
			output := ""

			for _, variable := range function.VariableList {
				output += variable.Data
			}

			fmt.Println(strings.ReplaceAll(output, "\"", ""))
		})

	api := NewFunction("api",
		make(map[string]*Variable), make(map[string]*Action), NewArgs("input-STRING"),
		func(function *Function) {
			output := ""

			type Message struct {
				Text string `json:"text"`
			}

			helloHandler := func(w http.ResponseWriter, r *http.Request) {
				message := Message{Text: "Hello, World!"}
				json.NewEncoder(w).Encode(message)
			}

			for _, variable := range function.VariableList {
				output += variable.Data
			}

			output = strings.ReplaceAll(output, "\"", "")

			http.HandleFunc("/"+output, helloHandler)
			log.Fatal(http.ListenAndServe(":8080", nil))
		})

	goCode := NewFunction("goCode",
		make(map[string]*Variable), make(map[string]*Action), NewArgs("input-STRING"),
		func(function *Function) {

			dir, errr := os.Getwd()
			if errr != nil {
				fmt.Println("Error:", errr)
				return
			}

			fmt.Println("Working directory:", dir)

			output := ""

			for _, variable := range function.VariableList {
				output += variable.Data
			}

			if !strings.Contains(output, ".grk") {

				fmt.Println("ROCKY ERROR: INVALID FILE TYPE " + output + "\nMust be filename.grk")

				return
			}

			// Open the file in read-only mode
			file, err := os.Open(strings.ReplaceAll(output, "\"", ""))
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			defer file.Close()

			// Read the contents of the file
			var content []byte
			buffer := make([]byte, 100)
			for {
				n, err := file.Read(buffer)
				if err != nil {
					break
				}
				content = append(content, buffer[:n]...)
			}

			output, err = golpal.New().ExecuteRaw(string(content))
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(output)
		})

	RegisterFunction(print)
	RegisterFunction(api)
	RegisterFunction(goCode)
}

type Function struct {
	Name         string
	Args         map[string]VTYPE
	VariableList map[string]*Variable
	ActionList   map[string]*Action
	Executable   func(function *Function)
}

func (function *Function) Execute() {
	// To be made
	function.Executable(function)
}

func (function *Function) SetVariableList(varList map[string]*Variable) {
	function.VariableList = varList
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
	ActionList map[string]*Action, ArgList map[string]VTYPE, function func(function *Function)) *Function {
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
