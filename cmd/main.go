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
)

func main() {

	if utils.IsArg("-tokenize", 1) && utils.ArgExists(2) && utils.ArgExists(3) {
		i, err := strconv.Atoi(utils.GetArgName(2))

		utils.Check(err)

		token := tokenizer.CreateToken(tokenizer.TOKEN(i), "{")

		fmt.Println(token.GetType().ToString())
	}

	if utils.IsArg("-eval", 1) && utils.ArgExists(2) {

		file.LoadFile(utils.GetArgName(2))

		fmt.Printf(*file.GetFile())
	}

	if utils.IsArg("-test", 1) && utils.ArgExists(2) {
		file.LoadFile(utils.GetArgName(2))
		//fmt.Printf("%s\n", *file.GetFile())
		lexer.LoopString(file.GetFile())

	} else if utils.IsArg("-test", 1) && !utils.ArgExists(2) {
		fmt.Printf("Error: Please enter filename!")
		os.Exit(1)
	}

	fmt.Println()
}
