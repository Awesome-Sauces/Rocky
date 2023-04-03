package eval

import (
	"fmt"

	"github.com/Awesome-Sauces/Rocky/core/tokenizer"
)

func Eval(tokenMap map[int]*tokenizer.Token) {

	math := NewMathematical("5*3")

	fmt.Println(math.GetOutput())

	/*
		for i := 0; i < len(tokenMap); i++ {
			fmt.Println("Order: "+strconv.Itoa(tokenMap[i].Order),
				" Type: "+tokenMap[i].Type.ToString(),
				" Value: "+tokenMap[i].Value)
		}
	*/
}
