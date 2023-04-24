package eval

import (
	"fmt"
	"strings"

	"github.com/Awesome-Sauces/Rocky/core/tokenizer"
)

func max(a int, b int) int {
	if b > a || b == a {
		return b
	} else {
		return a
	}
}

func min(a int, b int) int {
	if b < a || b == a {
		return b
	} else {
		return a
	}
}

func TestEval(tokenMap map[int]*tokenizer.Token) {

	DefaultRegistery()

	index := -1
	revr := 0
	Type := tokenMap[max(index, 0)].Type
	value := tokenMap[max(index, 0)].Value

	variables := NewVariableRegistrar()
	//functionList := make(map[string]*Function)

	forward := func(n int) bool {

		if index == len(tokenMap)-1 {
			return false
		}

		for j := 0; j < n; j++ {
			index = min(len(tokenMap), index+1)
			Type = tokenMap[index].Type
			value = tokenMap[index].Value
			revr++
		}

		return true
	}

	revert := func() {
		index -= revr
		revr = 0
		Type = tokenMap[index].Type
		value = tokenMap[index].Value

	}

	ifForward := func(typ tokenizer.TOKEN) bool {
		if tokenMap[index+1].Type != typ {
			//fmt.Println("First Type: " + tokenMap[i+1].Type.ToString() + " Second Type: " + typ.ToString())
			return false
		}

		return forward(1)
	}

	for forward(1) {

		// Variable Init Code
		if Type == tokenizer.TYPE {
			do := true

			vType := NewVType(value)

			do = ifForward(tokenizer.IDENTIFIER)

			Name := value

			do = ifForward(tokenizer.EQUAL)

			forward(1)

			// Later implement a handler for Data types and function returns
			Data := ""

			for {
				if Type != tokenizer.STOP {
					if Type == tokenizer.ADD {
						forward(1)
						continue
					}

					if Type == tokenizer.IDENTIFIER {

						Data += strings.ReplaceAll(variables.GetVariable(strings.ReplaceAll(value, " ", "")).Data, "\"", "")
						forward(1)
					} else {

						Data += strings.ReplaceAll(value, "\"", "")
						forward(1)
					}

				} else {

					break
				}
			}

			// Semicolon Detection

			do = Type == tokenizer.STOP

			if do {
				variable := NewVariable(Name, vType, Data)

				variables.RegisterVariable(variable)

			} else {
				revert()
			}
		}

		// Function Code
		if Type == tokenizer.IDENTIFIER &&
			tokenMap[index+1].Type == tokenizer.LPAREN {

			// Find the end point of function call
			i := 1
			for {
				i++
				if tokenMap[i+index].Type == tokenizer.RPAREN {
					break
				}

				if i >= 1000 {
					break
				}
			}

			function := GetFunction(strings.TrimSpace(value))
			list := make(map[string]*Variable)

			for forward(1) {

				if Type == tokenizer.IDENTIFIER {
					fmt.Println("Type Data: " + Type.ToString())
					fmt.Println("Value Data: " + value)

					vari := variables.GetVariable(strings.TrimSpace(value))

					fmt.Println("Variable: ", vari)

					list[value] = vari
				}
			}

			function.SetVariableList(list)

			function.Execute()
		}
	}

	fmt.Println(variables.GetVariable("filename").Data)

}

func Eval(tokenMap map[int]*tokenizer.Token) {

	DefaultRegistery()

	/*
				math := NewMathematical("6/2")

				fmt.Println(math.GetOutput())

		fmt.Println("Order: "+strconv.Itoa(order),
					" Type: "+Type.ToString(),
					" Value: "+value)

	*/

	for i := 0; i < len(tokenMap); i++ {

		//order := tokenMap[i].Order
		Type := tokenMap[i].Type
		value := tokenMap[i].Value

		revr := 0

		forward := func(n int) {
			for j := 0; j < n; j++ {
				i = min(len(tokenMap), i+1)
				Type = tokenMap[i].Type
				value = tokenMap[i].Value
				revr++
			}
		}

		revert := func() {
			i -= revr
			revr = 0
			Type = tokenMap[i].Type
			value = tokenMap[i].Value

		}

		ifForward := func(typ tokenizer.TOKEN) bool {
			if tokenMap[i+1].Type != typ {
				//fmt.Println("First Type: " + tokenMap[i+1].Type.ToString() + " Second Type: " + typ.ToString())
				return false
			}

			i = min(len(tokenMap), i+1)
			Type = tokenMap[i].Type
			value = tokenMap[i].Value

			revr++

			return true
		}

		//fmt.Println("Type: " + Type.ToString() + " Value: " + value)

		// Variable Init Code
		if Type == tokenizer.TYPE {
			do := true

			vType := NewVType(value)

			do = ifForward(tokenizer.IDENTIFIER)

			Name := value

			do = ifForward(tokenizer.EQUAL)

			forward(1)

			// Later implement a handler for Data types and function returns
			Data := value

			// Semicolon Detection

			do = ifForward(tokenizer.STOP)

			if do {
				NewVariable(Name, vType, Data)
			} else {
				revert()
			}
		}

		/*
					if Type == tokenizer.TYPE &&
				tokenMap[i+1].Type == tokenizer.IDENTIFIER &&
				tokenMap[i+2].Type == tokenizer.EQUAL &&
				(tokenMap[i+3].Type == tokenizer.NUMBER || tokenMap[i+3].Type == tokenizer.STRING) &&
				tokenMap[i+4].Type == tokenizer.STOP {
				vType := NewVType(value)

				if (vType == INT && tokenMap[i+3].Type == tokenizer.STRING) ||
					(vType == STRING && tokenMap[i+3].Type == tokenizer.NUMBER) {
					fmt.Println("An error occured when initializing the variable &" + tokenMap[i+1].Value + "\nINVALID TYPE ERROR")
					continue
				}

				RegisterVariable(NewVariable(tokenMap[i+1].Value,
					NewVType(value),
					tokenMap[i+3].Value))
			}
		*/

		if Type == tokenizer.IDENTIFIER &&
			tokenMap[i+1].Type == tokenizer.LPAREN {

			// Find the end point of function call
			index := 1
			for {
				index++
				if tokenMap[i+index].Type == tokenizer.RPAREN {
					break
				}

				if index >= 1000 {
					break
				}
			}

			function := GetFunction(value)
			list := make(map[string]*Variable)

			for e := 1; e < index; e++ {
				if tokenMap[i+e].Type == tokenizer.IDENTIFIER {

					list[tokenMap[i+e].Value] = &Variable{}

				}
			}

			function.SetVariableList(list)

			function.Execute()
		}

		/*
			Order: 0  Type: TYPE  Value: int
			Order: 1  Type: IDENTIFIER  Value: num
			Order: 2  Type: EQUAL  Value: =
			Order: 3  Type: NUMBER  Value: 12345
			Order: 4  Type: STOP  Value: ;
		*/

	}

}
