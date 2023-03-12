package main

import (
	"fmt"
	"os"
	"strconv"

	// Rocky imports
	"github.com/Awesome-Sauces/Rocky/core/file"
	"github.com/Awesome-Sauces/Rocky/core/lexer"
	"github.com/Awesome-Sauces/Rocky/core/tokenizer"
	"github.com/Awesome-Sauces/Rocky/core/utils"
	"github.com/novalagung/golpal"
)

func main() {

	// Example: -tokenize enumNumber
	if utils.IsArg("-tokenize", 1) && utils.ArgExists(2) {
		i, err := strconv.Atoi(utils.GetArgName(2))

		utils.Check(err)

		token := tokenizer.CreateToken(tokenizer.TOKEN(i), "{", 0)

		fmt.Println(token.GetType().ToString())
	}

	// Example: -eval filename.extension
	if utils.IsArg("-eval", 1) && utils.ArgExists(2) {

		file.LoadFile(utils.GetArgName(2))

		fmt.Printf(*file.GetFile())
	}

	if utils.IsArg("-testGoEval", 1) && utils.ArgExists(2) {
		file.LoadFile(utils.GetArgName(2))

		temp := *file.GetFile()

		output, err := golpal.New().ExecuteRaw(temp)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(output)

	}

	// Example: -test filename.extension
	if utils.IsArg("-test", 1) && utils.ArgExists(2) {
		file.LoadFile(utils.GetArgName(2))
		//fmt.Printf("%s\n", *file.GetFile())
		tokens := lexer.LoopString(file.GetFile())

		for _, element := range tokens {
			fmt.Println("TYPE -> " + element.GetType().ToString() + " -> VALUE -> " + element.GetValue() + " -> ORDER -> " + strconv.Itoa(element.GetOrder()))
		}

	} else if utils.IsArg("-test", 1) && !utils.ArgExists(2) {
		fmt.Printf("Error: Please enter filename!")
		os.Exit(1)
	}

	fmt.Println()
}
