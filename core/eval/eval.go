package eval

import (
	"fmt"

	"github.com/Awesome-Sauces/Rocky/core/tokenizer"
)

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

		//fmt.Println("Type: " + Type.ToString() + " Value: " + value)

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
<<<<<<< HEAD
				tokenMap[i+3].Value))
=======
				tokenMap[i+3].Value)
>>>>>>> parent of a1c6f28 (nil)
		}

		/*
			Order: 0  Type: TYPE  Value: int
			Order: 1  Type: IDENTIFIER  Value: num
			Order: 2  Type: EQUAL  Value: =
			Order: 3  Type: NUMBER  Value: 12345
			Order: 4  Type: STOP  Value: ;
		*/

	}

	for var_name, var_obj := range variables {
		fmt.Println("Name: " + var_name + " Type: " +
			var_obj.Type.ToString() + " Data: " + var_obj.Data)
	}

}
