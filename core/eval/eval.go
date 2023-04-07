package eval

import (
	"fmt"

	"github.com/Awesome-Sauces/Rocky/core/tokenizer"
)

func Eval(tokenMap map[int]*tokenizer.Token) {

	/*
				math := NewMathematical("6/2")

				fmt.Println(math.GetOutput())

		fmt.Println("Order: "+strconv.Itoa(order),
					" Type: "+Type.ToString(),
					" Value: "+value)

	*/

	variables := make(map[int]*Variable)

	for i := 0; i < len(tokenMap); i++ {

		//order := tokenMap[i].Order
		Type := tokenMap[i].Type
		value := tokenMap[i].Value

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

			variables[len(variables)] = NewVariable(tokenMap[i+1].Value,
				NewVType(value),
				tokenMap[i+3].Value)

		}

		/*
			Order: 0  Type: TYPE  Value: int
			Order: 1  Type: IDENTIFIER  Value: num
			Order: 2  Type: EQUAL  Value: =
			Order: 3  Type: NUMBER  Value: 12345
			Order: 4  Type: STOP  Value: ;
		*/

	}

	for i := 0; i < len(variables); i++ {
		fmt.Println("Name: " + variables[i].Name + " Type: " + variables[i].Type.ToString() + " Data: " + variables[i].Data)
	}

}
